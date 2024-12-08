package handlers

import (
	"encoding/json"
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	// "galihwicaksono90/musikmarching-be/internal/services/auth"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/spf13/viper"
)

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Logout")

	err := gothic.Logout(w, r)
	if err != nil {
		h.logger.Println(err)
		return
	}

	h.auth.RemoveUserSession(w, r)

	w.Header().Set("Location", "http://localhost:5173")
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
	json.NewEncoder(w).Encode(u)
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
