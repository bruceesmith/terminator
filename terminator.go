// Copyright © 2024 Bruce Smith <bruceesmith@gmail.com>
// Use of this source code is governed by the MIT
// License that can be found in the LICENSE file.

/*
Package terminator permits orderly stopping / shutdown of a group of goroutines via methods which mimic stop
of a [sync.WaitGroup].There is a default Terminator accessible through top level functions (Add, Done, Wait and
so on) that call the corresponding Terminator methods
*/
package terminator

//go:generate ./make_doc.sh

import (
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
)

// Terminator manages groups of goroutines
type Terminator struct {
	shutdown     chan struct{}
	shuttingDown bool
	wg           sync.WaitGroup
}

var (
	defaultTerminator atomic.Pointer[Terminator]
)

func init() {
	defaultTerminator.Store(
		&Terminator{
			shutdown: make(chan struct{}, 1),
		},
	)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		defer signal.Stop(c)
		<-c
		Stop()
	}()
}

// Add adds delta to the count of goroutines in the group
func Add(delta int) {
	Default().wg.Add(delta)
}

// Default returns the default [Terminator].
func Default() *Terminator { return defaultTerminator.Load() }

// Go runs a function inside an Add(1) --- Done() sequence
func Go(f func()) {
	Default().wg.Add(1)
	go func() {
		defer Default().wg.Done()
		f()
	}()
}

// Done decrements the count of goroutines in the group by one
func Done() {
	Default().wg.Done()
}

// ShutDown allows code to wait for a shut down signal
func ShutDown() <-chan struct{} {
	return Default().shutdown
}

// SetDefault sets the default Terminator
func SetDefault(t *Terminator) {
	defaultTerminator.Store(t)
}

// ShuttingDown returns true if shutdown is in progress
func ShuttingDown() bool {
	return Default().shuttingDown
}

// Stop signals that all goroutines in the group should safely exit
func Stop() {
	Default().shuttingDown = true
	close(Default().shutdown)
}

// Wait blocks until every goroutines in the group has called Done()
func Wait() {
	Default().wg.Wait()
}

// New returns a Terminator
func New() *Terminator {
	return &Terminator{
		shutdown: make(chan struct{}, 1),
	}
}

// Add adds delta to the count of goroutines in the group
func (t *Terminator) Add(delta int) {
	t.wg.Add(delta)
}

// Done decrements the count of goroutines in the group by one
func (t *Terminator) Done() {
	t.wg.Done()
}

// Go runs a function inside an Add(1) --- Done() sequence
func (t *Terminator) Go(f func()) {
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		f()
	}()
}

// ShutDown allows code to wait for a shut down signal
func (t *Terminator) ShutDown() <-chan struct{} {
	return t.shutdown
}

// ShuttingDown returns true if shutdown is in Default().progress
func (t *Terminator) ShuttingDown() bool {
	return t.shuttingDown
}

// Stop signals that all goroutines in the group should safely exit
func (t *Terminator) Stop() {
	t.shuttingDown = true
	close(t.shutdown)
}

// Wait blocks until every goroutines in the group has called Done()
func (t *Terminator) Wait() {
	t.wg.Wait()
}
