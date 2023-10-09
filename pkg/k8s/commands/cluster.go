package commands

import (
	"context"

	"golang.org/x/exp/slices"
	"golang.org/x/xerrors"

	k8sArtifacts "github.com/khulnasoft-lab/vul-kubernetes/pkg/artifacts"
	"github.com/khulnasoft-lab/vul-kubernetes/pkg/k8s"
	"github.com/khulnasoft-lab/vul-kubernetes/pkg/vulk8s"
	"github.com/khulnasoft-lab/vul/pkg/flag"
	"github.com/khulnasoft-lab/vul/pkg/log"
	"github.com/khulnasoft-lab/vul/pkg/types"
)

// clusterRun runs scan on kubernetes cluster
func clusterRun(ctx context.Context, opts flag.Options, cluster k8s.Cluster) error {
	if err := validateReportArguments(opts); err != nil {
		return err
	}
	var artifacts []*k8sArtifacts.Artifact
	var err error
	switch opts.Format {
	case types.FormatCycloneDX:
		artifacts, err = vulk8s.New(cluster, log.Logger).ListBomInfo(ctx)
		if err != nil {
			return xerrors.Errorf("get k8s artifacts with node info error: %w", err)
		}
	case types.FormatJSON, types.FormatTable:
		if opts.Scanners.AnyEnabled(types.MisconfigScanner) && slices.Contains(opts.Components, "infra") {
			artifacts, err = vulk8s.New(cluster, log.Logger).ListArtifactAndNodeInfo(ctx, opts.NodeCollectorNamespace, opts.ExcludeNodes, opts.Tolerations...)
			if err != nil {
				return xerrors.Errorf("get k8s artifacts with node info error: %w", err)
			}
		} else {
			artifacts, err = vulk8s.New(cluster, log.Logger).ListArtifacts(ctx)
			if err != nil {
				return xerrors.Errorf("get k8s artifacts error: %w", err)
			}
		}
	default:
		return xerrors.Errorf(`unknown format %q. Use "json" or "table" or "cyclonedx"`, opts.Format)
	}

	runner := newRunner(opts, cluster.GetCurrentContext())
	return runner.run(ctx, artifacts)
}
