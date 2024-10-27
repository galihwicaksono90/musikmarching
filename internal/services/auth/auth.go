package auth

import (
	"encoding/gob"
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AuthService interface {
	GetSessionUser(r *http.Request) (model.SessionUser, error)
	StoreUserSession(w http.ResponseWriter, r *http.Request, user *model.SessionUser) error
	RemoveUserSession(w http.ResponseWriter, r *http.Request)
}

type authService struct {
	logger *logrus.Logger
}

// GetSessionUser implements AuthService.
func (a *authService) GetSessionUser(r *http.Request) (model.SessionUser, error) {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		return model.SessionUser{}, err
	}

	u := session.Values["user"]

	if u == nil {
		return model.SessionUser{}, fmt.Errorf("user is not authenticated! %v", u)
	}
	return model.SessionUser{
		ID:       u.(model.SessionUser).ID,
		Email:    u.(model.SessionUser).Email,
		Name:     u.(model.SessionUser).Name,
		RoleName: u.(model.SessionUser).RoleName,
	}, nil
}

// RemoveUserSession implements AuthService.
func (a *authService) RemoveUserSession(w http.ResponseWriter, r *http.Request) {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		a.logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = model.SessionUser{}

	session.Options.MaxAge = -1

	session.Save(r, w)
}

// StoreUserSession implements AuthService.
func (a *authService) StoreUserSession(w http.ResponseWriter, r *http.Request, user *model.SessionUser) error {
	session, _ := gothic.Store.Get(r, SessionName)

	a.logger.Println(session)

	session.Values["user"] = user

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func NewAuthService(logger *logrus.Logger, store sessions.Store) AuthService {
	gob.Register(model.SessionUser{})

	gothic.Store = store

	goth.UseProviders(
		google.New(
			viper.GetString("GOOGLE_CLIENT_ID"),
			viper.GetString("GOOGLE_CLIENT_SECRET"),
			viper.GetString("GOOGLE_CALLBACK_URL"),
			"email",
			"profile",
		),
	)
	return &authService{
		logger,
	}
}
