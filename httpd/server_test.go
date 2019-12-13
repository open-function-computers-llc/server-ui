package main

import (
	"os"
)

func createDummyServer() server {
	os.Setenv("DB_TYPE", "sqlite")
	s := server{}
	s.bootstrap()

	return s
}
