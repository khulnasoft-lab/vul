// +build wireinject

package client

import (
	"context"
	"time"

	"github.com/google/wire"

	"github.com/aquasecurity/fanal/analyzer"
	"github.com/aquasecurity/fanal/analyzer/config"
	"github.com/aquasecurity/fanal/cache"
	"github.com/khulnasoft-lab/vul/pkg/rpc/client"
	"github.com/khulnasoft-lab/vul/pkg/scanner"
	"github.com/khulnasoft-lab/vul/pkg/vulnerability"
)

func initializeDockerScanner(ctx context.Context, imageName string, artifactCache cache.ArtifactCache, customHeaders client.CustomHeaders,
	url client.RemoteURL, timeout time.Duration, disabled []analyzer.Type, configScannerOption config.ScannerOption) (
	scanner.Scanner, func(), error) {
	wire.Build(scanner.RemoteDockerSet)
	return scanner.Scanner{}, nil, nil
}

func initializeArchiveScanner(ctx context.Context, filePath string, artifactCache cache.ArtifactCache,
	customHeaders client.CustomHeaders, url client.RemoteURL, timeout time.Duration, disabled []analyzer.Type,
	configScannerOption config.ScannerOption) (scanner.Scanner, error) {
	wire.Build(scanner.RemoteArchiveSet)
	return scanner.Scanner{}, nil
}

func initializeResultClient() vulnerability.Client {
	wire.Build(vulnerability.SuperSet)
	return vulnerability.Client{}
}
