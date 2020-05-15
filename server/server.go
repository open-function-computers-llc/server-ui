package server

import (
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // db connections via mysql/mariaDB
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3" // db connections via sqlite
	"github.com/open-function-computers-llc/server-ui/db"
	"github.com/open-function-computers-llc/server-ui/models"
	"github.com/sirupsen/logrus"
)

// Server this is a web server
type Server struct {
	db     *sqlx.DB
	logger *logrus.Logger
	router *httprouter.Router
}

func (s *Server) bootstrap() error {
	// attach logger
	if s.logger == nil {
		s.logger = logrus.New()
		s.logger.Out = os.Stdout
	}

	// bootstrap database connection
	if s.db == nil {
		err := s.connectToDB()
		if err != nil {
			return err
		}
	}

	// attach routes
	s.router = httprouter.New()
	s.bindRoutes()

	return nil
}

// Start will start up the server instance and bootstrap it if necessary
func (s *Server) Start() error {
	if s.logger == nil {
		err := s.bootstrap()
		if err != nil {
			return err
		}
	}

	// listen and serve
	s.Log("Started server on port " + os.Getenv("APP_PORT"))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), s.router)

	return nil
}

func (s *Server) connectToDB() error {
	dbType := os.Getenv("DB_TYPE")

	// default to sqlite
	if dbType == "" || dbType == "sqlite" {
		return s.setUpSQLite()
	}

	return s.setUpMariaDB()
}

func (s *Server) setUpSQLite() error {
	s.Log("connection set to SQLite")
	dbName := os.Getenv("DB_DATABASE")
	if dbName == "" {
		dbName = ":memory:"
	}
	db, err := sqlx.Connect("sqlite3", dbName)
	if err != nil {
		return err
	}
	s.db = db
	s.migrateDB()
	return nil
}

func (s *Server) setUpMariaDB() error {
	s.Log("connection set to MariaDB")
	dsn, err := db.GenerateMariaDSN(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	if err != nil {
		return err
	}

	// try to connect to DB
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	s.db = db
	s.Log("Database configured with DSN: ", dsn)

	return nil
}

func (s *Server) migrateDB() {
	s.Log("Migrating Database")
	if s.db.DriverName() != "sqlite3" {
		return
	}

	s.db.MustExec(models.GetSchemas("sqlite3"))
	s.Log("DB Migrated")
}

// Log will log any messages to the attached logger instance
func (s *Server) Log(messages ...string) {
	s.logger.Info(messages)
}

// LogError will log any messages to the attached logger instance
func (s *Server) LogError(messages ...string) {
	s.logger.Error(messages)
}
