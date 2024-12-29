package handlers

import (
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"
	"net/http"
	"reflect"

	"github.com/markbates/goth/gothic"
	"github.com/spf13/viper"
)

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	rerouteUrl := viper.GetString("GOOGLE_REROUTE_URL")
	h.logger.Println("Logout")

	err := gothic.Logout(w, r)
	if err != nil {
		h.logger.Println(err)
		return
	}

	h.auth.RemoveUserSession(w, r)

	w.Header().Set("Location", rerouteUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) HandleProviderLogin(w http.ResponseWriter, r *http.Request) {
	// try to get the user without re-authenticating
	if u, err := gothic.CompleteUserAuth(w, r); err == nil {
		h.logger.Printf("User already authenticated! %v", u)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (h *Handler) HandleAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	redirectUrl := viper.GetString("GOOGLE_REROUTE_URL")
	u, err := gothic.CompleteUserAuth(w, r)
	h.logger.Println("HandleAuthCallbackFunction======")
	h.logger.Println(u)
	h.logger.Println("redirectUrl", redirectUrl)
	h.logger.Println("======HandleAuthCallbackFunction")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	user, err := h.account.UpsertAccount(u)
	if err != nil {
		h.logger.Println(err)
		return
	}

	err = h.auth.StoreUserSession(w, r, user)
	if err != nil {
		h.logger.Println(err)
		return
	}

	http.Redirect(w, r, redirectUrl, http.StatusPermanentRedirect)
}

func (h *Handler) HandleMe(w http.ResponseWriter, r *http.Request) {
	u := h.getSessionUser(r)
	h.logger.Println("HandleMe======")
	h.logger.Println(u)
	h.logger.Println("======HandleMe")
	if u == nil {
		h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), u)
		return
	}

	if u.RoleName != db.RolenameContributor {
		h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), u)
		return
	}

	c, err := h.contributor.GetByID(u.ID)
	if err != nil {
		h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), u)
		return
	}

	user := make(map[string]interface{})
	user["id"] = u.ID
	user["email"] = u.Email
	user["name"] = c.FullName
	user["role_name"] = u.RoleName
	user["is_verified"] = c.IsVerified
	user["verified_at"] = c.VerifiedAt

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), user)
}

func (h *Handler) getSessionUser(r *http.Request) *model.SessionUser {
	u := r.Context().Value(middlewares.UserContextName)
	if u == nil {
		return nil
	}

	user, ok := u.(*model.SessionUser)
	if !ok {
		return nil
	}

	return user
}

func structToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		result[typ.Field(i).Name] = field.Interface()
	}
	return result
}
