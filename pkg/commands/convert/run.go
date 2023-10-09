package convert

import (
	"context"
	"encoding/json"
	"os"

	"golang.org/x/xerrors"

	"github.com/khulnasoft-lab/vul/pkg/commands/operation"
	"github.com/khulnasoft-lab/vul/pkg/flag"
	"github.com/khulnasoft-lab/vul/pkg/log"
	"github.com/khulnasoft-lab/vul/pkg/report"
	"github.com/khulnasoft-lab/vul/pkg/result"
	"github.com/khulnasoft-lab/vul/pkg/types"
)

func Run(ctx context.Context, opts flag.Options) (err error) {
	f, err := os.Open(opts.Target)
	if err != nil {
		return xerrors.Errorf("file open error: %w", err)
	}
	defer f.Close()

	var r types.Report
	if err = json.NewDecoder(f).Decode(&r); err != nil {
		return xerrors.Errorf("json decode error: %w", err)
	}

	// "convert" supports JSON results produced by Vul scanning other than AWS and Kubernetes
	if r.ArtifactName == "" && r.ArtifactType == "" {
		return xerrors.New("AWS and Kubernetes scanning reports are not yet supported")
	}

	if err = result.Filter(ctx, r, opts.FilterOpts()); err != nil {
		return xerrors.Errorf("unable to filter results: %w", err)
	}

	log.Logger.Debug("Writing report to output...")
	if err = report.Write(r, opts); err != nil {
		return xerrors.Errorf("unable to write results: %w", err)
	}

	operation.ExitOnEOL(opts, r.Metadata)
	operation.Exit(opts, r.Results.Failed())

	return nil
}