package main

import (
	"fmt"
	"net/http"

	"context"
	"log"
	"os"
    "os/signal"
    "time"
	
	"gin-starter/routers"
	"gin-starter/pkg/setting"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Cfg.GetInt("server.port")),
		Handler:        router,
		ReadTimeout:    setting.Cfg.GetDuration("server.read_timeout"),
		WriteTimeout:   setting.Cfg.GetDuration("server.write_timeout"),
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
        if err := s.ListenAndServe(); err != nil {
            log.Printf("Listen: %s\n", err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <- quit

    log.Println("Shutdown Server ...")

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    if err := s.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }

    log.Println("Server exiting")
}