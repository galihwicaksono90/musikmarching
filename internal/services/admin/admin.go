package admin

import (
	"context"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type AdminService interface {
	VerifyContributorApplication(id uuid.UUID) error
}

type adminService struct {
	logger *logrus.Logger
	store  db.Store
}

func (a *adminService) VerifyContributorApplication(id uuid.UUID) error {
	ctx := context.Background()
	return a.store.ExecTx(ctx, func(tx *db.Queries) error {
		if err := a.store.VerifyContributorApply(ctx, id); err != nil {
			return err
		}

		if _, err := a.store.UpdateAccountRole(ctx, db.UpdateAccountRoleParams{
			Rolename: "contributor",
			ID:       id,
		}); err != nil {
			return err
		}

		return a.store.CreateContributorFromContributorApply(ctx, id)
	})
}

// func (a *adminService) VerifyContributorApplication(id uuid.UUID) error {
// 	ctx := context.Background()
// 	return a.store.ExecTx(ctx, func(tx *db.Queries) error {
// 		application, err := tx.GetContributorApplyByAccountID(ctx, id)
// 		if err != nil {
// 			return err
// 		}
// 		if application.IsVerified {
// 			return errors.New("Contributor application already verified")
// 		}
//
// 		if _, err := tx.GetContributorById(ctx, application.ID); err == nil {
// 			return errors.New("Contributor alredy exists")
// 		}
//
// 		params := db.CreateContributorParams{
// 			ID:                id,
// 			FullName:          application.FullName,
// 			PhoneNumber:       application.PhoneNumber,
// 			MusicalBackground: application.MusicalBackground,
// 			Education:         application.Education,
// 			Experience:        application.Experience,
// 			PortofolioLink:    application.PortofolioLink,
// 		}
//
// 		if _, err := a.store.CreateContributor(ctx, params); err != nil {
// 			a.logger.Println("=====")
// 			a.logger.Errorln(err)
// 			a.logger.Println("=====")
// 			return err
// 		}
//
// 		return nil
// 	})
// }

func NewAdminService(logger *logrus.Logger, store db.Store) AdminService {
	return &adminService{
		logger,
		store,
	}
}
