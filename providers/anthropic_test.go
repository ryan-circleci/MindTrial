// Copyright (c) 2025 Petr Malik and CircleCI, Inc.
// SPDX-License-Identifier: MIT

package providers

import (
	"context"
	"testing"

	"github.com/CircleCI-Research/MindTrial/config"
	"github.com/CircleCI-Research/MindTrial/pkg/testutils"
	"github.com/stretchr/testify/require"
)

func TestAnthropic_FileTypeNotSupported(t *testing.T) {
	logger := testutils.NewTestLogger(t)
	p := &Anthropic{} // nil client is sufficient since error occurs before any API call

	runCfg := config.RunConfig{Name: "test-run", Model: "claude"}
	task := config.Task{
		Name:  "bad_file_type",
		Files: []config.TaskFile{mockTaskFile(t, "file.txt", "file://file.txt", "text/plain")},
	}
	_, err := p.Run(context.Background(), logger, runCfg, task)
	require.ErrorIs(t, err, ErrFileNotSupported)
}
