package initiator

import (
	"encoding/json"
	"fmt"
	"galihwicaksono90/musikmarching-be/internal/constants/routings"
	"galihwicaksono90/musikmarching-be/internal/handlers"
	"galihwicaksono90/musikmarching-be/internal/services/account"
	"galihwicaksono90/musikmarching-be/internal/services/allocation"
	"galihwicaksono90/musikmarching-be/internal/services/auth"
	"galihwicaksono90/musikmarching-be/internal/services/category"
	"galihwicaksono90/musikmarching-be/internal/services/contributor"
	"galihwicaksono90/musikmarching-be/internal/services/file"
	"galihwicaksono90/musikmarching-be/internal/services/instrument"
	"galihwicaksono90/musikmarching-be/internal/services/purchase"
	"galihwicaksono90/musikmarching-be/internal/services/score"
	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
	"galihwicaksono90/musikmarching-be/pkg/cors"
	"galihwicaksono90/musikmarching-be/pkg/dbpool"
	"galihwicaksono90/musikmarching-be/pkg/email"
	fileStorage "galihwicaksono90/musikmarching-be/pkg/file-storage"
	"galihwicaksono90/musikmarching-be/pkg/logger"
	"galihwicaksono90/musikmarching-be/pkg/validator"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"

	"galihwicaksono90/musikmarching-be/pkg/config"
)

func Init() {
	logger := logger.NewLogger()
	validate := validator.New()

	cors := cors.NewCorsHandler()

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
		CookiesKey: config.SessionSecret,
		MaxAge:     60 * 60 * 24 * 4,
		HttpOnly:   true,
		Secure:     false,
		Domain:     config.SessionDomain,
	})

	// services
	authService := auth.NewAuthService(logger, sessionStore)
	accountService := account.NewAccountService(logger, store)
	scoreService := score.NewScoreService(logger, store)
	purchaseService := purchase.NewPurchaseService(logger, store)
	instrumentService := instrument.NewInstrumentService(logger, store)
	categoryService := category.NewCategoryService(logger, store)
	allocationService := allocation.NewAllocationService(logger, store)
	contributorService := contributor.NewContributorService(logger, store, instrumentService)
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
		instrumentService,
		categoryService,
		allocationService,
		fileService,
		email,
		validate,
	)

	// routings
	router := mux.NewRouter()
	routings.AuthRouting(handler, router)
	routings.Routings(handler, router)
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("pong")
	})

	port := fmt.Sprintf(":%s", config.Port)

	fmt.Printf("listening to port %s \n", port)

	log.Fatalln(http.ListenAndServe(port, cors(router)))
}
