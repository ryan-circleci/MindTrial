// Copyright (c) 2025 Petr Malik and CircleCI, Inc.
// SPDX-License-Identifier: MIT

// Package formatters provides output formatting functionality for MindTrial results.
// It supports multiple output formats including HTML, CSV, and text logs.
package formatters

import (
	"errors"
	"io"
	"strings"

	"github.com/CircleCI-Research/MindTrial/pkg/utils"
	"github.com/CircleCI-Research/MindTrial/runners"
)

// ErrPrintResults indicates that result formatting failed.
var ErrPrintResults = errors.New("failed to print formatted results")

// Formatter handles converting results into specific output formats.
type Formatter interface {
	// FileExt returns the formatter's file extension.
	FileExt() string
	// Write outputs formatted results to the writer.
	Write(results runners.Results, out io.Writer) error
}

// formatAnswerText returns a single plain text block containing all possible formatted answers for
// CSV and other text-based outputs. If there is only one answer, it returns the answer directly.
// If there are multiple answers, it returns them as a bracket-formatted list.
func formatAnswerText(result runners.RunResult) string {
	answers := FormatAnswer(result, false)
	if len(answers) == 1 {
		return answers[0]
	}

	// Indent all lines in each answer.
	indentedAnswers := make([]string, len(answers))
	for i, answer := range answers {
		lines := utils.SplitLines(answer)
		for j, line := range lines {
			lines[j] = "    " + line
		}
		indentedAnswers[i] = strings.Join(lines, "\n")
	}

	return "[\n" + strings.Join(indentedAnswers, "\n  ,\n") + "\n]"
}
