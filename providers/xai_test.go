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

func TestXAI_FileTypeNotSupported(t *testing.T) {
	logger := testutils.NewTestLogger(t)
	p := &XAI{} // nil client is sufficient to exercise early validation

	runCfg := config.RunConfig{Name: "test-run", Model: "grok-test"}
	task := config.Task{
		Name:  "bad_file_type",
		Files: []config.TaskFile{mockTaskFile(t, "file.txt", "file://file.txt", "text/plain")},
	}
	_, err := p.Run(context.Background(), logger, runCfg, task)
	require.ErrorIs(t, err, ErrFileNotSupported)
}
