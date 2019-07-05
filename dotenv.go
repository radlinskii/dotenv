/*

Package dotenv is a simple API for reading environment variables stored in files inside your project.

Parsing rules:

1. Variables should be stored as Key-Value pairs separated by "=" sign.  
2. Whitespaces around key and value are trimmed.  
3. Lines starting with "#" sign are omitted.  
4. Blank lines are omitted.  
5. Environment variables that are already exported are not getting overwritten by those read from .env file.  

*/
package dotenv

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// SetEnv sets env variables specified in the .env file in cwd.
func SetEnv() error {
	return SetEnvFromPath(".env")
}

// SetEnvFromPath sets env variables specified in the file in given path.
func SetEnvFromPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("file: %q does not exist\n", path)
		return nil
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if line == "" || strings.HasPrefix(strings.TrimSpace(line), "#") {
			continue
		}

		varinfo := strings.SplitN(line, "=", 2)
		if len(varinfo) != 2 {
			return errors.New("Error parsing " + path + " file")
		}

		key := strings.TrimSpace(varinfo[0])
		val := strings.TrimSpace(varinfo[1])

		if _, ok := os.LookupEnv(key); !ok {
			err := os.Setenv(key, val)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
