package server

import (
	"net/http"
	"os"

	rice "github.com/GeertJohan/go.rice"
	"github.com/dchest/uniuri"
	_ "github.com/go-sql-driver/mysql" // db connections via mysql/mariaDB
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3" // db connections via sqlite
	"github.com/open-function-computers-llc/server-ui/db"
	"github.com/open-function-computers-llc/server-ui/models"
	"github.com/sirupsen/logrus"
)

// Server this is a web server
type Server struct {
	db          *sqlx.DB              // DB connection
	logger      *logrus.Logger        // logger
	router      *httprouter.Router    // router for all routes
	routePrefix string                // used to obfuscate routes
	assets      *rice.Box             // rice box for static assets
	sessions    *sessions.CookieStore //valid user sessions
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

	// set the route prefix
	prefix := os.Getenv("ROUTE_PREFIX")
	if prefix == "" {
		prefix = uniuri.New()
	}
	s.routePrefix = "/" + prefix
	s.Log("Route prefix: ", s.routePrefix)

	// load in static assets
	s.assets = rice.MustFindBox("../resources/views")
	s.Log("rice box assets loaded")

	// set up encrypted cookies for sessions
	s.sessions = sessions.NewCookieStore([]byte(uniuri.NewLen(64)))

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
func (s *Server) Log(messages ...interface{}) {
	s.logger.Info(messages...)
}

// LogError will log any messages to the attached logger instance
func (s *Server) LogError(messages ...string) {
	s.logger.Error(messages)
}
