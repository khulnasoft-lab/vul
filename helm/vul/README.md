# Vul Scanner

Vul vulnerability scanner standalone installation.

## TL;DR;

```
$ helm install vul . --namespace vul --create-namespace
```

## Introduction

This chart bootstraps a Vul deployment on a [Kubernetes](http://kubernetes.io) cluster using the
[Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.12+
- Helm 3+

## Installing from the Khulnasoft Chart Repository

```
helm repo add khulnasoft-lab https://khulnasoft-lab.github.io/helm-charts/
helm repo update
helm search repo vul
helm install my-vul khulnasoft-lab/vul
```

## Installing the Chart

To install the chart with the release name `my-release`:

```
$ helm install my-release .
```

The command deploys Vul on the Kubernetes cluster in the default configuration. The [Parameters](#parameters)
section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`.

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Parameters

The following table lists the configurable parameters of the Vul chart and their default values.

|                 Parameter             |                                Description                              |    Default     |
|---------------------------------------|-------------------------------------------------------------------------|----------------|
| `image.registry`                      | Image registry                                                          | `docker.io`    |
| `image.repository`                    | Image name                                                              | `khulnasoft/vul` |
| `image.tag`                           | Image tag                                                               | `{TAG_NAME}`   |
| `image.pullPolicy`                    | Image pull policy                                                       | `IfNotPresent` |
| `image.pullSecret`                    | The name of an imagePullSecret used to pull vul image from e.g. Docker Hub or a private registry  | |
| `replicaCount`                        | Number of Vul Pods to run                                   | `1`            |
| `vul.debugMode`                     | The flag to enable or disable Vul debug mode                          | `false` |
| `vul.gitHubToken`                   | The GitHub access token to download Vul DB. More info: https://github.com/khulnasoft-lab/vul#github-rate-limiting                          |      |
| `vul.registryUsername`              | The username used to log in at dockerhub. More info: https://khulnasoft-lab.github.io/vul/dev/advanced/private-registries/docker-hub/ |      |
| `vul.registryPassword`              | The password used to log in at dockerhub. More info: https://khulnasoft-lab.github.io/vul/dev/advanced/private-registries/docker-hub/ |      |
| `vul.registryCredentialsExistingSecret` | Name of Secret containing dockerhub credentials. Alternative to the 2 parameters above, has precedence if set.                    |      |
| `vul.serviceAccount.annotations`        | Additional annotations to add to the Kubernetes service account resource |     |
| `vul.skipDBUpdate`                    | The flag to enable or disable Vul DB downloads from GitHub            | `false`        |
| `vul.dbRepository`                  | OCI repository to retrieve the vul vulnerability database from        | `ghcr.io/khulnasoft-lab/vul-db`        |
| `vul.cache.redis.enabled`           | Enable Redis as caching backend                                         | `false` |
| `vul.cache.redis.url`               | Specify redis connection url, e.g. redis://redis.redis.svc:6379         | `` |
| `vul.cache.redis.ttl`               | Specify redis TTL, e.g. 3600s or 24h                                    | `` |
| `vul.cache.redis.tls`               | Enable Redis TLS with public certificates                               | `` |
| `vul.serverToken`                   | The token to authenticate Vul client with Vul server                | `` |
| `vul.existingSecret`                | existingSecret if an existing secret has been created outside the chart. Overrides gitHubToken, registryUsername, registryPassword, serverToken | `` |
| `vul.podAnnotations`                | Annotations for pods created by statefulset                             | `{}` |
| `vul.extraEnvVars`                  | extraEnvVars to be set on the container                                 | `{}` |
| `service.name`                        | If specified, the name used for the Vul service                       |     |
| `service.type`                        | Kubernetes service type                                                 | `ClusterIP` |
| `service.port`                        | Kubernetes service port                                                 | `4954`      |
| `service.sessionAffinity`             | Kubernetes service session affinity                                     | `ClientIP`  |
| `httpProxy`                           | The URL of the HTTP proxy server                                        |     |
| `httpsProxy`                          | The URL of the HTTPS proxy server                                       |     |
| `noProxy`                             | The URLs that the proxy settings do not apply to                        |     |
| `nodeSelector`                        | Node labels for pod assignment                                              |     |
| `affinity`                            | Affinity settings for pod assignment                                              |     |
| `tolerations`                         | Tolerations for pod assignment                                              |     |
| `podAnnotations`                      | Annotations for pods created by statefulset                             | `{}` |

The above parameters map to the env variables defined in [vul](https://github.com/khulnasoft-lab/vul#configuration).

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

```
$ helm install my-release . \
       --namespace my-namespace \
       --set "service.port=9090" \
       --set "vul.vulnType=os\,library"
```

## Storage

This chart uses a PersistentVolumeClaim to reduce the number of database downloads between POD restarts or updates. The storageclass should have the reclaim policy  `Retain`.

## Caching

You can specify a Redis server as cache backend. This Redis server has to be already present. You can use the [bitnami chart](https://bitnami.com/stack/redis/helm).
More Information about the caching backends can be found [here](https://github.com/khulnasoft-lab/vul#specify-cache-backend).
