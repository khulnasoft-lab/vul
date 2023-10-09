package rpc

import (
	"strings"
	"time"

	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	dbTypes "github.com/khulnasoft-lab/vul-db/pkg/types"
	"github.com/khulnasoft-lab/vul/pkg/digest"
	ftypes "github.com/khulnasoft-lab/vul/pkg/fanal/types"
	"github.com/khulnasoft-lab/vul/pkg/log"
	"github.com/khulnasoft-lab/vul/pkg/types"
	"github.com/khulnasoft-lab/vul/rpc/cache"
	"github.com/khulnasoft-lab/vul/rpc/common"
	"github.com/khulnasoft-lab/vul/rpc/scanner"
)

// ConvertToRPCPkgs returns the list of RPC package objects
func ConvertToRPCPkgs(pkgs []ftypes.Package) []*common.Package {
	var rpcPkgs []*common.Package
	for _, pkg := range pkgs {
		rpcPkgs = append(rpcPkgs, &common.Package{
			Id:         pkg.ID,
			Name:       pkg.Name,
			Version:    pkg.Version,
			Release:    pkg.Release,
			Epoch:      int32(pkg.Epoch),
			Arch:       pkg.Arch,
			Dev:        pkg.Dev,
			SrcName:    pkg.SrcName,
			SrcVersion: pkg.SrcVersion,
			SrcRelease: pkg.SrcRelease,
			SrcEpoch:   int32(pkg.SrcEpoch),
			Licenses:   pkg.Licenses,
			Layer:      ConvertToRPCLayer(pkg.Layer),
			FilePath:   pkg.FilePath,
			DependsOn:  pkg.DependsOn,
			Digest:     pkg.Digest.String(),
			Indirect:   pkg.Indirect,
		})
	}
	return rpcPkgs
}

func ConvertToRPCCustomResources(resources []ftypes.CustomResource) []*common.CustomResource {
	var rpcResources []*common.CustomResource
	for _, r := range resources {
		data, err := structpb.NewValue(r.Data)
		if err != nil {
			log.Logger.Warn(err)
		}
		rpcResources = append(rpcResources, &common.CustomResource{
			Type:     r.Type,
			FilePath: r.FilePath,
			Layer: &common.Layer{
				Digest: r.Layer.Digest,
				DiffId: r.Layer.DiffID,
			},
			Data: data,
		})
	}
	return rpcResources
}

func ConvertToRPCCode(code ftypes.Code) *common.Code {
	var rpcLines []*common.Line
	for _, line := range code.Lines {
		rpcLines = append(rpcLines, &common.Line{
			Number:      int32(line.Number),
			Content:     line.Content,
			IsCause:     line.IsCause,
			Annotation:  line.Annotation,
			Truncated:   line.Truncated,
			Highlighted: line.Highlighted,
			FirstCause:  line.FirstCause,
			LastCause:   line.LastCause,
		})
	}
	return &common.Code{
		Lines: rpcLines,
	}
}

func ConvertToRPCSecrets(secrets []ftypes.Secret) []*common.Secret {
	var rpcSecrets []*common.Secret
	for _, s := range secrets {
		rpcSecrets = append(rpcSecrets, &common.Secret{
			Filepath: s.FilePath,
			Findings: ConvertToRPCSecretFindings(s.Findings),
		})
	}
	return rpcSecrets
}

func ConvertToRPCSecretFindings(findings []ftypes.SecretFinding) []*common.SecretFinding {
	var rpcFindings []*common.SecretFinding
	for _, f := range findings {
		rpcFindings = append(rpcFindings, &common.SecretFinding{
			RuleId:    f.RuleID,
			Category:  string(f.Category),
			Severity:  f.Severity,
			Title:     f.Title,
			EndLine:   int32(f.EndLine),
			StartLine: int32(f.StartLine),
			Code:      ConvertToRPCCode(f.Code),
			Match:     f.Match,
			Layer:     ConvertToRPCLayer(f.Layer),
		})
	}
	return rpcFindings
}

// ConvertFromRPCPkgs returns list of Fanal package objects
func ConvertFromRPCPkgs(rpcPkgs []*common.Package) []ftypes.Package {
	var pkgs []ftypes.Package
	for _, pkg := range rpcPkgs {
		pkgs = append(pkgs, ftypes.Package{
			ID:         pkg.Id,
			Name:       pkg.Name,
			Version:    pkg.Version,
			Release:    pkg.Release,
			Epoch:      int(pkg.Epoch),
			Arch:       pkg.Arch,
			Dev:        pkg.Dev,
			SrcName:    pkg.SrcName,
			SrcVersion: pkg.SrcVersion,
			SrcRelease: pkg.SrcRelease,
			SrcEpoch:   int(pkg.SrcEpoch),
			Licenses:   pkg.Licenses,
			Layer:      ConvertFromRPCLayer(pkg.Layer),
			FilePath:   pkg.FilePath,
			DependsOn:  pkg.DependsOn,
			Digest:     digest.Digest(pkg.Digest),
			Indirect:   pkg.Indirect,
		})
	}
	return pkgs
}

// ConvertToRPCVulns returns common.Vulnerability
func ConvertToRPCVulns(vulns []types.DetectedVulnerability) []*common.Vulnerability {
	var rpcVulns []*common.Vulnerability
	for _, vuln := range vulns {
		severity, err := dbTypes.NewSeverity(vuln.Severity)
		if err != nil {
			log.Logger.Warn(err)
		}
		cvssMap := make(map[string]*common.CVSS) // This is needed because protobuf generates a map[string]*CVSS type
		for vendor, vendorSeverity := range vuln.CVSS {
			cvssMap[string(vendor)] = &common.CVSS{
				V2Vector: vendorSeverity.V2Vector,
				V3Vector: vendorSeverity.V3Vector,
				V2Score:  vendorSeverity.V2Score,
				V3Score:  vendorSeverity.V3Score,
			}
		}
		vensorSeverityMap := make(map[string]common.Severity)
		for vendor, vendorSeverity := range vuln.VendorSeverity {
			vensorSeverityMap[string(vendor)] = common.Severity(vendorSeverity)
		}

		var lastModifiedDate, publishedDate *timestamppb.Timestamp
		if vuln.LastModifiedDate != nil {
			lastModifiedDate = timestamppb.New(*vuln.LastModifiedDate) // nolint: errcheck
		}

		if vuln.PublishedDate != nil {
			publishedDate = timestamppb.New(*vuln.PublishedDate) // nolint: errcheck
		}

		var customAdvisoryData, customVulnData *structpb.Value
		if vuln.Custom != nil {
			customAdvisoryData, _ = structpb.NewValue(vuln.Custom) // nolint: errcheck
		}
		if vuln.Vulnerability.Custom != nil {
			customVulnData, _ = structpb.NewValue(vuln.Vulnerability.Custom) // nolint: errcheck
		}

		rpcVulns = append(rpcVulns, &common.Vulnerability{
			VulnerabilityId:    vuln.VulnerabilityID,
			VendorIds:          vuln.VendorIDs,
			PkgId:              vuln.PkgID,
			PkgName:            vuln.PkgName,
			PkgPath:            vuln.PkgPath,
			InstalledVersion:   vuln.InstalledVersion,
			FixedVersion:       vuln.FixedVersion,
			Status:             int32(vuln.Status),
			Title:              vuln.Title,
			Description:        vuln.Description,
			Severity:           common.Severity(severity),
			VendorSeverity:     vensorSeverityMap,
			References:         vuln.References,
			Layer:              ConvertToRPCLayer(vuln.Layer),
			Cvss:               cvssMap,
			SeveritySource:     string(vuln.SeveritySource),
			CweIds:             vuln.CweIDs,
			PrimaryUrl:         vuln.PrimaryURL,
			LastModifiedDate:   lastModifiedDate,
			PublishedDate:      publishedDate,
			CustomAdvisoryData: customAdvisoryData,
			CustomVulnData:     customVulnData,
			DataSource:         ConvertToRPCDataSource(vuln.DataSource),
		})
	}
	return rpcVulns
}

// ConvertToRPCMisconfs returns common.DetectedMisconfigurations
func ConvertToRPCMisconfs(misconfs []types.DetectedMisconfiguration) []*common.DetectedMisconfiguration {
	var rpcMisconfs []*common.DetectedMisconfiguration
	for _, m := range misconfs {
		severity, err := dbTypes.NewSeverity(m.Severity)
		if err != nil {
			log.Logger.Warn(err)
		}

		rpcMisconfs = append(rpcMisconfs, &common.DetectedMisconfiguration{
			Type:          m.Type,
			Id:            m.ID,
			AvdId:         m.AVDID,
			Title:         m.Title,
			Description:   m.Description,
			Message:       m.Message,
			Namespace:     m.Namespace,
			Query:         m.Query,
			Resolution:    m.Resolution,
			Severity:      common.Severity(severity),
			PrimaryUrl:    m.PrimaryURL,
			References:    m.References,
			Status:        string(m.Status),
			Layer:         ConvertToRPCLayer(m.Layer),
			CauseMetadata: ConvertToRPCCauseMetadata(m.CauseMetadata),
		})
	}
	return rpcMisconfs
}

// ConvertToRPCLayer returns common.Layer
func ConvertToRPCLayer(layer ftypes.Layer) *common.Layer {
	return &common.Layer{
		Digest:    layer.Digest,
		DiffId:    layer.DiffID,
		CreatedBy: layer.CreatedBy,
	}
}

func ConvertToRPCPolicyMetadata(policy ftypes.PolicyMetadata) *common.PolicyMetadata {
	return &common.PolicyMetadata{
		Id:                 policy.ID,
		AdvId:              policy.AVDID,
		Type:               policy.Type,
		Title:              policy.Title,
		Description:        policy.Description,
		Severity:           policy.Severity,
		RecommendedActions: policy.RecommendedActions,
		References:         policy.References,
	}
}

func ConvertToRPCCauseMetadata(cause ftypes.CauseMetadata) *common.CauseMetadata {
	return &common.CauseMetadata{
		Resource:  cause.Resource,
		Provider:  cause.Provider,
		Service:   cause.Service,
		StartLine: int32(cause.StartLine),
		EndLine:   int32(cause.EndLine),
		Code:      ConvertToRPCCode(cause.Code),
	}
}

// ConvertToRPCDataSource returns common.DataSource
func ConvertToRPCDataSource(ds *dbTypes.DataSource) *common.DataSource {
	if ds == nil {
		return nil
	}
	return &common.DataSource{
		Id:   string(ds.ID),
		Name: ds.Name,
		Url:  ds.URL,
	}
}

// ConvertFromRPCResults converts scanner.Result to types.Result
func ConvertFromRPCResults(rpcResults []*scanner.Result) []types.Result {
	var results []types.Result
	for _, result := range rpcResults {
		results = append(results, types.Result{
			Target:            result.Target,
			Vulnerabilities:   ConvertFromRPCVulns(result.Vulnerabilities),
			Misconfigurations: ConvertFromRPCMisconfs(result.Misconfigurations),
			Class:             types.ResultClass(result.Class),
			Type:              ftypes.TargetType(result.Type),
			Packages:          ConvertFromRPCPkgs(result.Packages),
			CustomResources:   ConvertFromRPCCustomResources(result.CustomResources),
			Secrets:           ConvertFromRPCSecretFindings(result.Secrets),
			Licenses:          ConvertFromRPCLicenses(result.Licenses),
		})
	}
	return results
}

func ConvertFromRPCLicenses(rpcLicenses []*common.DetectedLicense) []types.DetectedLicense {
	var licenses []types.DetectedLicense
	for _, l := range rpcLicenses {
		severity := dbTypes.Severity(l.Severity)
		licenses = append(licenses, types.DetectedLicense{
			Severity:   severity.String(),
			Category:   ConvertFromRPCLicenseCategory(l.Category),
			PkgName:    l.PkgName,
			FilePath:   l.FilePath,
			Name:       l.Name,
			Confidence: float64(l.Confidence),
			Link:       l.Link,
		})
	}
	return licenses
}

func ConvertFromRPCLicenseCategory(rpcCategory common.DetectedLicense_LicenseCategory) ftypes.LicenseCategory {
	if rpcCategory == common.DetectedLicense_UNSPECIFIED {
		return ""
	}
	return ftypes.LicenseCategory(strings.ToLower(rpcCategory.String()))
}

// ConvertFromRPCCustomResources converts array of cache.CustomResource to fanal.CustomResource
func ConvertFromRPCCustomResources(rpcCustomResources []*common.CustomResource) []ftypes.CustomResource {
	var resources []ftypes.CustomResource
	for _, res := range rpcCustomResources {
		resources = append(resources, ftypes.CustomResource{
			Type:     res.Type,
			FilePath: res.FilePath,
			Layer: ftypes.Layer{
				Digest: res.Layer.Digest,
				DiffID: res.Layer.DiffId,
			},
			Data: res.Data,
		})
	}
	return resources
}

func ConvertFromRPCCode(rpcCode *common.Code) ftypes.Code {
	var lines []ftypes.Line
	for _, line := range rpcCode.Lines {
		lines = append(lines, ftypes.Line{
			Number:      int(line.Number),
			Content:     line.Content,
			IsCause:     line.IsCause,
			Annotation:  line.Annotation,
			Truncated:   line.Truncated,
			Highlighted: line.Highlighted,
			FirstCause:  line.FirstCause,
			LastCause:   line.LastCause,
		})
	}
	return ftypes.Code{
		Lines: lines,
	}
}

func ConvertFromRPCSecretFindings(rpcFindings []*common.SecretFinding) []ftypes.SecretFinding {
	var findings []ftypes.SecretFinding
	for _, finding := range rpcFindings {
		findings = append(findings, ftypes.SecretFinding{
			RuleID:    finding.RuleId,
			Category:  ftypes.SecretRuleCategory(finding.Category),
			Severity:  finding.Severity,
			Title:     finding.Title,
			StartLine: int(finding.StartLine),
			EndLine:   int(finding.EndLine),
			Code:      ConvertFromRPCCode(finding.Code),
			Match:     finding.Match,
			Layer: ftypes.Layer{
				Digest:    finding.Layer.Digest,
				DiffID:    finding.Layer.DiffId,
				CreatedBy: finding.Layer.CreatedBy,
			},
		})
	}
	return findings
}

func ConvertFromRPCSecrets(recSecrets []*common.Secret) []ftypes.Secret {
	var secrets []ftypes.Secret
	for _, secret := range recSecrets {
		secrets = append(secrets, ftypes.Secret{
			FilePath: secret.Filepath,
			Findings: ConvertFromRPCSecretFindings(secret.Findings),
		})
	}
	return secrets
}

// ConvertFromRPCVulns converts []*common.Vulnerability to []types.DetectedVulnerability
func ConvertFromRPCVulns(rpcVulns []*common.Vulnerability) []types.DetectedVulnerability {
	var vulns []types.DetectedVulnerability
	for _, vuln := range rpcVulns {
		severity := dbTypes.Severity(vuln.Severity)
		cvssMap := make(dbTypes.VendorCVSS) // This is needed because protobuf generates a map[string]*CVSS type
		for vendor, vendorSeverity := range vuln.Cvss {
			cvssMap[dbTypes.SourceID(vendor)] = dbTypes.CVSS{
				V2Vector: vendorSeverity.V2Vector,
				V3Vector: vendorSeverity.V3Vector,
				V2Score:  vendorSeverity.V2Score,
				V3Score:  vendorSeverity.V3Score,
			}
		}
		vensorSeverityMap := make(dbTypes.VendorSeverity)
		for vendor, vendorSeverity := range vuln.VendorSeverity {
			vensorSeverityMap[dbTypes.SourceID(vendor)] = dbTypes.Severity(vendorSeverity)
		}

		var lastModifiedDate, publishedDate *time.Time
		if vuln.LastModifiedDate != nil {
			lastModifiedDate = lo.ToPtr(vuln.LastModifiedDate.AsTime())
		}
		if vuln.PublishedDate != nil {
			publishedDate = lo.ToPtr(vuln.PublishedDate.AsTime())
		}

		vulns = append(vulns, types.DetectedVulnerability{
			VulnerabilityID:  vuln.VulnerabilityId,
			VendorIDs:        vuln.VendorIds,
			PkgID:            vuln.PkgId,
			PkgName:          vuln.PkgName,
			PkgPath:          vuln.PkgPath,
			InstalledVersion: vuln.InstalledVersion,
			FixedVersion:     vuln.FixedVersion,
			Status:           dbTypes.Status(vuln.Status),
			Vulnerability: dbTypes.Vulnerability{
				Title:            vuln.Title,
				Description:      vuln.Description,
				Severity:         severity.String(),
				CVSS:             cvssMap,
				References:       vuln.References,
				CweIDs:           vuln.CweIds,
				LastModifiedDate: lastModifiedDate,
				PublishedDate:    publishedDate,
				Custom:           vuln.CustomVulnData.AsInterface(),
				VendorSeverity:   vensorSeverityMap,
			},
			Layer:          ConvertFromRPCLayer(vuln.Layer),
			SeveritySource: dbTypes.SourceID(vuln.SeveritySource),
			PrimaryURL:     vuln.PrimaryUrl,
			Custom:         vuln.CustomAdvisoryData.AsInterface(),
			DataSource:     ConvertFromRPCDataSource(vuln.DataSource),
		})
	}
	return vulns
}

// ConvertFromRPCMisconfs converts []*common.DetectedMisconfigurations to []types.DetectedMisconfiguration
func ConvertFromRPCMisconfs(rpcMisconfs []*common.DetectedMisconfiguration) []types.DetectedMisconfiguration {
	var misconfs []types.DetectedMisconfiguration
	for _, rpcMisconf := range rpcMisconfs {
		misconfs = append(misconfs, types.DetectedMisconfiguration{
			Type:          rpcMisconf.Type,
			ID:            rpcMisconf.Id,
			AVDID:         rpcMisconf.AvdId,
			Title:         rpcMisconf.Title,
			Description:   rpcMisconf.Description,
			Message:       rpcMisconf.Message,
			Namespace:     rpcMisconf.Namespace,
			Query:         rpcMisconf.Query,
			Resolution:    rpcMisconf.Resolution,
			Severity:      rpcMisconf.Severity.String(),
			PrimaryURL:    rpcMisconf.PrimaryUrl,
			References:    rpcMisconf.References,
			Status:        types.MisconfStatus(rpcMisconf.Status),
			Layer:         ConvertFromRPCLayer(rpcMisconf.Layer),
			CauseMetadata: ConvertFromRPCCauseMetadata(rpcMisconf.CauseMetadata),
		})
	}
	return misconfs
}

// ConvertFromRPCLayer converts *common.Layer to fanal.Layer
func ConvertFromRPCLayer(rpcLayer *common.Layer) ftypes.Layer {
	if rpcLayer == nil {
		return ftypes.Layer{}
	}
	return ftypes.Layer{
		Digest:    rpcLayer.Digest,
		DiffID:    rpcLayer.DiffId,
		CreatedBy: rpcLayer.CreatedBy,
	}
}

func ConvertFromRPCPolicyMetadata(rpcPolicy *common.PolicyMetadata) ftypes.PolicyMetadata {
	if rpcPolicy == nil {
		return ftypes.PolicyMetadata{}
	}

	return ftypes.PolicyMetadata{
		ID:                 rpcPolicy.Id,
		AVDID:              rpcPolicy.AdvId,
		Type:               rpcPolicy.Type,
		Title:              rpcPolicy.Title,
		Description:        rpcPolicy.Description,
		Severity:           rpcPolicy.Severity,
		RecommendedActions: rpcPolicy.RecommendedActions,
		References:         rpcPolicy.References,
	}
}

func ConvertFromRPCCauseMetadata(rpcCause *common.CauseMetadata) ftypes.CauseMetadata {
	if rpcCause == nil {
		return ftypes.CauseMetadata{}
	}
	return ftypes.CauseMetadata{
		Resource:  rpcCause.Resource,
		Provider:  rpcCause.Provider,
		Service:   rpcCause.Service,
		StartLine: int(rpcCause.StartLine),
		EndLine:   int(rpcCause.EndLine),
		Code:      ConvertFromRPCCode(rpcCause.Code),
	}
}

// ConvertFromRPCOS converts common.OS to fanal.OS
func ConvertFromRPCOS(rpcOS *common.OS) ftypes.OS {
	if rpcOS == nil {
		return ftypes.OS{}
	}
	return ftypes.OS{
		Family:   ftypes.OSType(rpcOS.Family),
		Name:     rpcOS.Name,
		Eosl:     rpcOS.Eosl,
		Extended: rpcOS.Extended,
	}
}

// ConvertFromRPCRepository converts common.Repository to fanal.Repository
func ConvertFromRPCRepository(rpcRepo *common.Repository) *ftypes.Repository {
	if rpcRepo == nil {
		return nil
	}
	return &ftypes.Repository{
		Family:  ftypes.OSType(rpcRepo.Family),
		Release: rpcRepo.Release,
	}
}

// ConvertFromRPCDataSource converts *common.DataSource to *dbTypes.DataSource
func ConvertFromRPCDataSource(ds *common.DataSource) *dbTypes.DataSource {
	if ds == nil {
		return nil
	}
	return &dbTypes.DataSource{
		ID:   dbTypes.SourceID(ds.Id),
		Name: ds.Name,
		URL:  ds.Url,
	}
}

// ConvertFromRPCPackageInfos converts common.PackageInfo to fanal.PackageInfo
func ConvertFromRPCPackageInfos(rpcPkgInfos []*common.PackageInfo) []ftypes.PackageInfo {
	var pkgInfos []ftypes.PackageInfo
	for _, rpcPkgInfo := range rpcPkgInfos {
		pkgInfos = append(pkgInfos, ftypes.PackageInfo{
			FilePath: rpcPkgInfo.FilePath,
			Packages: ConvertFromRPCPkgs(rpcPkgInfo.Packages),
		})
	}
	return pkgInfos
}

// ConvertFromRPCApplications converts common.Application to fanal.Application
func ConvertFromRPCApplications(rpcApps []*common.Application) []ftypes.Application {
	var apps []ftypes.Application
	for _, rpcApp := range rpcApps {
		apps = append(apps, ftypes.Application{
			Type:      ftypes.LangType(rpcApp.Type),
			FilePath:  rpcApp.FilePath,
			Libraries: ConvertFromRPCPkgs(rpcApp.Libraries),
		})
	}
	return apps
}

// ConvertFromRPCMisconfigurations converts common.Misconfiguration to fanal.Misconfiguration
func ConvertFromRPCMisconfigurations(rpcMisconfs []*common.Misconfiguration) []ftypes.Misconfiguration {
	var misconfs []ftypes.Misconfiguration
	for _, rpcMisconf := range rpcMisconfs {
		misconfs = append(misconfs, ftypes.Misconfiguration{
			FileType:   ftypes.ConfigType(rpcMisconf.FileType),
			FilePath:   rpcMisconf.FilePath,
			Successes:  ConvertFromRPCMisconfResults(rpcMisconf.Successes),
			Warnings:   ConvertFromRPCMisconfResults(rpcMisconf.Warnings),
			Failures:   ConvertFromRPCMisconfResults(rpcMisconf.Failures),
			Exceptions: ConvertFromRPCMisconfResults(rpcMisconf.Exceptions),
			Layer:      ftypes.Layer{},
		})
	}
	return misconfs
}

// ConvertFromRPCMisconfResults converts common.MisconfResult to fanal.MisconfResult
func ConvertFromRPCMisconfResults(rpcResults []*common.MisconfResult) []ftypes.MisconfResult {
	var results []ftypes.MisconfResult
	for _, r := range rpcResults {
		results = append(results, ftypes.MisconfResult{
			Namespace:      r.Namespace,
			Message:        r.Message,
			PolicyMetadata: ConvertFromRPCPolicyMetadata(r.PolicyMetadata),
			CauseMetadata:  ConvertFromRPCCauseMetadata(r.CauseMetadata),
		})
	}
	return results
}

// ConvertFromRPCPutArtifactRequest converts cache.PutArtifactRequest to fanal.PutArtifactRequest
func ConvertFromRPCPutArtifactRequest(req *cache.PutArtifactRequest) ftypes.ArtifactInfo {
	return ftypes.ArtifactInfo{
		SchemaVersion:   int(req.ArtifactInfo.SchemaVersion),
		Architecture:    req.ArtifactInfo.Architecture,
		Created:         req.ArtifactInfo.Created.AsTime(),
		DockerVersion:   req.ArtifactInfo.DockerVersion,
		OS:              req.ArtifactInfo.Os,
		HistoryPackages: ConvertFromRPCPkgs(req.ArtifactInfo.HistoryPackages),
	}
}

// ConvertFromRPCPutBlobRequest returns ftypes.BlobInfo
func ConvertFromRPCPutBlobRequest(req *cache.PutBlobRequest) ftypes.BlobInfo {
	return ftypes.BlobInfo{
		SchemaVersion:     int(req.BlobInfo.SchemaVersion),
		Digest:            req.BlobInfo.Digest,
		DiffID:            req.BlobInfo.DiffId,
		OS:                ConvertFromRPCOS(req.BlobInfo.Os),
		Repository:        ConvertFromRPCRepository(req.BlobInfo.Repository),
		PackageInfos:      ConvertFromRPCPackageInfos(req.BlobInfo.PackageInfos),
		Applications:      ConvertFromRPCApplications(req.BlobInfo.Applications),
		Misconfigurations: ConvertFromRPCMisconfigurations(req.BlobInfo.Misconfigurations),
		OpaqueDirs:        req.BlobInfo.OpaqueDirs,
		WhiteoutFiles:     req.BlobInfo.WhiteoutFiles,
		CustomResources:   ConvertFromRPCCustomResources(req.BlobInfo.CustomResources),
		Secrets:           ConvertFromRPCSecrets(req.BlobInfo.Secrets),
	}
}

// ConvertToRPCOS returns common.OS
func ConvertToRPCOS(fos ftypes.OS) *common.OS {
	return &common.OS{
		Family:   string(fos.Family),
		Name:     fos.Name,
		Eosl:     fos.Eosl,
		Extended: fos.Extended,
	}
}

// ConvertToRPCRepository returns common.Repository
func ConvertToRPCRepository(repo *ftypes.Repository) *common.Repository {
	if repo == nil {
		return nil
	}
	return &common.Repository{
		Family:  string(repo.Family),
		Release: repo.Release,
	}
}

// ConvertToRPCArtifactInfo returns PutArtifactRequest
func ConvertToRPCArtifactInfo(imageID string, imageInfo ftypes.ArtifactInfo) *cache.PutArtifactRequest {

	t := timestamppb.New(imageInfo.Created)
	if err := t.CheckValid(); err != nil {
		log.Logger.Warnf("invalid timestamp: %s", err)
	}

	return &cache.PutArtifactRequest{
		ArtifactId: imageID,
		ArtifactInfo: &cache.ArtifactInfo{
			SchemaVersion:   int32(imageInfo.SchemaVersion),
			Architecture:    imageInfo.Architecture,
			Created:         t,
			DockerVersion:   imageInfo.DockerVersion,
			Os:              imageInfo.OS,
			HistoryPackages: ConvertToRPCPkgs(imageInfo.HistoryPackages),
		},
	}
}

// ConvertToRPCPutBlobRequest returns PutBlobRequest
func ConvertToRPCPutBlobRequest(diffID string, blobInfo ftypes.BlobInfo) *cache.PutBlobRequest {
	var packageInfos []*common.PackageInfo
	for _, pkgInfo := range blobInfo.PackageInfos {
		packageInfos = append(packageInfos, &common.PackageInfo{
			FilePath: pkgInfo.FilePath,
			Packages: ConvertToRPCPkgs(pkgInfo.Packages),
		})
	}

	var applications []*common.Application
	for _, app := range blobInfo.Applications {
		applications = append(applications, &common.Application{
			Type:      string(app.Type),
			FilePath:  app.FilePath,
			Libraries: ConvertToRPCPkgs(app.Libraries),
		})
	}

	var misconfigurations []*common.Misconfiguration
	for _, m := range blobInfo.Misconfigurations {
		misconfigurations = append(misconfigurations, &common.Misconfiguration{
			FileType:   string(m.FileType),
			FilePath:   m.FilePath,
			Successes:  ConvertToMisconfResults(m.Successes),
			Warnings:   ConvertToMisconfResults(m.Warnings),
			Failures:   ConvertToMisconfResults(m.Failures),
			Exceptions: ConvertToMisconfResults(m.Exceptions),
		})

	}

	var customResources []*common.CustomResource
	for _, res := range blobInfo.CustomResources {
		data, err := structpb.NewValue(res.Data)
		if err != nil {

		} else {
			customResources = append(customResources, &common.CustomResource{
				Type:     res.Type,
				FilePath: res.FilePath,
				Layer: &common.Layer{
					Digest: res.Layer.Digest,
					DiffId: res.Layer.DiffID,
				},
				Data: data,
			})
		}
	}

	return &cache.PutBlobRequest{
		DiffId: diffID,
		BlobInfo: &cache.BlobInfo{
			SchemaVersion:     ftypes.BlobJSONSchemaVersion,
			Digest:            blobInfo.Digest,
			DiffId:            blobInfo.DiffID,
			Os:                ConvertToRPCOS(blobInfo.OS),
			Repository:        ConvertToRPCRepository(blobInfo.Repository),
			PackageInfos:      packageInfos,
			Applications:      applications,
			Misconfigurations: misconfigurations,
			OpaqueDirs:        blobInfo.OpaqueDirs,
			WhiteoutFiles:     blobInfo.WhiteoutFiles,
			CustomResources:   customResources,
			Secrets:           ConvertToRPCSecrets(blobInfo.Secrets),
		},
	}
}

// ConvertToMisconfResults returns common.MisconfResult
func ConvertToMisconfResults(results []ftypes.MisconfResult) []*common.MisconfResult {
	var rpcResults []*common.MisconfResult
	for _, r := range results {
		rpcResults = append(rpcResults, &common.MisconfResult{
			Namespace:      r.Namespace,
			Message:        r.Message,
			PolicyMetadata: ConvertToRPCPolicyMetadata(r.PolicyMetadata),
			CauseMetadata:  ConvertToRPCCauseMetadata(r.CauseMetadata),
		})
	}
	return rpcResults
}

// ConvertToMissingBlobsRequest returns MissingBlobsRequest object
func ConvertToMissingBlobsRequest(imageID string, layerIDs []string) *cache.MissingBlobsRequest {
	return &cache.MissingBlobsRequest{
		ArtifactId: imageID,
		BlobIds:    layerIDs,
	}
}

// ConvertToRPCScanResponse converts types.Result to ScanResponse
func ConvertToRPCScanResponse(results types.Results, fos ftypes.OS) *scanner.ScanResponse {
	var rpcResults []*scanner.Result
	for _, result := range results {
		rpcResults = append(rpcResults, &scanner.Result{
			Target:            result.Target,
			Class:             string(result.Class),
			Type:              string(result.Type),
			Vulnerabilities:   ConvertToRPCVulns(result.Vulnerabilities),
			Misconfigurations: ConvertToRPCMisconfs(result.Misconfigurations),
			Packages:          ConvertToRPCPkgs(result.Packages),
			CustomResources:   ConvertToRPCCustomResources(result.CustomResources),
			Secrets:           ConvertToRPCSecretFindings(result.Secrets),
			Licenses:          ConvertToRPCLicenses(result.Licenses),
		})
	}

	return &scanner.ScanResponse{
		Os:      ConvertToRPCOS(fos),
		Results: rpcResults,
	}
}

func ConvertToRPCLicenses(licenses []types.DetectedLicense) []*common.DetectedLicense {
	var rpcLicenses []*common.DetectedLicense
	for _, l := range licenses {
		severity, err := dbTypes.NewSeverity(l.Severity)
		if err != nil {
			log.Logger.Warn(err)
		}
		rpcLicenses = append(rpcLicenses, &common.DetectedLicense{
			Severity:   common.Severity(severity),
			Category:   ConvertToRPCLicenseCategory(l.Category),
			PkgName:    l.PkgName,
			FilePath:   l.FilePath,
			Name:       l.Name,
			Confidence: float32(l.Confidence),
			Link:       l.Link,
		})
	}

	return rpcLicenses
}

func ConvertToRPCLicenseCategory(category ftypes.LicenseCategory) common.DetectedLicense_LicenseCategory {
	rpcCategory, ok := common.DetectedLicense_LicenseCategory_value[strings.ToUpper(string(category))]
	if !ok {
		return common.DetectedLicense_UNSPECIFIED
	}
	return common.DetectedLicense_LicenseCategory(rpcCategory)
}

func ConvertToDeleteBlobsRequest(blobIDs []string) *cache.DeleteBlobsRequest {
	return &cache.DeleteBlobsRequest{BlobIds: blobIDs}
}

func ConvertFromDeleteBlobsRequest(deleteBlobsRequest *cache.DeleteBlobsRequest) []string {
	if deleteBlobsRequest == nil {
		return []string{}
	}
	return deleteBlobsRequest.GetBlobIds()
}