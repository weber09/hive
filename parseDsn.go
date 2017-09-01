package hive

import "strings"

type Config struct {
hiveType string
hostName string
port string
dbName string
user string
password string
args map[string]interface{}
}

//jdbc:hive2://localhost:10000/default

func ParseDsn (dsn string) (*Config, error) {
	splDsn := strings.Split(dsn, "//")
	pathPart := strings.Split(splDsn[1], "/")
	hostPort := strings.Split(pathPart[0], ":")
	dbName := pathPart[1]

	config := new(Config)

	config.hostName = hostPort[0]
	config.port = hostPort[1]

	config.dbName = dbName

	return config, nil
}
