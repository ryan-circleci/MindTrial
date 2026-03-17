// Copyright (c) 2025 Petr Malik and CircleCI, Inc.
// SPDX-License-Identifier: MIT

package formatters

import (
	"testing"

	"github.com/CircleCI-Research/MindTrial/runners"
	"github.com/stretchr/testify/assert"
)

func TestLogFormatterWrite(t *testing.T) {
	tests := []struct {
		name    string
		results runners.Results
		want    string
	}{
		{
			name:    "format no results",
			results: runners.Results{},
			want:    "testdata/empty.log",
		},
		{
			name:    "format some results",
			results: mockResults,
			want:    "testdata/results.log",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := NewLogFormatter()
			assertFormatterOutputFromFile(t, formatter, tt.results, tt.want)
		})
	}
}

func TestLogFormatterFileExt(t *testing.T) {
	formatter := NewLogFormatter()
	assert.Equal(t, "log", formatter.FileExt())
}
