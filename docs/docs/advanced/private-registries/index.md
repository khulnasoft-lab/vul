Vul can download images from a private registry without the need for installing Docker or any other 3rd party tools.
This makes it easy to run within a CI process.

## Credential
To use Vul with private images, simply install it and provide your credentials:

```shell
$ VUL_USERNAME=YOUR_USERNAME VUL_PASSWORD=YOUR_PASSWORD vul image YOUR_PRIVATE_IMAGE
```

Vul also supports providing credentials through CLI flags:

```shell
$ VUL_PASSWORD=YOUR_PASSWORD vul image --username YOUR_USERNAME YOUR_PRIVATE_IMAGE
```

!!! warning
    The CLI flag `--password` is available, but its use is not recommended for security reasons.

You can also store your credentials in `vul.yaml`.
For more information, please refer to [the documentation](../../references/configuration/config-file.md).

It can handle multiple sets of credentials as well:

```shell
$ export VUL_USERNAME=USERNAME1,USERNAME2
$ export VUL_PASSWORD=PASSWORD1,PASSWORD2
$ vul image YOUR_PRIVATE_IMAGE
```

In the example above, Vul attempts to use two pairs of credentials:

- USERNAME1/PASSWORD1
- USERNAME2/PASSWORD2

Please note that the number of usernames and passwords must be the same.

## docker login
If you have Docker configured locally and have set up the credentials, Vul can access them.

```shell
$ docker login ghcr.io
Username: 
Password:
$ vul image ghcr.io/your/private_image
```

!!! note
    `docker login` can be used with any container runtime, such as Podman.
