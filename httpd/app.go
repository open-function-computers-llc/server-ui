package main

import "github.com/open-function-computers-llc/server-ui/server"

func runApp() error {
	s := server.Server{}
	err := s.Start()
	if err != nil {
		return err
	}
	return nil
}
