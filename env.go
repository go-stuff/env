package env

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// File parses a key, value pair file as environment variables.
func File(path string) error {

	// dose the file exist
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Printf("INFO > env.go > File(): %s does not exist\n", path)
	} else {
		// open .env
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// read each line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if scanner.Text() != "" {
				regex := regexp.MustCompile(`([a-zA-Z0-9-_]*)\s*=\s*"(.*)"`)
				matches := regex.FindStringSubmatch(scanner.Text())
				if len(matches) != 3 {
					return fmt.Errorf("error in %s", path)
				}

				// []matches = [0]line, [1](group1), [2](group2)
				err := os.Setenv(matches[1], matches[2])
				if err != nil {
					return err
				}

				// DO NOT PRINT ENVIRONMENT VARIABLES OUT, IT IS UNSAFE
				// fmt.Printf("env:%s = %s\n", matches[1], os.Getenv(matches[1]))
			}
		}

		// catch scanner errors
		err = scanner.Err()
		if err != nil {
			return err
		}

		log.Printf("INFO > env.go > File(): %s loaded\n", path)
	}

	return nil
}
