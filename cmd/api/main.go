package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/shantanuvidhate/feeds-backend/internal/auth"
	"github.com/shantanuvidhate/feeds-backend/internal/db"
	"github.com/shantanuvidhate/feeds-backend/internal/env"
	"github.com/shantanuvidhate/feeds-backend/internal/mailer"
	"github.com/shantanuvidhate/feeds-backend/internal/store"
	"github.com/shantanuvidhate/feeds-backend/internal/store/cache"
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
		redisCfg: redisConfig{
			addr:     env.GetString("REDIS_ADDR", "localhost:6379"),
			password: env.GetString("REDIS_PASSWORD", "password"),
			db:       env.GetInt("REDIS_DB", 0),
			enabled:  env.GetBool("REDIS_ENABLED", true),
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

	// Redis
	var rdb *redis.Client

	if cfg.redisCfg.enabled {
		rdb = cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.password, cfg.redisCfg.db)
		logger.Info("Redis connection established")
	}
	mailtrap, err := mailer.NewMailTrapClient(cfg.mail.mailTrap.apiKey, cfg.mail.smtpUser)
	if err != nil {
		logger.Fatal(err.Error())
	}

	store := store.NewStorage(db)
	cacheStorage := cache.NewRedisStorage(rdb)

	jwtAuthenticator := auth.NewJWTAuthenticator(
		cfg.auth.token.secret,
		cfg.auth.token.iss,
		cfg.auth.token.aud,
	)
	app := &application{
		config:        cfg,
		store:         store,
		cacheStorage:  cacheStorage,
		logger:        logger,
		mailer:        mailtrap,
		authenticator: jwtAuthenticator,
	}

	mux := app.mount()
	logger.Fatal("Failed to start server", zap.Error(app.run(mux)))
}
