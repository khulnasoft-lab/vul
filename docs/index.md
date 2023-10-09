---
hide:
- toc
---
![logo](imgs/logo.png){ align=right }

# Vul Documentation

👋 Welcome to Vul Documentation! To help you get around, please notice the different sections at the top global menu:

- You are currently in the [Getting Started] section where you can find general information and help with first steps.
- In the [Tutorials] section you can find step-by-step guides that help you accomplish specific tasks.
- In the [Docs] section you can find the complete reference documentation for all of the different features and settings that Vul has to offer.
- In the [Ecosystem] section you can find how Vul works together with other tools and applications that you might already use.
- In the [Contributing] section you can find technical developer documentation and contribution guidelines.

# About Vul

Vul ([pronunciation][pronunciation]) is a comprehensive and versatile security scanner. Vul has *scanners* that look for security issues, and *targets* where it can find those issues.

Targets (what Vul can scan):

- Container Image
- Filesystem
- Git Repository (remote)
- Virtual Machine Image
- Kubernetes
- AWS

Scanners (what Vul can find there):

- OS packages and software dependencies in use (SBOM)
- Known vulnerabilities (CVEs)
- IaC issues and misconfigurations
- Sensitive information and secrets
- Software licenses

Vul supports most popular programming languages, operating systems, and platforms. For a complete list, see the [Scanning Coverage] page.

To learn more, go to the [Vul homepage][homepage] for feature highlights, or to the [Documentation site][Docs] for detailed information.

## Quick Start

### Get Vul

Vul is available in most common distribution channels. The complete list of installation options is available in the [Installation] page. Here are a few popular examples:

- `brew install vul`
- `docker run khulnasoft/vul`
- Download binary from <https://github.com/khulnasoft-lab/vul/releases/latest/>
- See [Installation] for more

Vul is integrated with many popular platforms and applications. The complete list of integrations is available in the [Ecosystem] page. Here are a few popular options examples:

- [GitHub Actions](https://github.com/khulnasoft-lab/vul-action)
- [Kubernetes operator](https://github.com/khulnasoft-lab/vul-operator)
- [VS Code plugin](https://github.com/khulnasoft-lab/vul-vscode-extension)
- See [Ecosystem] for more

### General usage

```bash
vul <target> [--scanners <scanner1,scanner2>] <subject>
```

Examples:

```bash
vul image python:3.4-alpine
```

<details>
<summary>Result</summary>

<figure style="text-align: center">
  <video width="1000" autoplay muted controls loop>
    <source src="https://user-images.githubusercontent.com/1161307/171013513-95f18734-233d-45d3-aaf5-d6aec687db0e.mov" type="video/mp4" />
  </video>
  <figcaption>Demo: Vulnerability Detection</figcaption>
</figure>

</details>

```bash
vul fs --scanners vuln,secret,config myproject/
```

<details>
<summary>Result</summary>

<figure style="text-align: center">
  <video width="1000" autoplay muted controls loop>
    <source src="https://user-images.githubusercontent.com/1161307/171013917-b1f37810-f434-465c-b01a-22de036bd9b3.mov" type="video/mp4" />
  </video>
  <figcaption>Demo: Misconfiguration Detection</figcaption>
</figure>

</details>

```bash
vul k8s --report summary cluster
```

<details>
<summary>Result</summary>

<figure style="text-align: center">
  <img src="imgs/secret-demo.gif" width="1000">
  <figcaption>Demo: Secret Detection</figcaption>
</figure>

</details>

# Want more? Check out Khulnasoft

If you liked Vul, you will love Khulnasoft which builds on top of Vul to provide even more enhanced capabilities for a complete security management offering.  
You can find a high level comparison table specific to Vul users [here](https://github.com/khulnasoft-lab/resources/blob/main/vul-khulnasoft.md).  
In addition check out the <https://khulnasoft.com> website for more information about our products and services.
If you'd like to contact Khulnasoft or request a demo, please use this form: <https://www.khulnasoft.com/demo>

---

Vul is an [Khulnasoft Security][khulnasoft] open source project.  
Learn about our open source work and portfolio [here][oss].  
Contact us about any matter by opening a GitHub Discussion [here][discussions]

[Ecosystem]: ./ecosystem/index.md
[Installation]: getting-started/installation.md
[pronunciation]: #how-to-pronounce-the-name-vul
[Scanning Coverage]: ./docs/coverage/index.md

[khulnasoft]: https://khulnasoft.com
[oss]: https://www.khulnasoft.com/products/open-source-projects/
[discussions]: https://github.com/khulnasoft-lab/vul/discussions

[homepage]: https://vul.dev
[Tutorials]: ./tutorials/overview
[Docs]: ./docs
[Getting Started]: ./
[Contributing]: ./community/contribute/issue
