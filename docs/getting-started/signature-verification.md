# Signature Verification

## Verifying a Cosign signature
All binaries and container images are signed by [Cosign](https://github.com/sigstore/cosign).

You need the following tool:

- [Cosign](https://docs.sigstore.dev/cosign/installation/)

### Verifying signed container images
1. Use the following command for keyless [verification](https://docs.sigstore.dev/cosign/verify/):
   ```shell
   cosign verify khulnasoft/vul:<version> \
   --certificate-identity-regexp 'https://github\.com/khulnasoft-lab/vul/\.github/workflows/.+' \
   --certificate-oidc-issuer "https://token.actions.githubusercontent.com"
   ```
   
2. You should get the following output
   ```shell
   Verification for index.docker.io/khulnasoft/vul:latest --
   The following checks were performed on each of these signatures:
     - The cosign claims were validated
     - Existence of the claims in the transparency log was verified offline
     - The code-signing certificate was verified using trusted certificate authority certificates

     ....
   ```

### Verifying signed binaries

1. Download the required tarball, associated signature and certificate files
2. Use the following command for keyless verification:
   ```shell
   cosign verify-blob <path to binray> \
   --certificate <path to cert> \
   --signature <path to sig> \
   --certificate-identity-regexp 'https://github\.com/khulnasoft-lab/vul/\.github/workflows/.+' \
   --certificate-oidc-issuer "https://token.actions.githubusercontent.com"
   ```
3. You should get the following output
   ```
   Verified OK
   ```
   
For example:

```shell
$ wget "https://github.com/khulnasoft-lab/vul/releases/download/v0.45.0/vul_0.45.0_Linux-32bit.tar.gz"
$ wget "https://github.com/khulnasoft-lab/vul/releases/download/v0.45.0/vul_0.45.0_Linux-32bit.tar.gz.pem"
$ wget "https://github.com/khulnasoft-lab/vul/releases/download/v0.45.0/vul_0.45.0_Linux-32bit.tar.gz.sig"
$ cosign verify-blob vul_0.45.0_Linux-32bit.tar.gz \
  --certificate vul_0.45.0_Linux-32bit.tar.gz.pem \
  --signature vul_0.45.0_Linux-32bit.tar.gz.sig \
  --certificate-identity-regexp 'https://github\.com/khulnasoft-lab/vul/\.github/workflows/.+' \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com" 
  
Vetified OK
```

## Verifying a GPG signature

RPM and Deb packages are also signed by GPG.

### Verifying RPM

The public key downloaded [here](https://khulnasoft-lab.github.io/vul-repo/rpm/public.key).

1. Download the public key
   ```shell
   curl https://khulnasoft-lab.github.io/vul-repo/rpm/public.key \ 
   --output pub.key
   ```
2. Import the key
   ```shell
   rpm --import pub.key
   ```
3. Verify that the key has been imported
   ```shell
   rpm -q --queryformat "%{SUMMARY}\n" $(rpm -q gpg-pubkey)
   ```
   You should get the following output
   ```shell
   gpg(vul)
   ```
   
4. Download the required binary
   ```shell
   curl -L https://github.com/khulnasoft-lab/vul/releases/download/<version>/<file name>.rpm \
   --output vul.rpm
   ```
5. Check the binary with the following command
   ```shell
   rpm -K vul.rpm
   ```
   You should get the following output
   ```shell
   vul.rpm: digests signatures OK
   ```

