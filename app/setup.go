package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	contestHttpDelivery "github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/contest/delivery/http"
	contestRepository "github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/contest/repository"
	contestUsecase "github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/contest/usecase"

	googleHttpDelivery "github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/google/delivery/http"
	googleUsecase "github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/google/usecase"
	oauthRepository "github.com/KampungBudaya/Kampung-Budaya-2023-BE/api/oauth/repository"

	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/config"
	"github.com/KampungBudaya/Kampung-Budaya-2023-BE/util/response"
)

func Run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	db, err := config.StartMySQLConn()
	if err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln(err.Error())
		}
	}()

	port := os.Getenv("APP_PORT")

	app := mux.NewRouter()

	app.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+port+"/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	api := app.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response.Success(w, http.StatusOK, "I'm fine and healthy! nice to see you :)")
	}).Methods(http.MethodGet)

	sheetsService, err := config.SetupSheets(os.Getenv("SHEETS_ID"), os.Getenv("SHEETS_CREDENTIAL_PATH"))
	if err != nil {
		return err
	}

	contestRepository := contestRepository.NewContestRepository(db)
	contestUsecase := contestUsecase.NewContestUsecase(contestRepository, sheetsService)
	contestHttpDelivery.NewContestHandler(v1, contestUsecase)

	oauthRepository := oauthRepository.NewOAuthRepository(db)
	googleUsecase := googleUsecase.NewGoogleUsecase(oauthRepository, os.Getenv("GOOGLE_CLIENT_ID"))
	googleHttpDelivery.NewGoogleHandler(v1, googleUsecase)

	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "Origin", "X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{os.Getenv("ALLOWED_ORIGIN")})
	allowedMethods := handlers.AllowedMethods([]string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPatch,
	})

	fmt.Println("Server running on port " + port)
	if err := http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(app)); err != nil {
		return err
	}

	return nil
}
