// Copyright 2025 Erst Users
// SPDX-License-Identifier: Apache-2.0

package errors

import (
	"errors"
	"fmt"
)

var (
	ErrTransactionNotFound  = errors.New("transaction not found")
	ErrRPCConnectionFailed  = errors.New("RPC connection failed")
	ErrSimulatorNotFound    = errors.New("simulator binary not found")
	ErrSimulationFailed     = errors.New("simulation execution failed")
	ErrInvalidNetwork       = errors.New("invalid network")
	ErrMarshalFailed        = errors.New("failed to marshal request")
	ErrUnmarshalFailed      = errors.New("failed to unmarshal response")
	ErrSimulationLogicError = errors.New("simulation logic error")
	ErrGasModelInvalid      = errors.New("invalid gas model")
	ErrAuthorizationFailed  = errors.New("authorization failed")
	ErrLedgerEntryNotFound  = errors.New("ledger entry not found")
)

func Wrap(sentinel error, format string, args ...interface{}) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: %s", sentinel, format)
	}
	return fmt.Errorf("%w: "+format, append([]interface{}{sentinel}, args...)...)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
