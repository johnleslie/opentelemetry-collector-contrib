// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nfsscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/nfsscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/nfsscraper/internal/metadata"
)

// nfsScraper for NFS Metrics
type nfsScraper struct {
	settings scraper.Settings
	config   *Config
	mb       *metadata.MetricsBuilder

	// for mocking
	getNfsNetStats		func() (*NfsNetStats, error)
	getNfsRpcStats		func() (*NfsRpcStats, error)
	getNfsRPCProcStats	func() ([]*RPCProcedureStats, error)
	getNfsdRPCProcStats	func() ([]*RPCProcedureStats, error)
	getNfsdRepcacheStats	func() (*NfsdRepcacheStats, error)
	getNfsdFhStats		func() (*NfsdFhStats, error)
	getNfsdIoStats		func() (*NfsdIoStats, error)
	getNfsdThreadStats	func() (*NfsdThreadStats, error)
	getNfsdNetStats		func() (*NfsdNetStats, error)
	getNfsdRpcStats		func() (*NfsdRpcStats, error)

}

// newNfsScraper creates an Uptime related metric
func newNfsScraper(_ context.Context, settings scraper.Settings, cfg *Config) *nfsScraper {
	return &nfsScraper{settings: settings, config: cfg}
}

func (s *nfsScraper) start(_ context.Context, _ component.Host) error {
	s.mb = metadata.NewMetricsBuilder(s.config.MetricsBuilderConfig, s.settings)
	return nil
}

func (s *nfsScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	return pmetric.NewMetrics(), nil
}
