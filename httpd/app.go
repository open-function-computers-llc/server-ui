package main

import (
	"github.com/joho/godotenv"
	"github.com/open-function-computers-llc/server-ui/server"
)

func runApp() error {
	// load in ENV
	err := godotenv.Load()
	if err != nil {
		return err
	}

	// instantiate a server
	s, err := server.NewServer()
	if err != nil {
		return err
	}

	// start the new server instance
	err = s.Start()
	if err != nil {
		return err
	}
	return nil
}
