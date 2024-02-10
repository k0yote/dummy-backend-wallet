package util

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func CurrentDir() string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get current directory")
	}

	for {
		_, err := os.ReadFile(filepath.Join(currentDir, "go.mod"))
		if os.IsNotExist(err) {
			if currentDir == filepath.Dir(currentDir) {
				// at the root
				break
			}
			currentDir = filepath.Dir(currentDir)
			continue
		} else if err != nil {
			log.Fatal().Err(err).Msg("failed to read go.mod")
		}
		break
	}

	return currentDir
}
