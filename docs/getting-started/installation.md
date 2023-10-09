# Installing Vul

In this section you will find an aggregation of the different ways to install Vul. installations are listed as either "official" or "community". Official integrations are developed by the core Vul team and supported by it. Community integrations are integrations developed by the community, and collected here for your convenience. For support or questions about community integrations, please contact the original developers.

## Install using Package Manager

### RHEL/CentOS (Official)

=== "Repository"
    Add repository setting to `/etc/yum.repos.d`.

    ``` bash
    RELEASE_VERSION=$(grep -Po '(?<=VERSION_ID=")[0-9]' /etc/os-release)
    cat << EOF | sudo tee -a /etc/yum.repos.d/vul.repo
    [vul]
    name=Vul repository
    baseurl=https://khulnasoft-lab.github.io/vul-repo/rpm/releases/$RELEASE_VERSION/\$basearch/
    gpgcheck=1
    enabled=1
    gpgkey=https://khulnasoft-lab.github.io/vul-repo/rpm/public.key
    EOF
    sudo yum -y update
    sudo yum -y install vul
    ```

=== "RPM"

    ``` bash
    rpm -ivh https://github.com/khulnasoft-lab/vul/releases/download/{{ git.tag }}/vul_{{ git.tag[1:] }}_Linux-64bit.rpm
    ```

### Debian/Ubuntu (Official)

=== "Repository"
    Add repository setting to `/etc/apt/sources.list.d`.

    ``` bash
    sudo apt-get install wget apt-transport-https gnupg lsb-release
    wget -qO - https://khulnasoft-lab.github.io/vul-repo/deb/public.key | gpg --dearmor | sudo tee /usr/share/keyrings/vul.gpg > /dev/null
    echo "deb [signed-by=/usr/share/keyrings/vul.gpg] https://khulnasoft-lab.github.io/vul-repo/deb $(lsb_release -sc) main" | sudo tee -a /etc/apt/sources.list.d/vul.list
    sudo apt-get update
    sudo apt-get install vul
    ```

=== "DEB"

    ``` bash
    wget https://github.com/khulnasoft-lab/vul/releases/download/{{ git.tag }}/vul_{{ git.tag[1:] }}_Linux-64bit.deb
    sudo dpkg -i vul_{{ git.tag[1:] }}_Linux-64bit.deb
    ```

### Homebrew (Official)

Homebrew for MacOS and Linux.

```bash
brew install vul
```

### Arch Linux (Community)

Arch Community Package Manager.

```bash
pacman -S vul
```

References: 
- <https://archlinux.org/packages/community/x86_64/vul/>
- <https://github.com/archlinux/svntogit-community/blob/packages/vul/trunk/PKGBUILD>


### MacPorts (Community)

[MacPorts](https://www.macports.org) for MacOS.

```bash
sudo port install vul
```

References:
- <https://ports.macports.org/port/vul/details/>

### Nix/NixOS (Community)

Nix package manager for Linux and MacOS.

=== "Command line"

`nix-env --install -A nixpkgs.vul`

=== "Configuration"

```nix
  # your other config ...
  environment.systemPackages = with pkgs; [
    # your other packages ...
    vul
  ];
```

=== "Home Manager"

```nix
  # your other config ...
  home.packages = with pkgs; [
    # your other packages ...
    vul
  ];
```

References: 
-  <https://github.com/NixOS/nixpkgs/blob/master/pkgs/tools/admin/vul/default.nix>

## Install from GitHub Release (Official)

### Download Binary

1. Download the file for your operating system/architecture from [GitHub Release assets](https://github.com/khulnasoft-lab/vul/releases/tag/{{ git.tag }}) (`curl -LO https://url.to/vul.tar.gz`).  
2. Unpack the downloaded archive (`tar -xzf ./vul.tar.gz`).
3. Put the binary somewhere in your `$PATH` (e.g `mv ./vul /usr/local/bin/`).
4. Make sure the binary has execution bit turned on (`chmod +x ./vul`).

### Install Script

The process above can be automated by the following script:

```bash
curl -sfL https://raw.githubusercontent.com/khulnasoft-lab/vul/main/contrib/install.sh | sh -s -- -b /usr/local/bin {{ git.tag }}
```

### Install from source

```bash
git clone --depth 1 --branch {{ git.tag }} https://github.com/khulnasoft-lab/vul
cd vul
go install ./cmd/vul
```

## Use container image

1. Pull Vul image (`docker pull khulnasoft/vul:{{ git.tag[1:] }}`)
   2. It is advisable to mount a consistent [cache dir](../docs/configuration/cache.md) on the host into the Vul container.
3. For scanning container images with Vul, mount `docker.sock` from the host into the Vul container.

Example:

``` bash
docker run -v /var/run/docker.sock:/var/run/docker.sock -v $HOME/Library/Caches:/root/.cache/ khulnasoft/vul:{{ git.tag[1:] }} image python:3.4-alpine
```

Registry | Repository | Link | Supportability
Docker Hub | `docker.io/khulnasoft/vul` | https://hub.docker.com/r/khulnasoft/vul | Official
GitHub Container Registry (GHCR) | `ghcr.io/khulnasoft-lab/vul` | https://github.com/orgs/khulnasoft-lab/packages/container/package/vul | Official
AWS Elastic Container Registry (ECR) | `public.ecr.aws/khulnasoft-lab/vul` | https://gallery.ecr.aws/khulnasoft-lab/vul | Official

## Other Tools to use and deploy Vul

For additional tools and ways to install and use Vul in different environments such as in IDE, Kubernetes or CI/CD, see [Ecosystem section](../ecosystem/index.md).
