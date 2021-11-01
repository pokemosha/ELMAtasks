package main

import (
	"ELMAcourses/config"
	"ELMAcourses/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go func() {
		err := http.ListenAndServe(
			config.AddrUsers,
			services.CreateServer(),
		)

		if err != nil {
			panic(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-interrupt
}
