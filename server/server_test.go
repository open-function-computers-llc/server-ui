package server

import (
	"os"
)

func createDummyServer() Server {
	os.Setenv("DB_TYPE", "sqlite")
	os.Setenv("ROUTE_PREFIX", "test")
	s := Server{}
	s.bootstrap()

	return s
}
