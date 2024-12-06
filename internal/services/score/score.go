package score

import (
	"context"
	"errors"
	"galihwicaksono90/musikmarching-be/internal/constants/model"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

type ScoreService interface {
	GetManyByContributorId(account_id uuid.UUID) ([]db.Score, error)
	Create(model.CreateScoreDTO) (uuid.UUID, error)
	Update(uuid.UUID, model.UpdateScoreDTO) error
	GetVerified(db.GetVerifiedScoresParams) *[]db.GetVerifiedScoresRow
	GetVerifiedById(id uuid.UUID) (db.GetVerifiedScoreByIdRow, error)
	UploadPdfFile(*http.Request) (url string, err error)
	UploadMusicFile(*http.Request) (url string, err error)
	GetById(id uuid.UUID) (db.Score, error)
	GetByContirbutorID(db.GetScoresByContributorIDParams) ([]db.GetScoresByContributorIDRow, error)
	GetAll() ([]db.Score, error)
	Verify(id uuid.UUID) error
}

type scoreService struct {
	logger      *logrus.Logger
	store       db.Store
	fileStorage *minio.Client
}

// GetByContirbutorID implements ScoreService.
func (s *scoreService) GetByContirbutorID(params db.GetScoresByContributorIDParams) ([]db.GetScoresByContributorIDRow, error) {
	return s.store.GetScoresByContributorID(context.Background(), params)
}

// VerifyScore implements ScoreService.
func (s *scoreService) Verify(id uuid.UUID) error {
	return s.store.VerifyScore(context.Background(), id)
}

// GetAll implements ScoreService.
func (s *scoreService) GetAll() ([]db.Score, error) {
	ctx := context.Background()
	result, err := s.store.GetScoresPaginated(ctx, db.GetScoresPaginatedParams{
		Limit:   2,
		Offset:  0,
		Column3: "price_desc",
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetScoreById implements ScoreService.
func (s *scoreService) GetById(id uuid.UUID) (db.Score, error) {
	ctx := context.Background()
	return s.store.GetScoreById(ctx, id)
}

// GetManyByContributorId implements ScoreService.
func (s *scoreService) GetManyByContributorId(account_id uuid.UUID) ([]db.Score, error) {
	panic("unimplemented")
}

// GetVerified implements ScoreService.
func (s *scoreService) GetVerified(params db.GetVerifiedScoresParams) *[]db.GetVerifiedScoresRow {
	scores, err := s.store.GetVerifiedScores(context.Background(), params)

	if err != nil {
		return &[]db.GetVerifiedScoresRow{}
	}
	return &scores
}

// GetVerifiedById implements ScoreService.
func (s *scoreService) GetVerifiedById(id uuid.UUID) (db.GetVerifiedScoreByIdRow, error) {
	return s.store.GetVerifiedScoreById(context.Background(), id)
}

// UploadPdfFile implements ScoreService.
func (s *scoreService) UploadPdfFile(r *http.Request) (url string, err error) {
	file, header, err := r.FormFile("pdf-file")
	if err != nil {
		s.logger.Errorln(err)
		return "", err
	}

	pdfUploadInfo, err := utils.UploadFile(s.fileStorage, file, header, model.PDF_LOCATION)
	if err != nil {
		s.logger.Errorln(err)
		return "", err
	}

	return pdfUploadInfo.Location, nil
}

// UploadMusicFile implements ScoreService.
func (s *scoreService) UploadMusicFile(r *http.Request) (url string, err error) {
	file, header, err := r.FormFile("music-file")
	if err != nil {
		s.logger.Errorln(err)
		return "", err
	}

	pdfUploadInfo, err := utils.UploadFile(s.fileStorage, file, header, model.MUSIC_LOCATION)
	if err != nil {
		s.logger.Errorln(err)
		return "", err
	}

	return pdfUploadInfo.Location, nil
}

// CreateScore implements ScoreService.
func (s *scoreService) Create(params model.CreateScoreDTO) (uuid.UUID, error) {
	ctx := context.Background()

	return s.store.CreateScore(ctx, db.CreateScoreParams{
		Title: params.Title,
		Price: pgtype.Numeric{
			Int:   params.Price,
			Valid: true,
		},
		PdfUrl: pgtype.Text{
			String: params.PdfUrl,
			Valid:  true,
		},
		MusicUrl: pgtype.Text{
			String: params.MusicUrl,
			Valid:  true,
		},
		ContributorID: params.ContributorID,
	})
}

// Update implements ScoreService.
func (s *scoreService) Update(scoreId uuid.UUID, params model.UpdateScoreDTO) error {
	ctx := context.Background()

	scoreCheck, err := s.store.GetScoreById(ctx, scoreId)
	if err != nil {
		return err
	}

	if scoreCheck.ContributorID != params.ContributorID {
		return errors.New("you are not the owner of this score")
	}

	if err := s.store.UpdateScore(ctx, db.UpdateScoreParams{
		Title: params.Title,
		Price: params.Price,
		ID:    scoreId,
	}); err != nil {
		return err
	}

	return nil
}

// GetScoresByContributorId implements ScoreService.
func (s *scoreService) GetScoresByContributorId(account_id uuid.UUID) ([]db.Score, error) {
	panic("unimplemented")
}

func NewScoreService(logger *logrus.Logger, store db.Store, fileStorage *minio.Client) ScoreService {
	return &scoreService{
		logger,
		store,
		fileStorage,
	}
}
