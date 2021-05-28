package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string, registerHandlers func()) (context.Context, error) {
	registerHandlers()
	ctx = startService(ctx, serviceName, host, port)
	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	srv := http.Server{Addr: ":" + port}
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()
	go func() {
		fmt.Printf("%v startd. Press any key to stop. \n", serviceName)
		var s string
		_, _ = fmt.Scanln(&s)
		_ = srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
