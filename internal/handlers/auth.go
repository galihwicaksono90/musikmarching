package handlers

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth/gothic"
)

func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Logout")

	err := gothic.Logout(w, r)
	if err != nil {
		h.logger.Println(err)
		return
	}

	h.auth.RemoveUserSession(w, r)

	w.Header().Set("Location", "/")
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

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
