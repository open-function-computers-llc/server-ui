package db

import "errors"

// GenerateMariaDSN return a string that is formatted correctly for a
// MariaDB/MySQL db connection
func GenerateMariaDSN(user, pass, db, host, port string) (string, error) {
	if user == "" || pass == "" || db == "" || host == "" || port == "" {
		return "", errors.New("All fields are required. Please check your settings and try again")
	}
	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	return dsn, nil
}
