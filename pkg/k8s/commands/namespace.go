package commands

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/khulnasoft-lab/vul-kubernetes/pkg/k8s"
	"github.com/khulnasoft-lab/vul-kubernetes/pkg/vulk8s"
	"github.com/khulnasoft-lab/vul/pkg/flag"
	"github.com/khulnasoft-lab/vul/pkg/log"
)

// namespaceRun runs scan on kubernetes cluster
func namespaceRun(ctx context.Context, opts flag.Options, cluster k8s.Cluster) error {
	if err := validateReportArguments(opts); err != nil {
		return err
	}
	var vulk vulk8s.VulK8S
	if opts.AllNamespaces {
		vulk = vulk8s.New(cluster, log.Logger).AllNamespaces()
	} else {
		vulk = vulk8s.New(cluster, log.Logger).Namespace(getNamespace(opts, cluster.GetCurrentNamespace()))
	}

	artifacts, err := vulk.ListArtifacts(ctx)
	if err != nil {
		return xerrors.Errorf("get k8s artifacts error: %w", err)
	}

	runner := newRunner(opts, cluster.GetCurrentContext())
	return runner.run(ctx, artifacts)
}

func getNamespace(opts flag.Options, currentNamespace string) string {
	if len(opts.K8sOptions.Namespace) > 0 {
		return opts.K8sOptions.Namespace
	}

	return currentNamespace
}
