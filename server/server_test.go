package server

import (
	"os"
)

func createDummyServer() Server {
	os.Setenv("DB_TYPE", "sqlite")
	os.Setenv("ROUTE_PREFIX", "test")
	os.Setenv("APP_HOST_TYPE", "modx")
	os.Setenv("SESSION_LENGTH", "60")
	s := Server{}
	err := s.bootstrap()
	if err != nil {
		panic(err)
	}

	return s
}
