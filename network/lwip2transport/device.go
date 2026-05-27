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

package lwip2transport

import (
	"io"
	"sync"

	lwip "github.com/eycorsican/go-tun2socks/core"
	"golang.getoutline.org/sdk/network"
	"golang.getoutline.org/sdk/transport"
)

const packetMTU = 1500

// Compilation guard against interface implementation
var _ network.IPDevice = (*lwIPDevice)(nil)

type lwIPDevice struct {
	tcp   *tcpHandler
	udp   *udpHandler
	stack lwip.LWIPStack

	// whether the device has been closed
	done chan struct{}

	// async read call and its result
	rdBuf chan []byte
	rdN   chan int
}

// Singleton instance
var instMu sync.Mutex
var inst *lwIPDevice = nil

// ConfigureDevice configures the singleton LwIP device using the [transport.StreamDialer] to handle TCP streams and
// the [transport.PacketProxy] to handle UDP packets.
//
// LwIP device is a [network.IPDevice] that can translate IP packets to TCP/UDP traffic and vice versa. It uses the
// [lwIP library] to perform the translation.
//
// LwIP device must be a singleton object due to limitations in [lwIP library]. If you try to call ConfigureDevice more
// than once, we will Close the previous device and reconfigure it.
//
// To use a LwIP device:
//  1. Call [ConfigureDevice] with two handlers for TCP and UDP traffic.
//  2. Write IP packets to the device. The device will translate the IP packets to TCP/UDP traffic and send them to the
//     appropriate handlers.
//  3. Read IP packets from the device to get the TCP/UDP responses.
//
// A LwIP device is NOT thread-safe. However it is safe to use Write, Read/WriteTo and Close in different goroutines.
// Keep in mind that only one goroutine can call Write at a time, and only one goroutine can use either Read or
// WriteTo at a time.
//
// [lwIP library]: https://savannah.nongnu.org/projects/lwip/
func ConfigureDevice(sd transport.StreamDialer, pp network.PacketProxy) (network.IPDevice, error) {
	_ = "STUB: not implemented"
	return *new(network.IPDevice), nil
}

// Close implements [io.Closer] and [network.IPDevice]. It closes the device, rendering it unusable for I/O.
//
// Close does not close other objects that are passed to this device, such as the [transport.StreamDialer],
// [transport.PacketListener] or [io.Writer]. You are responsible for closing these objects yourself.
func (d *lwIPDevice) Close() error {
	_ = "STUB: not implemented"
	// make sure we don't close the channel twice
	return nil
}

// MTU implements [network.IPDevice]. It returns the maximum buffer size of a single IP packet that can be processed by
// this device.
func (d *lwIPDevice) MTU() int {
	_ = "STUB: not implemented"

	// forwardOutgoingIPPacket writes an IP packet response `b` to this device. The packet can be read by calling the Read
	// function, or it can be redirected to an [io.Writer] if the WriteTo function has been called. forwardOutgoingIPPacket
	// blocks until the packet is successfully consumed by a Read or WriteTo.
	//
	// forwardOutgoingIPPacket can be used as an output function for lwIP.
	//
	// forwardOutgoingIPPacket might be called by multiple goroutines (for example, when multiple UDP packets arrive at the
	// same time). We sequentialize the calls by using channels, if performance issues arise in the future, we can use
	// other more performant but more error-prone methods (e.g. the [sync] package) to resolve them.
	return 0
}

func (d *lwIPDevice) forwardOutgoingIPPacket(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Read implements [io.Reader] and [network.IPDevice]. It reads one IP packet from the TCP/UDP response, blocking until
// a packet arrives or this device is closed. If a packet is too long to fit in the supplied buffer `p`, the excess
// bytes are discarded.
//
// Read returns [io.EOF] error if this device is closed or no more data is available.
func (d *lwIPDevice) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteTo implements [io.WriterTo]. It writes all IP packets from TCP/UDP responses to `w` until all data is written
// or an error occurs. This function will not allocate any intermediate buffers.
//
// WriteTo returns the total number of bytes written and any error encountered during the write. If there are no more
// data available, WriteTo returns nil error instead of [io.EOF].
func (d *lwIPDevice) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Write implements [io.Writer] and [network.IPDevice]. It writes a single IP packet to this device. The device will
// then translate the IP packet into a TCP or UDP traffic.
//
// Write returns [network.ErrClosed] if this device is already closed.
func (d *lwIPDevice) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Workaround: lwip netstack did not use a typed error.
