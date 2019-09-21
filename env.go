package env

import (
	"bufio"
	"os"
	"regexp"

	"go.uber.org/zap"
)

// File parses a key, value pair file as environment variables.
func File(path string, logger *zap.Logger) {
	// if the path does not exist
	// log a warning and use environment variables
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		logger.Warn(
			"path does not exist, using environment variables",
			zap.String("path", path),
		)
		return
	}

	// check if we are dealing with a file or directory
	if fi.IsDir() {
		logger.Fatal(
			"path is a directory not a file",
			zap.String("path", path),
		)
	}

	// open environment file (example: .env)
	file, err := os.Open(path)
	if err != nil {
		logger.Fatal(
			"cannot open path for reading",
			zap.String("path", path),
			zap.String("err", err.Error()),
		)
	}
	defer file.Close()

	// read each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			regex := regexp.MustCompile(`^\s*([a-zA-Z0-9-_]*)\s*=\s*"(.*)"\s*$`)
			matches := regex.FindStringSubmatch(scanner.Text())

			// if matches is not 3 it does not follow the structure of a variable file
			// match 1 must be an unquoted variable name: VARIABLE_NAME
			// match 2 must be an equals sign: =
			// match 3 must be a quoted value: "this is my value"
			if len(matches) != 3 {
				logger.Fatal(
					"unable to parse",
					zap.String("line", scanner.Text()),
				)
			}

			// []matches = [0]line, [1](group1), [2](group2)
			err := os.Setenv(matches[1], matches[2])
			if err != nil {
				logger.Fatal(
					"unable to set environment variable",
					zap.String("line", scanner.Text()),
					zap.String("err", err.Error()),
				)
			}

			// DO NOT PRINT ENVIRONMENT VARIABLES OUT, IT IS UNSAFE
			// fmt.Printf("env:%s = %s\n", matches[1], os.Getenv(matches[1]))
		}
	}

	// catch scanner errors
	err = scanner.Err()
	if err != nil {
		logger.Fatal(
			"error scanning file",
			zap.String("err", err.Error()),
		)
	}

	logger.Info(
		"environment file was loaded",
		zap.String("path", path),
	)
}
