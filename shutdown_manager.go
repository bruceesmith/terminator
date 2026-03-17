// Copyright © 2026 Bruce Smith <bruceesmith@gmail.com>
// Use of this source code is governed by the MIT
// License that can be found in the LICENSE file.

// Acknowledgement to https://medium.com/@yaninyzwitty (Deterministic Shutdown Flows)

package terminator

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/bruceesmith/logger"
)

// ShutdownManager orchestrates ordered, phased shutdown.
type ShutdownManager struct {
	mu      sync.Mutex
	phases  []shutdownPhase
	timeout time.Duration
}

type shutdownPhase struct {
	name    string
	timeout time.Duration
	fn      func(ctx context.Context) error
}

func NewShutdownManager(total time.Duration) *ShutdownManager {
	return &ShutdownManager{timeout: total}
}

// Register adds a shutdown function to a named phase.
// Phases execute in registration order.
func (s *ShutdownManager) Register(name string, timeout time.Duration, fn func(ctx context.Context) error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.phases = append(s.phases, shutdownPhase{name, timeout, fn})
}

// Shutdown runs all phases sequentially with independent timeouts.
func (s *ShutdownManager) Shutdown() error {
	s.mu.Lock()
	phases := s.phases
	s.mu.Unlock()

	var errs []error
	for _, phase := range phases {
		ctx, cancel := context.WithTimeout(context.Background(), phase.timeout)

		if err := phase.fn(ctx); err != nil {
			logger.Error("shutdown phase failed", "phase", phase.name, "error", err)
			errs = append(errs, fmt.Errorf("phase %s: %w", phase.name, err))
		} else {
			logger.Info("shutdown phase complete", "phase", phase.name)
		}

		cancel()
	}

	return errors.Join(errs...)
}
