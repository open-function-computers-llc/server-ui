package server

import (
	"os"
)

func createDummyServer() Server {
	os.Setenv("DB_TYPE", "sqlite")
	s := Server{}
	s.bootstrap()

	return s
}
