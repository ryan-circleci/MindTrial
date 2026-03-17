// Copyright (c) 2025 Petr Malik and CircleCI, Inc.
// SPDX-License-Identifier: MIT

package formatters

import (
	"testing"

	"github.com/CircleCI-Research/MindTrial/runners"
	"github.com/stretchr/testify/assert"
)

func TestSummaryLogFormatterWrite(t *testing.T) {
	tests := []struct {
		name    string
		results runners.Results
		want    string
	}{
		{
			name:    "format no results",
			results: runners.Results{},
			want:    "testdata/empty.summary.log",
		},
		{
			name:    "format some results",
			results: mockResults,
			want:    "testdata/results.summary.log",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := NewSummaryLogFormatter()
			assertFormatterOutputFromFile(t, formatter, tt.results, tt.want)
		})
	}
}

func TestSummaryLogFormatterFileExt(t *testing.T) {
	formatter := NewSummaryLogFormatter()
	assert.Equal(t, "summary.log", formatter.FileExt())
}
