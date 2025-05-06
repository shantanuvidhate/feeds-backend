package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/shantanuvidhate/feeds-backend/internal/auth"
	"github.com/shantanuvidhate/feeds-backend/internal/db"
	"github.com/shantanuvidhate/feeds-backend/internal/env"
	"github.com/shantanuvidhate/feeds-backend/internal/mailer"
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
		addr:        env.GetString("ADDR", ":8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:3000"),
		db: dbConfig{addr: env.GetString("DB_ADDR", "postgres://user:password@localhost:5432/feeds?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:      time.Hour * 24, // 24 hours
			smtpUser: env.GetString("SMTP_USER", ""),
			mailTrap: mailTrapConfig{
				apiKey: env.GetString("MAILTRAP_API_KEY", "818f1f9243ae0e"),
			},
		},
		auth: authConfig{
			basic: authBasicConfig{
				user: env.GetString("BASIC_AUTH_USER", "admin"),
				pass: env.GetString("BASIC_AUTH_PASS", "admin"),
			},
			token: authTokenConfig{
				secret: env.GetString("JWT_SECRET", "secret"),
				iss:    env.GetString("JWT_ISS", "http://localhost:8080"),
				aud:    env.GetString("JWT_AUD", "http://localhost:8080"),
				exp:    time.Hour * 24, // 24 hours
			},
		},
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

	mailtrap, err := mailer.NewMailTrapClient(cfg.mail.mailTrap.apiKey, cfg.mail.smtpUser)
	if err != nil {
		logger.Fatal(err.Error())
	}

	store := store.NewStorage(db)

	jwtAuthenticator := auth.NewJWTAuthenticator(
		cfg.auth.token.secret,
		cfg.auth.token.iss,
		cfg.auth.token.aud,
	)
	app := &application{
		config:        cfg,
		store:         store,
		logger:        logger,
		mailer:        mailtrap,
		authenticator: jwtAuthenticator,
	}

	mux := app.mount()
	logger.Fatal("Failed to start server", zap.Error(app.run(mux)))
}
