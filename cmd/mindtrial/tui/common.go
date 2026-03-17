// Copyright (c) 2025 Petr Malik and CircleCI, Inc.
// SPDX-License-Identifier: MIT

// Package tui provides terminal-based UI for MindTrial CLI.
package tui

import (
	"errors"
	"os"

	"github.com/charmbracelet/x/term"
)

const (
	highlightColor  = "170" // pink/magenta
	helpTextColor   = "240" // gray
	initializingMsg = "Initializing UI, please wait..."
)

var (
	// ErrInteractiveMode is returned when interactive mode encounters an error.
	ErrInteractiveMode = errors.New("interactive mode error")

	// ErrTerminalRequired is returned when interactive mode is attempted
	// in a non-terminal environment.
	ErrTerminalRequired = errors.New("interactive mode requires a terminal environment")
)

// UserInputEvent represents the type of user actions in an interactive session.
type UserInputEvent int

const (
	// Exit indicates that the user wants to exit the application.
	Exit UserInputEvent = iota
	// Quit indicates that the user wants to quit the current interactive screen while continuing execution.
	Quit
	// Continue indicates that the user wants to proceed with the current selections.
	Continue
)

// IsTerminal reports whether the current output is connected to a terminal.
func IsTerminal() bool {
	return term.IsTerminal(os.Stdout.Fd())
}
