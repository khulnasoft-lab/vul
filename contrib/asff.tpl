[
{{- $t_first := true -}}
{{- range . -}}
{{- $target := .Target -}}
{{- range .Vulnerabilities -}}
{{- if $t_first -}}
  {{- $t_first = false -}}
{{- else -}}
  ,
{{- end -}}
{{- $vulProductSev := 0 -}}
{{- $vulNormalizedSev := 0 -}}
{{- if eq .Severity "LOW" -}}
    {{- $vulProductSev = 1 -}}
    {{- $vulNormalizedSev = 10 -}}
  {{- else if eq .Severity "MEDIUM" -}}
    {{- $vulProductSev = 4 -}}
    {{- $vulNormalizedSev = 40 -}}
  {{- else if eq .Severity "HIGH" -}}
    {{- $vulProductSev = 7 -}}
    {{- $vulNormalizedSev = 70 -}}
  {{- else if eq .Severity "CRITICAL" -}}
    {{- $vulProductSev = 9 -}}
    {{- $vulNormalizedSev = 90 -}}
{{- end }}
{{- $description := .Description -}}
{{- if gt (len $description ) 1021 -}}
    {{- $description = (slice $description 0 1021) | printf "%v .." -}}
{{- end}}
    {
        "SchemaVersion": "2018-10-08",
        "Id": "{{ $target }}/{{ .VulnerabilityID }}",
        "ProductArn": "arn:aws:securityhub:{{ getEnv "AWS_REGION" }}::product/khulnasoft/aquasecurity",
        "GeneratorId": "Vul",
        "AwsAccountId": "{{ getEnv "AWS_ACCOUNT_ID" }}",
        "Types": [ "Software and Configuration Checks/Vulnerabilities/CVE" ],
        "CreatedAt": "{{ getCurrentTime }}",
        "UpdatedAt": "{{ getCurrentTime }}",
        "Severity": {
            "Product": {{ $vulProductSev }},
            "Normalized": {{ $vulNormalizedSev }}
        },
        "Title": "Vul found a vulnerability to {{ .VulnerabilityID }} in container {{ $target }}",
        "Description": {{ escapeString $description | printf "%q" }},
        "Remediation": {
            "Recommendation": {
                "Text": "More information on this vulnerability is provided in the hyperlink",
                "Url": "{{ .PrimaryURL }}"
            }
        },
        "ProductFields": { "Product Name": "Vul" },
        "Resources": [
            {
                "Type": "Container",
                "Id": "{{ $target }}",
                "Partition": "aws",
                "Region": "{{ getEnv "AWS_REGION" }}",
                "Details": {
                    "Container": { "ImageName": "{{ $target }}" },
                    "Other": {
                        "CVE ID": "{{ .VulnerabilityID }}",
                        "CVE Title": {{ .Title | printf "%q" }},
                        "PkgName": "{{ .PkgName }}",
                        "Installed Package": "{{ .InstalledVersion }}",
                        "Patched Package": "{{ .FixedVersion }}",
                        "NvdCvssScoreV3": "{{ (index .CVSS "nvd").V3Score }}",
                        "NvdCvssVectorV3": "{{ (index .CVSS "nvd").V3Vector }}",
                        "NvdCvssScoreV2": "{{ (index .CVSS "nvd").V2Score }}",
                        "NvdCvssVectorV2": "{{ (index .CVSS "nvd").V2Vector }}"
                    }
                }
            }
        ],
        "RecordState": "ACTIVE"
    }
   {{- end -}}
  {{- end }}
]