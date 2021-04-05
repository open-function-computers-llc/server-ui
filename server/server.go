package server

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	rice "github.com/GeertJohan/go.rice"
	"github.com/dchest/uniuri"
	_ "github.com/go-sql-driver/mysql" // db connections via mysql/mariaDB
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3" // db connections via sqlite
	"github.com/open-function-computers-llc/server-ui/db"
	"github.com/open-function-computers-llc/server-ui/models"
	"github.com/open-function-computers-llc/server-ui/site"
	"github.com/sirupsen/logrus"
)

// Server this is a web server
type Server struct {
	db           *sqlx.DB              // DB connection
	logger       *logrus.Logger        // logger
	router       *httprouter.Router    // router for all routes
	routePrefix  string                // used to obfuscate routes
	assets       *rice.Box             // rice box for static assets
	publicAssets *rice.Box             // rice box for public assets
	sessions     *sessions.CookieStore //valid user sessions
	sites        []site.Site           // the websites on this server
	hostType     string                // what type of sites are intended to be ran on this server (modx|wordpress)
}

func (s *Server) bootstrap() error {
	// verify host type
	s.hostType = os.Getenv("APP_HOST_TYPE")
	err := verifyHostType(s.hostType)
	if err != nil {
		return errors.New("Invalid host type. Check your env var for APP_HOST_TYPE")
	}

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
	prefix := os.Getenv("APP_ROUTE_PREFIX")
	if prefix == "" {
		prefix = uniuri.New()
	}
	s.routePrefix = "/" + prefix
	s.Log("Route prefix: ", s.routePrefix)

	// load in static assets
	s.assets = rice.MustFindBox("../resources/views")
	s.Log("rice box template assets loaded")

	// load in public assets
	s.publicAssets = rice.MustFindBox("../public")
	s.Log("rice box public assets loaded")

	// set up encrypted cookies for sessions
	encLen, err := strconv.Atoi(os.Getenv("SESSION_LENGTH"))
	if err != nil {
		return errors.New("Invalid session length. Check your env var for SESSION_LENGTH.")
	}
	if encLen < 1 {
		return errors.New("Invalid session length. Check your env var for SESSION_LENGTH.")
	}
	encKey := os.Getenv("SESSION_ENC_KEY")
	if encKey == "" {
		encKey = uniuri.New()
	}
	s.sessions = sessions.NewCookieStore([]byte(encKey))
	s.sessions.Options = &sessions.Options{
		MaxAge: 60 * encLen, // 30 minute sessions
	}

	// set up routes
	s.router = httprouter.New()

	// bind http handler routes
	s.bindRoutes()

	// add in static assets that are cool to serve to anyone
	assetServer := http.FileServer(s.publicAssets.HTTPBox())
	s.router.NotFound = assetServer

	// populate the sites for this webserver
	err = s.loadSites()
	if err != nil {
		return err
	}

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

// NewServer - create and bootstrap a new server struct
func NewServer() (Server, error) {
	s := Server{}
	err := s.bootstrap()
	return s, err
}
