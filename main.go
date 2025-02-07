package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	go func() {
		<-signals
		log.Println("Receive SIGTERM, Sleep 10s")
		time.Sleep(10 * time.Second)
		log.Println("Exit")
		os.Exit(0)
	}()

	http.HandleFunc("/pre-stop", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Exec preStop Hook, Sleep 10s")
		time.Sleep(10 * time.Second)
		log.Println("pre-Stop Hook successfully")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
