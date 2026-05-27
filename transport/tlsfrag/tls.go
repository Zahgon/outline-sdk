// Copyright 2023 The Outline Authors
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

package tlsfrag

type tlsHandshakeRecordHeader []byte

// TLS record layout from [RFC 8446]:
//
//	+-------------+ 0
//	| RecordType  |
//	+-------------+ 1
//	|  Protocol   |
//	|  Version    |
//	+-------------+ 3
//	|   Record    |
//	|   Length    |
//	+-------------+ 5
//	|   Message   |
//	|    Data     |
//	|     ...     |
//	+-------------+ Message Length + 5
//
//	RecordType := invalid(0) | handshake(22) | application_data(23) | ...
//	LegacyRecordVersion := 0x0301 ("TLS 1.0") | 0x0302 ("TLS 1.1") | 0x0303 ("TLS 1.2")
//	0 < Message Length (of handshake)        ≤ 2^14
//	0 ≤ Message Length (of application_data) ≤ 2^14
//
// [RFC 8446]: https://datatracker.ietf.org/doc/html/rfc8446#section-5.1
const (
	recordHeaderLen     = 5
	maxRecordPayloadLen = 1 << 14

	recordTypeHandshake byte = 22

	versionTLS10 uint16 = 0x0301
	versionTLS11 uint16 = 0x0302
	versionTLS12 uint16 = 0x0303
	versionTLS13 uint16 = 0x0304
)

func newTLSHandshakeRecordHeader(p []byte) (tlsHandshakeRecordHeader, error) {
	_ = "STUB: not implemented"
	return *new(tlsHandshakeRecordHeader), nil
}

func (h tlsHandshakeRecordHeader) Validate() error { _ = "STUB: not implemented"; return nil }

func (h tlsHandshakeRecordHeader) PayloadLen() uint16 { _ = "STUB: not implemented"; return 0 }

func (h tlsHandshakeRecordHeader) SetPayloadLen(len uint16) error {
	_ = "STUB: not implemented"
	return nil
}
