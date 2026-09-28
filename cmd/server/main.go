package main

import (
	"auth-service/internal/config"
	"auth-service/internal/router"
	"auth-service/internal/shared/clients"
	"auth-service/internal/shared/database"
	"auth-service/internal/shared/middleware"
	"auth-service/internal/shared/otp"
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//Load env
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Info("ERR: ", "error", err.Error())
		return
	}

	//connect and initiate postgres db and defer close
	db, err := database.ConnectDB(cfg.DB_URL)
	if err != nil {
		slog.Error("DB ERROR: ", "error", err.Error())
		return
	}
	slog.Info("Database connected successfully")
	defer db.Close()

	// connect and initialize redis db and defer close
	rdb, err := clients.CreateNewRedisClient(cfg.RedisUrl, "", 0)
	if err != nil {
		log.Fatalf("failed to initialize redis, %v\n", err)
	}
	slog.Info("redis connected")
	defer rdb.Close()

	// connect and initialize resend client and defer close
	emailSender := clients.CreateResendClient(cfg.ResendApiKey)

	// initialize shared otp store
	otpStore := otp.NewOTPStore(rdb)

	// create serve mux
	mux := http.NewServeMux()

	router.RegisterRouter(mux, db, cfg, rdb, emailSender, otpStore)
	server := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      middleware.CorsConfig(middleware.Logger(mux)),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go startServer(&server)
	sig := <-sigChan

	log.Printf("Recieved signal %v\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown, %v\n", err)
		return
	}
	log.Printf("Server shutdown gracefully. Goodbye!!")
}

func startServer(server *http.Server) {
	slog.Info("Server started", "host", "localhost", "port", "2110")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
