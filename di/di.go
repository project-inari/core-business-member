// Package di provides dependency injection for the server
package di

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"

	"github.com/project-inari/core-business-member/config"
	"github.com/project-inari/core-business-member/handler"
	"github.com/project-inari/core-business-member/repository"
	"github.com/project-inari/core-business-member/service"
)

// New injects the dependencies for the server
func New(c *config.Config) {
	ctx := context.Background()

	// Sentry initialization
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:                c.SentryConfig.SentryDSN,
		Debug:              false,
		EnableTracing:      true,
		TracesSampleRate:   1.0,
		ProfilesSampleRate: 1.0,
	}); err != nil {
		slog.Error("error - [main.New] sentry initialization failed", slog.Any("error", err))
	}

	// Echo server initialization
	e := echo.New()
	setupServer(ctx, e, c)

	// MySQL initialization
	mysqlDB, err := newMySQL(mySQLOptions{
		host:         c.MySQLConfig.Host,
		username:     c.MySQLConfig.Username,
		password:     c.MySQLConfig.Password,
		database:     c.MySQLConfig.Database,
		timeout:      c.MySQLConfig.Timeout,
		maxIdleConns: c.MySQLConfig.MaxIdleConns,
		maxOpenConns: c.MySQLConfig.MaxOpenConns,
		maxLifetime:  c.MySQLConfig.MaxLifetime,
	})
	if err != nil {
		log.Panicf("error - [main.New] unable to connect to MySQL: %v", err)
	}
	defer func() {
		if err := mysqlDB.client.Close(); err != nil {
			slog.Error("error - [main.New] unable to close MySQL connection", slog.Any("error", err))
		}
	}()

	// Redis initialization
	redisClient, err := newRedis(redisOptions{
		host:     c.RedisConfig.Host,
		password: c.RedisConfig.Password,
		timeout:  c.RedisConfig.Timeout,
		maxRetry: c.RedisConfig.MaxRetry,
		poolSize: c.RedisConfig.PoolSize,
		db:       c.RedisConfig.DB,
	})
	if err != nil {
		log.Panicf("error - [main.New] unable to connect to Redis: %v", err)
	}
	defer func() {
		if err := redisClient.client.Close(); err != nil {
			slog.Error("error - [main.New] unable to close Redis connection", slog.Any("error", err))
		}
	}()

	// Repository initialization
	databaseRepo := repository.NewDatabaseRepository(repository.DatabaseRepositoryConfig{
		Database: c.MySQLConfig.Database,
	}, repository.DatabaseRepositoryDependencies{
		Client: mysqlDB.client,
	})

	cacheRepo := repository.NewCacheRepository(repository.CacheRepositoryConfig{
		KeyUserVerifiedAccount: c.RedisConfig.KeyUserVerifiedAccount,
	}, repository.CacheRepositoryDependencies{
		Client: redisClient.client,
	})

	// Service initialization
	service := service.New(service.Dependencies{
		DatabaseRepository: databaseRepo,
		CacheRepository:    cacheRepo,
	})

	// Handler initialization
	handler.New(e, handler.Dependencies{
		Service: service,
	})

	// HTTP Listening
	if err := e.Start(":" + c.AppConfig.Port); err != nil && err != http.ErrServerClosed {
		log.Panicf("error - [main.New] unable to start server: %v", err)
	}
}
