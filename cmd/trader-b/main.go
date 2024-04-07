// Package main provides the entry point for the application.
package main

import (
	"github.com/twk/trader-b/cmd/trader-b/commands"
	"go.uber.org/zap"

	"github.com/twk/trader-b/internal/logger"
)

func main() {
	log := logger.NewLogger(nil)

	cmd, err := commands.NewRootCommand(log)
	if err != nil {
		log.Fatal("Failed to create root command", zap.Error(err))
	}

	err = cmd.Execute()
	if err != nil {
		log.Fatal("Failed to execute command", zap.Error(err))
	}

	log.Info("Command executed successfully")
}
