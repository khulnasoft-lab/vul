# CircleCI

```
$ cat .circleci/config.yml
jobs:
  build:
    docker:
      - image: docker:stable-git
    steps:
      - checkout
      - setup_remote_docker
      - run:
          name: Build image
          command: docker build -t vul-ci-test:${CIRCLE_SHA1} .
      - run:
          name: Install vul
          command: |
            apk add --update-cache --upgrade curl
            curl -sfL https://raw.githubusercontent.com/khulnasoft-lab/vul/main/contrib/install.sh | sh -s -- -b /usr/local/bin
      - run:
          name: Scan the local image with vul
          command: vul --exit-code 0 --no-progress vul-ci-test:${CIRCLE_SHA1}
workflows:
  version: 2
  release:
    jobs:
      - build
```

[Example][example]
[Repository][repository]

[example]: https://circleci.com/gh/khulnasoft-lab/vul-ci-test
[repository]: https://github.com/khulnasoft-lab/vul-ci-test
