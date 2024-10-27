package handlers

import (
	"galihwicaksono90/musikmarching-be/internal/services/account"
	"galihwicaksono90/musikmarching-be/internal/services/auth"
	"galihwicaksono90/musikmarching-be/internal/services/score"
	"galihwicaksono90/musikmarching-be/internal/services/profile"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/sirupsen/logrus"
)

type Handler struct {
  logger *logrus.Logger
  store *db.Store
  auth auth.AuthService
  account account.AccountService
  profile profile.ProfileService
  score score.ScoreService
}

func New(
  logger * logrus.Logger, 
  store *db.Store, 
  auth auth.AuthService, 
  account account.AccountService,
  profile profile.ProfileService,
  score score.ScoreService,
) *Handler {
  return &Handler{
    logger,
    store,
    auth,
    account,
    profile,
    score,
  }
}
