package handlers

import (
  "galihwicaksono90/musikmarching-be/internal/services/account"
  "galihwicaksono90/musikmarching-be/internal/services/auth"
  "galihwicaksono90/musikmarching-be/internal/services/score"
  "galihwicaksono90/musikmarching-be/internal/services/purchase"
  db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

  "github.com/sirupsen/logrus"
)

type Handler struct {
  logger *logrus.Logger
  store *db.Store
  auth auth.AuthService
  account account.AccountService
  score score.ScoreService
  purchase purchase.PurchaseService
}

func New(
  logger * logrus.Logger, 
  store *db.Store, 
  auth auth.AuthService, 
  account account.AccountService,
  score score.ScoreService,
  purchase purchase.PurchaseService,
) *Handler {
  return &Handler{
    logger,
    store,
    auth,
    account,
    score,
    purchase,
  }
}
