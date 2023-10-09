# CBL-Mariner
Vul supports the following scanners for OS packages.

| Version          | SBOM  | Vulnerability | License |
| ---------------- | :---: | :-----------: | :-----: |
| 1.0              |   ✔   |       ✔       |    ✔    |
| 1.0 (Distroless) |   ✔   |       ✔       |         |
| 2.0              |   ✔   |       ✔       |    ✔    |
| 2.0 (Distroless) |   ✔   |       ✔       |         |


The following table provides an outline of the targets Vul supports.

| Version | Container image | Virtual machine |     Arch     |
| ------- | :-------------: | :-------------: | :----------: |
| 1.0     |        ✔        |        ✔        | amd64, arm64 |
| 2.0     |        ✔        |        ✔        | amd64, arm64 |

The table below outlines the features offered by Vul.

|               Feature                | Supported |
|:------------------------------------:|:---------:|
|    Detect unfixed vulnerabilities    |     ✓     |
| [Dependency graph][dependency-graph] |     ✓     |

## SBOM
Vul detects packages that have been installed through package managers such as `dnf` and `yum`.

## Vulnerability
CBL-Mariner offers its own security advisories, and these are utilized when scanning CBL-Mariner for vulnerabilities.

### Data Source
See [here](../../scanner/vulnerability.md#data-sources).

### Fixed Version
Vul takes fixed versions from [CBL-Mariner OVAL][oval].

### Severity
Vul calculates the severity of an issue based on the severity provided in [CBL-Mariner OVAL][oval].

### Status
Vul supports the following [vulnerability statuses] for CBL-Mariner.

|       Status        | Supported |
| :-----------------: | :-------: |
|        Fixed        |     ✓     |
|      Affected       |     ✓     |
| Under Investigation |           |
|    Will Not Fix     |           |
|    Fix Deferred     |           |
|     End of Life     |           |


## License
Vul identifies licenses by examining the metadata of RPM packages.

!!! note
    License detection is not supported for CBL-Mariner Distroless.


[dependency-graph]: ../../configuration/reporting.md#show-origins-of-vulnerable-dependencies
[cbl-mariner]: https://github.com/microsoft/CBL-Mariner

[oval]: https://github.com/microsoft/CBL-MarinerVulnerabilityData/

[vulnerability statuses]: ../../configuration/filtering.md#by-status
