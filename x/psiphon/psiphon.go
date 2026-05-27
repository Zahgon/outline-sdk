// Copyright 2024 The Outline Authors
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

package psiphon

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"sync"

	"golang.getoutline.org/sdk/transport"
)

// The single [Dialer] we can have.
var singletonDialer Dialer

var (
	errNotStartedDial = errors.New("dialer has not been started yet")
	errNotStartedStop = errors.New("tried to stop dialer that is not running")
)

// DialerConfig specifies the parameters for [Dialer].
type DialerConfig struct {
	// Used as the directory for the datastore, remote server list, and obfuscasted
	// server list.
	// Empty string means the default will be used (current working directory).
	// Strongly recommended.
	DataRootDirectory string

	// Raw JSON config provided by Psiphon.
	ProviderConfig json.RawMessage
}

// Dialer is a [transport.StreamDialer] that uses Psiphon to connect to a destination.
// There's only one possible Psiphon Dialer available at any time, which is accessible via [GetSingletonDialer].
// The zero value of this type is invalid.
//
// The Dialer must be configured first with [Dialer.Start] before it can be used, and [Dialer.Stop] must be
// called before you can start it again with a new configuration. Dialer.Stop should be called
// when you no longer need the Dialer in order to release resources.
type Dialer struct {
	// Controls the Dialer state and Psiphon's global state.
	mu sync.Mutex
	// Used by DialStream.
	tunnel psiphonTunnel
	// Used by Stop.
	stop func()
}

type psiphonTunnel interface {
	Dial(remoteAddr string) (net.Conn, error)
	Stop()
}

var _ transport.StreamDialer = (*Dialer)(nil)

// DialStream implements [transport.StreamDialer].
// The context is not used because Psiphon's implementation doesn't support it. If you need cancellation,
// you will need to add it independently.
func (d *Dialer) DialStream(unusedContext context.Context, addr string) (transport.StreamConn, error) {
	_ = "STUB: not implemented"
	return *new(transport.StreamConn), nil
}

func getClientPlatform() string { _ = "STUB: not implemented"; return "" }

// Allows for overriding in tests.
var startTunnel func(ctx context.Context, config *DialerConfig) (psiphonTunnel, error) = psiphonStartTunnel

func psiphonStartTunnel(tunnelCtx context.Context, config *DialerConfig) (psiphonTunnel, error) {
	_ = "STUB: not implemented"
	return *new(psiphonTunnel), nil
}

// Note that these parameters override anything in the provider config.

// Disable Psiphon's local proxy servers, which we don't use.

// Start configures and runs the Dialer. It must be called before you can use the Dialer. It returns when the tunnel is ready.
func (d *Dialer) Start(startCtx context.Context, config *DialerConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// If we are already started stop first.

// Make sure we unlock the mutex so that the previous Start can complete.

// startCtx is intended for the lifetime of the startup.
// dialerCtx is intended for the lifetime of the tunnel.

// This ties startCtx and dialerCtx together
// so dialerCtx will be cancelled if startCtx is cancelled.
// We run detatchContexts after startTunnel to disconnect them.

// Tell start to stop.

// Wait for tunnel to be done.

// Cleanup.

// wait for Stop

// Stop stops the Dialer background processes, releasing resources and allowing it to be reconfigured.
// It returns when the Dialer is completely stopped.
func (d *Dialer) Stop() error { _ = "STUB: not implemented"; return nil }

// GetSingletonDialer returns the single Psiphon dialer instance.
func GetSingletonDialer() *Dialer { _ = "STUB: not implemented"; return nil }

// streamConn wraps a [net.Conn] to provide a [transport.StreamConn] interface.
type streamConn struct {
	net.Conn
}

var _ transport.StreamConn = (*streamConn)(nil)

func (c streamConn) CloseWrite() error { _ = "STUB: not implemented"; return nil }

func (c streamConn) CloseRead() error { _ = "STUB: not implemented"; return nil }
