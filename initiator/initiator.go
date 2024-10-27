package initiator

import (
	"context"
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/routings"
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/internal/services/account"
	"galihwicaksono90/musikmarching-be/internal/services/auth"
	"galihwicaksono90/musikmarching-be/internal/services/profile"
	"galihwicaksono90/musikmarching-be/internal/services/score"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"

	"github.com/galihwicaksono90/musikmarching-be/pkg/config"
)

func Init() {
	ctx := context.Background()
	logger := logrus.New()

	config, err := config.LoadConfig("./")

	conn, err := pgx.Connect(ctx, config.DB_SOURCE)
	if err != nil {
		logger.Fatal(err)
	}

	store := db.NewStore(conn)

	sessionStore := auth.NewSessionStore(auth.SessionOptions{
		CookiesKey: "secretkey",
		MaxAge:     60 * 60 * 24 * 2,
		HttpOnly:   true,
		Secure:     true,
	})

	// services
	authService := auth.NewAuthService(logger, sessionStore)
	accountService := account.NewAccountService(logger, store)
	profileService := profile.NewProfileService(logger, store)
	scoreService := score.NewProfileService(logger, store)

	handler := handlers.New(
		logger, 
		&store, 
		authService, 
		accountService,
		profileService,
		scoreService,
	)

	// routings
	router := mux.NewRouter()

	routings.AuthRouting(handler, router)
	routings.HomeRouting(handler, router)
	routings.ScoreRouting(handler, router)
	routings.ProfileRouting(handler, router)

	port := fmt.Sprintf(":%s", config.PORT)

	fmt.Printf("listening to port %s \n", port)

	// log.Printf("Server: Listening on %s:%s\n", "http://localohst", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
