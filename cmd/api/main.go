package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/shantanuvidhate/feeds-backend/internal/db"
	"github.com/shantanuvidhate/feeds-backend/internal/env"
	"github.com/shantanuvidhate/feeds-backend/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			Feed Backend API
//	@description	API for Feed Generation, a social network for user
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{addr: env.GetString("DB_ADDR", "postgres://user:password@localhost:5432/feeds?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	// Logger
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	defer db.Close()
	logger.Info("Database connection established")

	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	mux := app.mount()
	logger.Fatal("Failed to start server", zap.Error(app.run(mux)))
}
