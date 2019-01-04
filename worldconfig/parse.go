package worldconfig

import (
	"bufio"
	"os"
	"strings"
)

const (
	BACKEND_SQLITE3 string = "sqlite3"
	BACKEND_FILES string = "files"
	BACKEND_POSTGRES string = "postgresql"
)

const (
	CONFIG_BACKEND string = "backend"
	CONFIG_PLAYER_BACKEND string = "player_backend"
	CONFIG_PSQL_CONNECTION string = "pgsql_connection"
	CONFIG_PSQL_PLAYER_CONNECTION string = "pgsql_player_connection"
)

type PsqlConfig struct {
	Host string
	Port int
	Username string
	Password string
	DbName string
}

type WorldConfig struct {
	Backend       string
	PlayerBackend string

	PsqlConnection *PsqlConfig
	PsqlPlayerConnection *PsqlConfig
}

func parseConnectionString(str string) *PsqlConfig {
	return &PsqlConfig{}
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
		sc := bufio.NewScanner(strings.NewReader(scanner.Text()))
		sc.Split(bufio.ScanWords)
		lastPart := ""
		for sc.Scan() {
			switch (lastPart) {
			case CONFIG_BACKEND:
				cfg.Backend = sc.Text()
			case CONFIG_PLAYER_BACKEND:
				cfg.PlayerBackend = sc.Text()
			case CONFIG_PSQL_CONNECTION:
				cfg.PsqlConnection = parseConnectionString(sc.Text())
			case CONFIG_PSQL_PLAYER_CONNECTION:
				cfg.PsqlPlayerConnection = parseConnectionString(sc.Text())
			}

			if sc.Text() != "=" {
				lastPart = sc.Text()
			}
		}
	}

	return cfg
}
