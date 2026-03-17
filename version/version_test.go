// Copyright (c) 2025 Petr Malik and CircleCI, Inc.
// SPDX-License-Identifier: MIT

package version

import (
	"runtime/debug"
	"sync"
	"testing"

	"github.com/CircleCI-Research/MindTrial/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

var sourceLock sync.Mutex

func TestName(t *testing.T) {
	assert.Equal(t, "MindTrial", Name)
}

func TestGetVersion(t *testing.T) {
	testutils.SyncCall(&sourceLock, func() {
		originalSource := source
		source = func() debug.Module {
			return debug.Module{
				Version: "driver",
			}
		}
		defer func() { source = originalSource }()
		assert.Equal(t, "driver", GetVersion())
	})
}

func TestGetSource(t *testing.T) {
	testutils.SyncCall(&sourceLock, func() {
		originalSource := source
		source = func() debug.Module {
			return debug.Module{
				Path: "neat-thread.com/necessary-baggy/unaware-polenta",
			}
		}
		defer func() { source = originalSource }()
		assert.Equal(t, "neat-thread.com/necessary-baggy/unaware-polenta", GetSource())
	})
}
