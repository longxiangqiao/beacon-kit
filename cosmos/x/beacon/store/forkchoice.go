// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package store

import (
	"errors"

	"cosmossdk.io/store"
)

// Forkchoice represents the fork choice rule in the blockchain.
type Forkchoice struct {
	store store.KVStore

	lastValidHead [32]byte
}

// NewForkchoice creates a new instance of Forkchoice.
func NewForkchoice(store store.KVStore) *Forkchoice {
	return &Forkchoice{
		store: store,
	}
}

// SetSafeBlockHash sets the safe block hash in the store.
func (f *Forkchoice) SetSafeBlockHash(safeBlockHash [32]byte) error {
	f.store.Set([]byte("forkchoice_safe"), safeBlockHash[:])
	return nil
}

// GetSafeBlockHash retrieves the safe block hash from the store.
func (f *Forkchoice) GetSafeBlockHash() ([32]byte, error) {
	bz := f.store.Get([]byte("forkchoice_safe"))
	if bz == nil {
		return [32]byte{}, errors.New("safe block hash not found")
	}
	var safeBlockHash [32]byte
	copy(safeBlockHash[:], bz)
	return safeBlockHash, nil
}

// SetFinalizedBlockHash sets the finalized block hash in the store.
func (f *Forkchoice) SetFinalizedBlockHash(finalizedBlockHash [32]byte) error {
	f.store.Set([]byte("forkchoice_finalized"), finalizedBlockHash[:])
	return nil
}

// GetFinalizedBlockHash retrieves the finalized block hash from the store.
func (f *Forkchoice) GetFinalizedBlockHash() ([32]byte, error) {
	bz := f.store.Get([]byte("forkchoice_finalized"))
	if bz == nil {
		return [32]byte{}, errors.New("finalized block hash not found")
	}
	var finalizedBlockHash [32]byte
	copy(finalizedBlockHash[:], bz)
	return finalizedBlockHash, nil
}

// SetLastValidHead sets the last valid head in the store.
func (f *Forkchoice) SetLastValidHead(lastValidHead [32]byte) error {
	f.lastValidHead = lastValidHead
	return nil
}

// GetLastValidHead retrieves the last valid head from the store.
func (f *Forkchoice) GetLastValidHead() [32]byte {
	return f.lastValidHead
}
