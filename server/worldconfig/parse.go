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
	CONFIG_BACKEND         string = "backend"
	CONFIG_PLAYER_BACKEND  string = "player_backend"
	CONFIG_PSQL_CONNECTION string = "pgsql_connection"
	CONFIG_PSQL_MAPSERVER  string = "pgsql_mapserver_connection"
)

func Parse(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cfg := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sepIndex := strings.Index(line, "=")
		if sepIndex < 0 {
			continue
		}

		valueStr := strings.Trim(line[sepIndex+1:], " ")
		keyStr := strings.Trim(line[:sepIndex], " ")

		cfg[keyStr] = valueStr
	}

	return cfg
}
