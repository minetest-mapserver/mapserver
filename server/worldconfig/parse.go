package worldconfig

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	CONFIG_PSQL_PLAYER_CONNECTION string = "pgsql_player_connection"
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

	PsqlConnection       *PsqlConfig
	PsqlPlayerConnection *PsqlConfig
}

func parseConnectionString(str string) *PsqlConfig {
	cfg := PsqlConfig{}

	pairs := strings.Split(str, " ")
	for _, pair := range pairs {
		fmt.Println(pair)
		kv := strings.Split(pair, "=")
		switch kv[0] {
		case "host":
			cfg.Host = kv[1]
		case "port":
			cfg.Port, _ = strconv.Atoi(kv[1])
		case "user":
			cfg.Host = kv[1]
		case "password":
			cfg.Password = kv[1]
		case "dbname":
			cfg.DbName = kv[1]
		}
	}

	return &cfg
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
			cfg.PsqlConnection = parseConnectionString(valueStr)
		case CONFIG_PSQL_PLAYER_CONNECTION:
			cfg.PsqlPlayerConnection = parseConnectionString(valueStr)
		}

	}

	return cfg
}
