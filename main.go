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

	port := fmt.Sprintf(":%d", setting.Cfg.GetInt("server.port"))

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    setting.Cfg.GetDuration("server.read_timeout") * time.Second,
		WriteTimeout:   setting.Cfg.GetDuration("server.write_timeout") * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http server listening %s", port)
	
	go func() {
        if err := s.ListenAndServe(); err != nil {
            log.Printf("[info] Listen: %s\n", err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <- quit

    log.Println("[info] Shutdown Server ...")

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    if err := s.Shutdown(ctx); err != nil {
        log.Fatal("[info] Server Shutdown:", err)
    }

    log.Println("[info] Server exiting")
}