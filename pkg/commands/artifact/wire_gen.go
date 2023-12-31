// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package artifact

import (
	"context"
	"github.com/aquasecurity/fanal/analyzer"
	"github.com/aquasecurity/fanal/analyzer/config"
	"github.com/aquasecurity/fanal/applier"
	image2 "github.com/aquasecurity/fanal/artifact/image"
	local2 "github.com/aquasecurity/fanal/artifact/local"
	"github.com/aquasecurity/fanal/artifact/remote"
	"github.com/aquasecurity/fanal/cache"
	"github.com/aquasecurity/fanal/image"
	"github.com/khulnasoft-lab/vul-db/pkg/db"
	"github.com/khulnasoft-lab/vul/pkg/detector/ospkg"
	"github.com/khulnasoft-lab/vul/pkg/scanner"
	"github.com/khulnasoft-lab/vul/pkg/scanner/local"
	"github.com/khulnasoft-lab/vul/pkg/types"
	"github.com/khulnasoft-lab/vul/pkg/vulnerability"
	"time"
)

// Injectors from inject.go:

func initializeDockerScanner(ctx context.Context, imageName string, artifactCache cache.ArtifactCache, localArtifactCache cache.LocalArtifactCache, timeout time.Duration, disableAnalyzers []analyzer.Type, configScannerOption config.ScannerOption) (scanner.Scanner, func(), error) {
	applierApplier := applier.NewApplier(localArtifactCache)
	detector := ospkg.Detector{}
	localScanner := local.NewScanner(applierApplier, detector)
	dockerOption, err := types.GetDockerOption(timeout)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	imageImage, cleanup, err := image.NewDockerImage(ctx, imageName, dockerOption)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	artifact, err := image2.NewArtifact(imageImage, artifactCache, disableAnalyzers, configScannerOption)
	if err != nil {
		cleanup()
		return scanner.Scanner{}, nil, err
	}
	scannerScanner := scanner.NewScanner(localScanner, artifact)
	return scannerScanner, func() {
		cleanup()
	}, nil
}

func initializeArchiveScanner(ctx context.Context, filePath string, artifactCache cache.ArtifactCache, localArtifactCache cache.LocalArtifactCache, timeout time.Duration, disableAnalyzers []analyzer.Type, configScannerOption config.ScannerOption) (scanner.Scanner, error) {
	applierApplier := applier.NewApplier(localArtifactCache)
	detector := ospkg.Detector{}
	localScanner := local.NewScanner(applierApplier, detector)
	imageImage, err := image.NewArchiveImage(filePath)
	if err != nil {
		return scanner.Scanner{}, err
	}
	artifact, err := image2.NewArtifact(imageImage, artifactCache, disableAnalyzers, configScannerOption)
	if err != nil {
		return scanner.Scanner{}, err
	}
	scannerScanner := scanner.NewScanner(localScanner, artifact)
	return scannerScanner, nil
}

func initializeFilesystemScanner(ctx context.Context, dir string, artifactCache cache.ArtifactCache, localArtifactCache cache.LocalArtifactCache, disableAnalyzers []analyzer.Type, configScannerOption config.ScannerOption) (scanner.Scanner, func(), error) {
	applierApplier := applier.NewApplier(localArtifactCache)
	detector := ospkg.Detector{}
	localScanner := local.NewScanner(applierApplier, detector)
	artifact, err := local2.NewArtifact(dir, artifactCache, disableAnalyzers, configScannerOption)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	scannerScanner := scanner.NewScanner(localScanner, artifact)
	return scannerScanner, func() {
	}, nil
}

func initializeRepositoryScanner(ctx context.Context, url string, artifactCache cache.ArtifactCache, localArtifactCache cache.LocalArtifactCache, disableAnalyzers []analyzer.Type, configScannerOption config.ScannerOption) (scanner.Scanner, func(), error) {
	applierApplier := applier.NewApplier(localArtifactCache)
	detector := ospkg.Detector{}
	localScanner := local.NewScanner(applierApplier, detector)
	artifact, cleanup, err := remote.NewArtifact(url, artifactCache, disableAnalyzers, configScannerOption)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	scannerScanner := scanner.NewScanner(localScanner, artifact)
	return scannerScanner, func() {
		cleanup()
	}, nil
}

func initializeResultClient() vulnerability.Client {
	dbConfig := db.Config{}
	client := vulnerability.NewClient(dbConfig)
	return client
}
