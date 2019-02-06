package worldconfig

import (
	"bufio"
	"os"
	"strings"
)

const (
	BACKEND_SQLITE3  string = "sqlite3"
	BACKEND_FILES    string = "files"
	BACKEND_POSTGRES string = "postgresql"
)

const (
	CONFIG_BACKEND                string = "backend"
	CONFIG_PLAYER_BACKEND         string = "player_backend"
	CONFIG_PSQL_CONNECTION        string = "pgsql_connection"
	CONFIG_PSQL_MAPSERVER         string = "pgsql_mapserver_connection"
)

type PsqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

type WorldConfig struct {
	Backend       string
	PlayerBackend string

	PsqlConnection       string
	MapObjectConnection  string
}


func Parse(filename string) WorldConfig {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cfg := WorldConfig{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sepIndex := strings.Index(line, "=")
		if sepIndex < 0 {
			continue
		}

		valueStr := strings.Trim(line[sepIndex+1:], " ")
		keyStr := strings.Trim(line[:sepIndex], " ")

		switch keyStr {
		case CONFIG_BACKEND:
			cfg.Backend = valueStr
		case CONFIG_PLAYER_BACKEND:
			cfg.PlayerBackend = valueStr
		case CONFIG_PSQL_CONNECTION:
			cfg.PsqlConnection = valueStr
		case CONFIG_PSQL_MAPSERVER:
			cfg.MapObjectConnection = valueStr
		}

	}

	return cfg
}
