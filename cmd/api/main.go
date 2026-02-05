package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	router:=http.NewServeMux()
	server:=http.Server{
		Addr: ":3480",
		Handler: router,
	}
	out:=make(chan os.Signal, 1)
	log.Println("Listening on port 3480")
	signal.Notify(out,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)
	go func() {
		err:=server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()
	<-out
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	err:=server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Error in shuting down")
	}
}