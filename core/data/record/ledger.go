// Copyright (C) 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package record

import (
	"github.com/google/gapid/core/event"
	"github.com/google/gapid/core/log"
)

// Ledger is the interface to a sequential immutable record store.
// New records can only be added to a ledger, and old records cannot be modified.
type Ledger interface {
	// Read will read all records already in the ledger and feed them to the supplied writer.
	Read(ctx log.Context, handler event.Handler) error
	// Watch will feed new entries as they arrive into the supplied handler.
	Watch(ctx log.Context, handler event.Handler)
	// Add a new entry to the ledger.
	// The record must be of the same type that the ledger was opened with.
	Add(ctx log.Context, record interface{}) error
	// Close is called to close the register, and notify all watchers.
	Close(ctx log.Context)
	// New returns a new record of the type the ledger stores.
	New(ctx log.Context) interface{}
}

type ledger struct {
	instance LedgerInstance
	onAdd    event.Broadcast
}

// LedgerInstance is the interface to an implementation support class for a ledger.
// It is responsible for reading and writing the backing store.
type LedgerInstance interface {
	Write(ctx log.Context, event interface{}) error
	Reader(ctx log.Context) event.Source
	New(ctx log.Context) interface{}
	Close(ctx log.Context)
}

// NewLedger returns a ledger from a backing store.
func NewLedger(ctx log.Context, instance LedgerInstance) Ledger {
	return &ledger{instance: instance}
}

func (l *ledger) Read(ctx log.Context, handler event.Handler) error {
	r := l.instance.Reader(ctx)
	defer r.Close(ctx)
	return event.Feed(ctx, handler, r.Next)
}

func (l *ledger) Watch(ctx log.Context, handler event.Handler) {
	l.onAdd.Listen(ctx, handler)
}

func (l *ledger) Add(ctx log.Context, record interface{}) error {
	if err := l.instance.Write(ctx, record); err != nil {
		return err
	}
	if err := l.onAdd.Send(ctx, record); err != nil {
		return err
	}
	return nil
}

func (l *ledger) Close(ctx log.Context) {
	l.instance.Close(ctx)
	l.onAdd.Send(ctx, nil)
}

func (l *ledger) New(ctx log.Context) interface{} {
	return l.instance.New(ctx)
}
