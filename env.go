package env

import (
	"bufio"
	"os"
	"regexp"

	"go.uber.org/zap"
)

// File parses a key, value pair file as environment variables.
func File(path string) error {

	// initialize logging
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer logger.Sync()

	// if the path does not exist
	// log a warning and use environment variables
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		logger.Warn(
			"path does not exist",
			zap.String("path", path),
			zap.String("err", err.Error()),
		)
		return nil
	}

	// open environment file (example: .env)
	file, err := os.Open(path)
	if err != nil {
		logger.Error(
			"unable to open path",
			zap.String("err", err.Error()),
		)
		return err
	}
	defer file.Close()

	// read each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			regex := regexp.MustCompile(`^\s*([a-zA-Z0-9-_]*)\s*=\s*"(.*)"\s*$`)
			matches := regex.FindStringSubmatch(scanner.Text())

			if len(matches) != 3 {
				logger.Error(
					"unable to parse",
					zap.String("line", scanner.Text()),
					zap.String("err", err.Error()),
				)
				return err
			}

			// []matches = [0]line, [1](group1), [2](group2)
			err := os.Setenv(matches[1], matches[2])
			if err != nil {
				logger.Error(
					"unable to set environment variable",
					zap.String("line", scanner.Text()),
					zap.String("err", err.Error()),
				)
				return err
			}

			// DO NOT PRINT ENVIRONMENT VARIABLES OUT, IT IS UNSAFE
			// fmt.Printf("env:%s = %s\n", matches[1], os.Getenv(matches[1]))
		}
	}

	// catch scanner errors
	err = scanner.Err()
	if err != nil {
		logger.Error(
			"errors with file",
			zap.String("err", err.Error()),
		)
		return err
	}

	logger.Info(
		"environment file was loaded",
		zap.String("path", path),
	)
	return nil
}
