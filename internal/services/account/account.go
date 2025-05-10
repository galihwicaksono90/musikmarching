package account

import (
	"context"
	"encoding/json"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/markbates/goth"
	"github.com/sirupsen/logrus"
)

type AccountService interface {
	GetUserByID(uuid.UUID) (*model.Account, error)
	GetUserByEmail(string) (*model.Account, error)
	UpsertAccount(goth.User) (*model.SessionUser, error)
	UpdateRole(uuid.UUID, db.Rolename) error
	ApplyContributor(db.CreateContributorApplyParams) (db.ContributorApply, error)
}

type accountService struct {
	logger *logrus.Logger
	store  db.Store
}

// ApplyContributor implements AccountService.
func (s *accountService) ApplyContributor(params db.CreateContributorApplyParams) (db.ContributorApply, error) {
	ctx := context.Background()
	return s.store.CreateContributorApply(ctx, params)
}

// GetUserByID implements AccountService.
func (s *accountService) GetUserByID(id uuid.UUID) (*model.Account, error) {
	account, err := s.store.GetAccountById(context.Background(), id)

	if err != nil {
		s.logger.Error(err)
		return &model.Account{}, err
	}
	return &model.Account{
		ID:       account.ID,
		Email:    account.Email,
		Name:     account.Name,
		RoleName: account.RoleName,
	}, nil
}

// UpdateRole implements AccountService.
func (s *accountService) UpdateRole(id uuid.UUID, role db.Rolename) error {
	_, err := s.store.UpdateAccountRole(context.Background(), db.UpdateAccountRoleParams{
		ID:       id,
		Rolename: role,
	})
	return err
}

// CreateOrUpdateAccount implements AccountService.
func (s *accountService) UpsertAccount(user goth.User) (*model.SessionUser, error) {
	ctx := context.Background()
	accountCheck, err := s.store.GetAccountByEmail(ctx, user.Email)
	var id uuid.UUID

	if err != nil {
		newAccount, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
			Email: user.Email,
			Name:  user.Name,
			Pictureurl: pgtype.Text{
				String: user.AvatarURL,
				Valid:  true,
			}},
		)
		if err != nil {
			return nil, err
		}

		id = newAccount.ID
	} else {
		id, err = s.store.UpdateAccount(ctx,
			db.UpdateAccountParams{
				Name: user.Name,
				Pictureurl: pgtype.Text{
					String: user.AvatarURL,
					Valid:  true,
				},
				ID: accountCheck.ID,
			},
		)
		if err != nil {
			return nil, err
		}
	}

	res, err := s.store.GetAccountById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.SessionUser{
		ID:         id,
		Email:      user.Email,
		Name:       user.Name,
		RoleName:   res.RoleName,
		PictureUrl: user.AvatarURL,
	}, nil
}

// GetUserByEmail implements AccountService.
func (a *accountService) GetUserByEmail(email string) (*model.Account, error) {
	account, err := a.store.GetAccountByEmail(context.Background(), email)
	if err != nil {
		a.logger.Error(err)
		return &model.Account{}, err
	}
	return &model.Account{
		ID:       account.ID,
		Email:    email,
		Name:     account.Name,
		RoleName: account.RoleName,
	}, nil
}

func NewAccountService(logger *logrus.Logger, store db.Store) AccountService {
	return &accountService{
		logger,
		store,
	}
}
