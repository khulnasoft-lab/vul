# Travis CI

```
$ cat .travis.yml
services:
  - docker

env:
  global:
    - COMMIT=${TRAVIS_COMMIT::8}

before_install:
  - docker build -t vul-ci-test:${COMMIT} .
  - export VERSION=$(curl --silent "https://api.github.com/repos/khulnasoft-lab/vul/releases/latest" | grep '"tag_name":' | sed -E 's/.*"v([^"]+)".*/\1/')
  - wget https://github.com/khulnasoft-lab/vul/releases/download/v${VERSION}/vul_${VERSION}_Linux-64bit.tar.gz
  - tar zxvf vul_${VERSION}_Linux-64bit.tar.gz
script:
  - ./vul image --exit-code 0 --severity HIGH --no-progress vul-ci-test:${COMMIT}
  - ./vul image --exit-code 1 --severity CRITICAL --no-progress vul-ci-test:${COMMIT}
cache:
  directories:
    - $HOME/.cache/vul
```

[Example][example]
[Repository][repository]

[example]: https://travis-ci.org/khulnasoft-lab/vul-ci-test
[repository]: https://github.com/khulnasoft-lab/vul-ci-test
