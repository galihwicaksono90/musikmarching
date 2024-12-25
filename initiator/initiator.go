package initiator

import (
	"encoding/json"
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/routings"
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/internal/services/account"
	"galihwicaksono90/musikmarching-be/internal/services/auth"
	"galihwicaksono90/musikmarching-be/internal/services/contributor"
	"galihwicaksono90/musikmarching-be/internal/services/file"
	"galihwicaksono90/musikmarching-be/internal/services/purchase"
	"galihwicaksono90/musikmarching-be/internal/services/score"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/pkg/dbpool"
	"galihwicaksono90/musikmarching-be/pkg/email"
	fileStorage "galihwicaksono90/musikmarching-be/pkg/file-storage"
	"galihwicaksono90/musikmarching-be/pkg/logger"
	"galihwicaksono90/musikmarching-be/pkg/middlewares"
	"galihwicaksono90/musikmarching-be/pkg/validator"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"

	"galihwicaksono90/musikmarching-be/pkg/config"
)

func Init() {
	logger := logger.NewLogger()
	validate := validator.New()

	// load env config
	config, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}

	// file storage (minio)
	fileStorage := fileStorage.NewStorage(logger, config)
	email := email.NewEmail(config)

	pool, err := dbpool.NewDBPool(config)
	if err != nil {
		logger.Fatal(err)
	}

	store := db.NewStore(pool)

	sessionStore := auth.NewSessionStore(auth.SessionOptions{
		CookiesKey: "secretkey",
		MaxAge:     60 * 60 * 24 * 4,
		HttpOnly:   true,
		Secure:     true,
	})

	// services
	authService := auth.NewAuthService(logger, sessionStore)
	accountService := account.NewAccountService(logger, store)
	scoreService := score.NewScoreService(logger, store)
	purchaseService := purchase.NewPurchaseService(logger, store)
	contributorService := contributor.NewContributorService(logger, store)
	fileService := file.NewFileService(logger, fileStorage)

	// initiate new handler
	handler := handlers.New(
		logger,
		&store,
		authService,
		accountService,
		scoreService,
		purchaseService,
		contributorService,
		fileService,
		email,
		validate,
	)

	// routings
	router := mux.NewRouter()
	router.Use(middlewares.SessionMiddleware)
	routings.AuthRouting(handler, router)

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("pong")
	})

	routings.Routings(handler, router)

	// serve static files
	fs := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	port := fmt.Sprintf(":%s", config.PORT)

	fmt.Printf("listening to port %s \n", port)

	log.Fatalln(http.ListenAndServe(port, router))
}
