package main

import (
	"context"
	"library/pkg/db"
	"library/pkg/handlers"
	"library/pkg/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	conn, err := db.InitDB("")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repo := db.New(conn)
	serv := service.New(repo)
	hand := handlers.New(serv)

	r.POST("/", hand.Create)
	r.POST("/:book_name/", hand.AddPage)
	r.GET("/", hand.GetByTitle)
	r.GET("/", hand.GetAll)
	r.GET("/:book_name/read", hand.GetPage)
	r.DELETE("/:book_name", hand.Delete)

	srv := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
