// Copyright 2020 The Outline Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shadowsocks

import (
	"errors"
)

// ErrShortPacket indicates that the destination packet given to Unpack is too short.
var ErrShortPacket = errors.New("short packet")

// Assumes all ciphers have NonceSize() <= 12.
var zeroNonce [12]byte

// PackSalt encrypts a Shadowsocks-UDP packet and returns a slice containing the encrypted packet.
// dst must be big enough to hold the encrypted packet.
// If plaintext and dst overlap but are not aligned for in-place encryption, this
// function will panic.
// It uses the given [SaltGenerator] to generate the salt.
func PackSalt(dst, plaintext []byte, key *EncryptionKey, sg SaltGenerator) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pack calls PackSalt with the [RandomSaltGenerator].
func Pack(dst, plaintext []byte, key *EncryptionKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unpack decrypts a Shadowsocks-UDP packet in the format [salt][cipherText][AEAD tag] and returns a slice containing
// the decrypted payload or an error.
// If dst is present, it is used to store the plaintext, and must have enough capacity.
// If dst is nil, decryption proceeds in-place.
func Unpack(dst, pkt []byte, key *EncryptionKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
