# Built-in Policies

## Policy Sources
Built-in policies are mainly written in [Rego][rego] and Go.
Those policies are managed under [defsec repository][defsec].
See [here](../../../coverage/iac/index.md) for the list of supported config types.

For suggestions or issues regarding policy content, please open an issue under the [defsec][defsec] repository.

## Policy Distribution
defsec policies are distributed as an OPA bundle on [GitHub Container Registry][ghcr] (GHCR).
When misconfiguration detection is enabled, Vul pulls the OPA bundle from GHCR as an OCI artifact and stores it in the cache.
Those policies are then loaded into Vul OPA engine and used for detecting misconfigurations.
If Vul is unable to pull down newer policies, it will use the embedded set of policies as a fallback. This is also the case in air-gap environments where `--skip-policy-update` might be passed.

## Update Interval
Vul checks for updates to OPA bundle on GHCR every 24 hours and pulls it if there are any updates.

[rego]: https://www.openpolicyagent.org/docs/latest/policy-language/

[kubernetes-policies]: https://github.com/khulnasoft/defsec/tree/master/rules/kubernetes/policies
[docker-policies]: https://github.com/khulnasoft/defsec/tree/master/rules/docker/policies
[defsec]: https://github.com/khulnasoft/defsec
[ghcr]: https://github.com/khulnasoft/defsec/pkgs/container/defsec