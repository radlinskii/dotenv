package dotenv

/*

Parsing rules:

1. BASIC=basic
2. WHITE_SPACES = are trimmed
3. # lines starting with "#"" are omitted
4. # blank lines are omitted
5. ALREADY_EXPORTED_VARIABLES="are not overwritten"

*/

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// SetEnv sets env variables specified in the .env file in cwd.
func SetEnv() {
	SetEnvFromPath(".env")
}

// SetEnvFromPath sets env variables specified in the file in given path.
func SetEnvFromPath(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		varinfo := strings.SplitN(line, "=", 2)
		if len(varinfo) != 2 {
			log.Fatalln("Error parsing " + path + " file")
		}

		key := strings.TrimSpace(varinfo[0])
		val := strings.TrimSpace(varinfo[1])

		if os.Getenv(key) == "" {
			err := os.Setenv(key, val)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
