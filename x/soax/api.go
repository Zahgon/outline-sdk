// Copyright 2025 The Outline Authors
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

package soax

import (
	"context"
	"net/http"
)

const (
	apiHost = "api.soax.com"
)

// ConnType specifies the connection type for SOAX API calls.
type ConnType string

const (
	// ConnTypeResidential is for residential proxies, referred to as "wifi" by the API.
	ConnTypeResidential ConnType = "wifi"
	// ConnTypeMobile is for mobile proxies.
	ConnTypeMobile ConnType = "mobile"
)

// Client allows you to access the SOAX REST API.
type Client struct {
	APIKey     string
	PackageKey string
	// HTTPClient is the client to use for API calls. If nil, a default client will be used.
	HTTPClient *http.Client
	// BaseURL for testing. If empty, "https://api.soax.com" is used. This can be a plain URL, or one
	// with a path.
	BaseURL string
}

func (c *Client) httpClient() *http.Client { _ = "STUB: not implemented"; return nil }

func (c *Client) newRequest(ctx context.Context, apiPath string, queryParams map[string]string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) doAndDecode(req *http.Request, result any) error {
	_ = "STUB: not implemented"
	return nil
}

// GetResidentialISPs returns the available ISPs for the given location.
// Requires a Residential package; returns an error with a Mobile package.
// The official documentation refers to residential ISPs as "WiFi ISPs".
// API reference: https://helpcenter.soax.com/en/articles/6228391-getting-a-list-of-wifi-isps
func (c *Client) GetResidentialISPs(ctx context.Context, countryCode, regionID, cityID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetMobileISPs returns the available mobile carriers for the given location.
// Requires a Mobile package; returns an error with a Residential package.
// API reference: https://helpcenter.soax.com/en/articles/6228381-getting-a-list-of-mobile-carriers
func (c *Client) GetMobileISPs(ctx context.Context, countryCode, regionID, cityID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetRegions returns the available regions for the given country and ISP.
// connType must match the package type: [ConnTypeResidential] for Residential packages, [ConnTypeMobile] for Mobile packages.
// API reference: https://helpcenter.soax.com/en/articles/6227864-getting-a-list-of-regions
func (c *Client) GetRegions(ctx context.Context, connType ConnType, countryCode, isp string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetCities returns the available cities for the given country, ISP, and region.
// connType must match the package type: [ConnTypeResidential] for Residential packages, [ConnTypeMobile] for Mobile packages.
// API reference: https://helpcenter.soax.com/en/articles/6228092-getting-a-list-of-cities
func (c *Client) GetCities(ctx context.Context, connType ConnType, countryCode, isp, regionID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
