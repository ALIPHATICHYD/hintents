// Copyright 2025 Erst Users
// SPDX-License-Identifier: Apache-2.0

package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentinelErrors(t *testing.T) {
	assert.NotNil(t, ErrTransactionNotFound)
	assert.NotNil(t, ErrRPCConnectionFailed)
	assert.NotNil(t, ErrSimulatorNotFound)
	assert.NotNil(t, ErrSimulationFailed)
	assert.NotNil(t, ErrInvalidNetwork)
	assert.NotNil(t, ErrMarshalFailed)
	assert.NotNil(t, ErrUnmarshalFailed)
	assert.NotNil(t, ErrSimulationLogicError)
	assert.NotNil(t, ErrGasModelInvalid)
	assert.NotNil(t, ErrAuthorizationFailed)
	assert.NotNil(t, ErrLedgerEntryNotFound)
}

func TestWrap(t *testing.T) {
	baseErr := fmt.Errorf("base error")

	tests := []struct {
		name     string
		sentinel error
		format   string
		args     []interface{}
		checkFn  func(error) bool
		checkMsg string
	}{
		{
			name:     "wrap with error",
			sentinel: ErrTransactionNotFound,
			format:   "%w",
			args:     []interface{}{baseErr},
			checkFn:  func(e error) bool { return errors.Is(e, ErrTransactionNotFound) && errors.Is(e, baseErr) },
			checkMsg: "should contain both errors",
		},
		{
			name:     "wrap with string",
			sentinel: ErrSimulatorNotFound,
			format:   "test message",
			args:     []interface{}{},
			checkFn:  func(e error) bool { return errors.Is(e, ErrSimulatorNotFound) },
			checkMsg: "should contain message",
		},
		{
			name:     "wrap with formatted args",
			sentinel: ErrMarshalFailed,
			format:   "failed: %w, detail: %s",
			args:     []interface{}{baseErr, "extra info"},
			checkFn:  func(e error) bool { return errors.Is(e, ErrMarshalFailed) },
			checkMsg: "should preserve format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.sentinel, tt.format, tt.args...)
			assert.True(t, tt.checkFn(err), tt.checkMsg)
		})
	}
}

func TestIsFunction(t *testing.T) {
	baseErr := fmt.Errorf("base error")
	wrappedErr := Wrap(ErrTransactionNotFound, "%w", baseErr)

	assert.True(t, Is(wrappedErr, ErrTransactionNotFound))
	assert.False(t, Is(wrappedErr, ErrRPCConnectionFailed))
	assert.True(t, Is(wrappedErr, baseErr))
}

func TestErrorComparison(t *testing.T) {
	baseErr := fmt.Errorf("test")
	err1 := Wrap(ErrTransactionNotFound, "%w", baseErr)
	err2 := Wrap(ErrRPCConnectionFailed, "%w", baseErr)

	assert.True(t, errors.Is(err1, ErrTransactionNotFound))
	assert.False(t, errors.Is(err1, ErrRPCConnectionFailed))
	assert.True(t, errors.Is(err2, ErrRPCConnectionFailed))
	assert.False(t, errors.Is(err2, ErrTransactionNotFound))

	assert.True(t, errors.Is(err2, ErrRPCConnectionFailed))
	assert.False(t, errors.Is(err2, ErrTransactionNotFound))
}
