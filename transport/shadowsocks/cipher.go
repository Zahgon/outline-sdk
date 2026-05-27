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
	"crypto/cipher"

	"golang.org/x/crypto/chacha20poly1305"
)

type cipherSpec struct {
	newInstance func(key []byte) (cipher.AEAD, error)
	keySize     int
	saltSize    int
	tagSize     int
}

// List of supported AEAD ciphers, as specified at https://shadowsocks.org/guide/aead.html
var (
	CHACHA20IETFPOLY1305 = "AEAD_CHACHA20_POLY1305"
	AES256GCM            = "AEAD_AES_256_GCM"
	AES192GCM            = "AEAD_AES_192_GCM"
	AES128GCM            = "AEAD_AES_128_GCM"
)

var (
	chacha20IETFPOLY1305Cipher = &cipherSpec{chacha20poly1305.New, chacha20poly1305.KeySize, 32, 16}
	aes256GCMCipher            = &cipherSpec{newAesGCM, 32, 32, 16}
	aes192GCMCipher            = &cipherSpec{newAesGCM, 24, 24, 16}
	aes128GCMCipher            = &cipherSpec{newAesGCM, 16, 16, 16}
)

var supportedCiphers = [](string){CHACHA20IETFPOLY1305, AES256GCM, AES192GCM, AES128GCM}

// ErrUnsupportedCipher is returned by [CypherByName] when the named cipher is not supported.
type ErrUnsupportedCipher struct {
	// The name of the requested [Cipher]
	Name string
}

func (err ErrUnsupportedCipher) Error() string { _ = "STUB: not implemented"; return "" }

// Largest tag size among the supported ciphers. Used by the TCP buffer pool
const maxTagSize = 16

// CipherByName returns a [*Cipher] with the given name, or an error if the cipher is not supported.
// The name must be the IETF name (as per https://www.iana.org/assignments/aead-parameters/aead-parameters.xhtml) or the
// Shadowsocks alias from https://shadowsocks.org/guide/aead.html.
func cipherByName(name string) (*cipherSpec, error) { _ = "STUB: not implemented"; return nil, nil }

func newAesGCM(key []byte) (cipher.AEAD, error) {
	_ = "STUB: not implemented"
	return *new(cipher.AEAD), nil
}

// EncryptionKey encapsulates a Shadowsocks AEAD spec and a secret
type EncryptionKey struct {
	cipher *cipherSpec
	secret []byte
}

// SaltSize is the size of the salt for this Cipher
func (c *EncryptionKey) SaltSize() int { _ = "STUB: not implemented"; return 0 }

// TagSize is the size of the AEAD tag for this Cipher
func (c *EncryptionKey) TagSize() int { _ = "STUB: not implemented"; return 0 }

var subkeyInfo = []byte("ss-subkey")

// NewAEAD creates the AEAD for this cipher
func (c *EncryptionKey) NewAEAD(salt []byte) (cipher.AEAD, error) {
	_ = "STUB: not implemented"
	return *new(cipher.AEAD), nil
}

// Function definition at https://www.openssl.org/docs/manmaster/man3/EVP_BytesToKey.html
func simpleEVPBytesToKey(data []byte, keyLen int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewEncryptionKey creates a Cipher with a cipher name and a secret.
// The cipher name must be the IETF name (as per https://www.iana.org/assignments/aead-parameters/aead-parameters.xhtml)
// or the Shadowsocks alias from https://shadowsocks.org/guide/aead.html.
func NewEncryptionKey(cipherName string, secretText string) (*EncryptionKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Key derivation as per https://shadowsocks.org/en/spec/AEAD-Ciphers.html
