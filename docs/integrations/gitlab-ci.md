# GitLab CI

```
$ cat .gitlab-ci.yml
stages:
  - test

vul:
  stage: test
  image: docker:stable
  services:
    - name: docker:dind
      entrypoint: ["env", "-u", "DOCKER_HOST"]
      command: ["dockerd-entrypoint.sh"]
  variables:
    DOCKER_HOST: tcp://docker:2375/
    DOCKER_DRIVER: overlay2
    # See https://github.com/docker-library/docker/pull/166
    DOCKER_TLS_CERTDIR: ""
    IMAGE: vul-ci-test:$CI_COMMIT_SHA
  before_script:
    - export VUL_VERSION=$(wget -qO - "https://api.github.com/repos/khulnasoft-lab/vul/releases/latest" | grep '"tag_name":' | sed -E 's/.*"v([^"]+)".*/\1/')
    - echo $VUL_VERSION
    - wget --no-verbose https://github.com/khulnasoft-lab/vul/releases/download/v${VUL_VERSION}/vul_${VUL_VERSION}_Linux-64bit.tar.gz -O - | tar -zxvf -
  allow_failure: true
  script:
    # Build image
    - docker build -t $IMAGE .
    # Build report
    - ./vul --exit-code 0 --cache-dir .vulcache/ --no-progress --format template --template "@contrib/gitlab.tpl" -o gl-container-scanning-report.json $IMAGE
    # Print report
    - ./vul --exit-code 0 --cache-dir .vulcache/ --no-progress --severity HIGH $IMAGE
    # Fail on severe vulnerabilities
    - ./vul --exit-code 1 --cache-dir .vulcache/ --severity CRITICAL --no-progress $IMAGE
  cache:
    paths:
      - .vulcache/
  # Enables https://docs.gitlab.com/ee/user/application_security/container_scanning/ (Container Scanning report is available on GitLab EE Ultimate or GitLab.com Gold)
  artifacts:
    reports:
      container_scanning: gl-container-scanning-report.json
```

[Example][example]
[Repository][repository]

### GitLab CI using Vul container

To scan a previously built image that has already been pushed into the
GitLab container registry the following CI job manifest can be used.
Note that `entrypoint` needs to be unset for the `script` section to work.
In case of a non-public GitLab project Vul additionally needs to
authenticate to the registry to be able to pull your application image.
Finally, it is not necessary to clone the project repo as we only work
with the container image.

```yaml
container_scanning:
  image:
    name: docker.io/khulnasoft/vul:latest
    entrypoint: [""]
  variables:
    # No need to clone the repo, we exclusively work on artifacts.  See
    # https://docs.gitlab.com/ee/ci/runners/README.html#git-strategy
    GIT_STRATEGY: none
    VUL_USERNAME: "$CI_REGISTRY_USER"
    VUL_PASSWORD: "$CI_REGISTRY_PASSWORD"
    VUL_AUTH_URL: "$CI_REGISTRY"
    FULL_IMAGE_NAME: $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_SLUG
  script:
    - vul --version
    # cache cleanup is needed when scanning images with the same tags, it does not remove the database
    - time vul image --clear-cache
    # update vulnerabilities db
    - time vul --download-db-only --no-progress --cache-dir .vulcache/
    # Builds report and puts it in the default workdir $CI_PROJECT_DIR, so `artifacts:` can take it from there
    - time vul --exit-code 0 --cache-dir .vulcache/ --no-progress --format template --template "@/contrib/gitlab.tpl"
        --output "$CI_PROJECT_DIR/gl-container-scanning-report.json" "$FULL_IMAGE_NAME"
    # Prints full report
    - time vul --exit-code 0 --cache-dir .vulcache/ --no-progress "$FULL_IMAGE_NAME"
    # Fails on high and critical vulnerabilities
    - time vul --exit-code 1 --cache-dir .vulcache/ --severity CRITICAL --no-progress "$FULL_IMAGE_NAME"
  cache:
    paths:
      - .vulcache/
  # Enables https://docs.gitlab.com/ee/user/application_security/container_scanning/ (Container Scanning report is available on GitLab EE Ultimate or GitLab.com Gold)
  artifacts:
    when:                          always
    reports:
      container_scanning:          gl-container-scanning-report.json
  tags:
    - docker-runner
```

[example]: https://gitlab.com/khulnasoft-lab/vul-ci-test/pipelines
[repository]: https://github.com/khulnasoft-lab/vul-ci-test
