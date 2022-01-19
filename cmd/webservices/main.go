package main

import (
	"fmt"
	"mocking-goway/internal/appserver"
	"mocking-goway/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	/*
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		c := &config.Config{Port: "8080"}
		server := appserver.NewAppServer(c)
		server.Start()

		<-sigs
		fmt.Println("Done")

	*/
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	c := &config.Config{Port: ":8080"}
	server := appserver.NewAppServer(c)
	go server.Start()

	//go appserver.Start()
	<-sigs
	fmt.Println("Done")
}
