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

package main

import (
	"context"

	"golang.getoutline.org/sdk/x/examples/outline-connectivity-app/shared_backend"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	_ = "STUB: not implemented"

	// startup is called when the app starts. The context is saved
	// so we can call the runtime methods
	return nil
}

func (a *App) startup(ctx context.Context) { _ = "STUB: not implemented"; return }

func (a *App) Request(resourceName string, parameters string) (shared_backend.Response, error) {
	_ = "STUB: not implemented"
	return *new(shared_backend.Response), nil
}

// TODO: make this non-blocking with goroutines/channels
