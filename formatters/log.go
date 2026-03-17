// Copyright (c) 2025 Petr Malik and CircleCI, Inc.
// SPDX-License-Identifier: MIT

package formatters

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/CircleCI-Research/MindTrial/runners"
)

// NewLogFormatter creates a new formatter that outputs detailed results as an ASCII table.
func NewLogFormatter() Formatter {
	return &logFormatter{}
}

type logFormatter struct{}

func (f logFormatter) FileExt() string {
	return "log"
}

func (f logFormatter) Write(results runners.Results, out io.Writer) error {
	tab := tabwriter.NewWriter(out, 0, 0, 1, ' ', tabwriter.Debug)
	defer tab.Flush()
	if _, err := fmt.Fprintln(tab, "TraceID\tProvider\tRun\tTask\tStatus\tScore\tDuration\tAnswer\t"); err != nil {
		return fmt.Errorf("%w: %v", ErrPrintResults, err)
	}

	return ForEachOrdered(results, func(_ string, runResults []runners.RunResult) error {
		for _, result := range runResults {
			if _, err := fmt.Fprintf(tab, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n", result.TraceID, result.Provider, result.Run, result.Task, ToStatus(result.Kind), Score(result), RoundToMS(result.Duration), formatAnswerText(result)); err != nil {
				return fmt.Errorf("%w: %v", ErrPrintResults, err)
			}
		}
		return nil
	})
}
