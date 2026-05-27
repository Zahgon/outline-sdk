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

package main

import (
	"image/color"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.getoutline.org/sdk/transport"
)

type runningProxy struct {
	server  *http.Server
	Address string
}

func (p *runningProxy) Close() {
	_ = "STUB: not implemented"

	// newFilteredStreamDialer creates a direct [transport.StreamDialer] that blocks
	// non public IPs to prevent access to localhost or the local network.
	return
}

func newFilteredStreamDialer() transport.StreamDialer {
	_ = "STUB: not implemented"
	return *new(transport.StreamDialer)
}

func runServer(address, transport string) (*runningProxy, error) {
	_ = "STUB: not implemented"
	// TODO: block localhost, maybe local net.
	return nil, nil
}

type appTheme struct {
	fyne.Theme
}

const ColorNameOnPrimary = "OnPrimary"

func (t *appTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	_ = "STUB: not implemented"
	return *new(color.Color)
}

func makeAppHeader(title string) *fyne.Container { _ = "STUB: not implemented"; return nil }

func main() {
	fyneApp := app.New()
	if meta := fyneApp.Metadata(); meta.Name == "" {
		// App not packaged, probably from `go run`.
		meta.Name = "Local Proxy"
		app.SetMetadata(meta)
	}
	fyneApp.Settings().SetTheme(&appTheme{theme.DefaultTheme()})

	mainWin := fyneApp.NewWindow(fyneApp.Metadata().Name)
	mainWin.Resize(fyne.Size{Width: 350})

	addressEntry := widget.NewEntry()
	addressEntry.SetPlaceHolder("Enter proxy local address")
	addressEntry.Text = "localhost:8080"

	configEntry := widget.NewMultiLineEntry()
	configEntry.Wrapping = fyne.TextWrapBreak
	configEntry.SetPlaceHolder("Enter transport config")

	statusBox := widget.NewLabel("")
	statusBox.Wrapping = fyne.TextWrapWord

	startStopButton := widget.NewButton("", func() {})
	startStopButton.Importance = widget.HighImportance
	setProxyUI := func(proxy *runningProxy, err error) {
		if proxy != nil {
			statusBox.SetText("Proxy listening on " + proxy.Address)
			addressEntry.Disable()
			configEntry.Disable()
			startStopButton.SetText("Stop")
			startStopButton.SetIcon(theme.MediaStopIcon())
			return
		}
		if err != nil {
			statusBox.SetText("❌ ERROR: " + err.Error())
		} else {
			statusBox.SetText("Proxy not running")
		}
		addressEntry.Enable()
		configEntry.Enable()
		startStopButton.SetText("Start")
		startStopButton.SetIcon(theme.MediaPlayIcon())
	}
	var proxy *runningProxy
	startStopButton.OnTapped = func() {
		log.Println(startStopButton.Text)
		var err error
		if proxy == nil {
			// Start proxy.
			proxy, err = runServer(addressEntry.Text, configEntry.Text)
		} else {
			// Stop proxy
			proxy.Close()
			proxy = nil
		}
		setProxyUI(proxy, err)
	}
	setProxyUI(proxy, nil)

	content := container.NewVBox(
		makeAppHeader(fyneApp.Metadata().Name),
		container.NewPadded(
			container.NewVBox(
				widget.NewLabelWithStyle("Local address", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				addressEntry,
				widget.NewRichTextFromMarkdown("**Transport config** ([format](https://pkg.go.dev/golang.getoutline.org/sdk/x/configurl#hdr-Config_Format))"),
				configEntry,
				container.NewHBox(layout.NewSpacer(), startStopButton),
				widget.NewLabelWithStyle("Status", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				statusBox,
			),
		),
	)
	mainWin.SetContent(content)
	mainWin.Show()
	fyneApp.Run()
}
