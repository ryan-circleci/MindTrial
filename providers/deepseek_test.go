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

func TestDeepseek_FileUploadNotSupported(t *testing.T) {
	logger := testutils.NewTestLogger(t)
	p := &Deepseek{} // nil client is sufficient to test early error
	runCfg := config.RunConfig{Name: "test-run", Model: "directional"}
	task := config.Task{
		Name:  "with_file",
		Files: []config.TaskFile{mockTaskFile(t, "img.png", "file://img.png", "image/png")},
	}
	_, err := p.Run(context.Background(), logger, runCfg, task)
	require.ErrorIs(t, err, ErrFileUploadNotSupported)
}
