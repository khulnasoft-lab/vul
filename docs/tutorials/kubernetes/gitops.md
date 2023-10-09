# Installing the Vul-Operator through GitOps

This tutorial shows you how to install the Vul Operator through GitOps platforms, namely ArgoCD and FluxCD.

## ArgoCD

Make sure to have [ArgoCD installed](https://argo-cd.readthedocs.io/en/stable/getting_started/) and running in your Kubernetes cluster.

You can either deploy the Vul Operator through the argocd CLI or by applying a Kubernetes manifest.

ArgoCD command:
```
> kubectl create ns vul-system
> argocd app create vul-operator --repo https://github.com/khulnasoft-lab/vul-operator --path deploy/helm --dest-server https://kubernetes.default.svc --dest-namespace vul-system
```
Note that this installation is directly related to our official Helm Chart. If you want to change any of the value, we'd suggest you to create a separate values.yaml file.

Kubernetes manifest `vul-operator.yaml`:
```
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: vul-operator
  namespace: argocd
spec:
  project: default
  source:
    chart: vul-operator
    repoURL: https://khulnasoft-lab.github.io/helm-charts/
    targetRevision: 0.0.3
    helm:
      values: |
        vul:
          ignoreUnfixed: true
  destination:
    server: https://kubernetes.default.svc
    namespace: vul-system
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
```

To apply the Kubernetes manifest, if you have the manifest locally, you can use the following command through kubectl:
```
> kubectl apply -f vul-operator.yaml

application.argoproj.io/vul-operator created
```

If you have the manifest in a Git repository, you can apply it to your cluster through the following command:
```
> kubectl apply -n argocd -f https://raw.githubusercontent.com/AnaisUrlichs/argocd-starboard/main/starboard/argocd-starboard.yaml
```
The latter command would allow you to make changes to the YAML manifest that ArgoCD would register automatically.

Once deployed, you want to tell ArgoCD to sync the application from the actual state to the desired state:
```
argocd app sync vul-operator
```

Now you can see the deployment in the ArgoCD UI. Have a look at the ArgoCD documentation to know how to access the UI.

![ArgoCD UI after deploying the Vul Operator](../../imgs/argocd-ui.png)

Note that ArgoCD is unable to show the Vul CRDs as synced.


## FluxCD

Make sure to have [FluxCD installed](https://fluxcd.io/docs/installation/#install-the-flux-cli) and running in your Kubernetes cluster.

You can either deploy the Vul Operator through the Flux CLI or by applying a Kubernetes manifest.

Flux command:
```
> kubectl create ns vul-system
> flux create source helm vul-operator --url https://khulnasoft-lab.github.io/helm-charts --namespace vul-system
> flux create helmrelease vul-operator --chart vul-operator
  --source HelmRepository/vul-operator
  --chart-version 0.0.3
  --namespace vul-system
```

Kubernetes manifest `vul-operator.yaml`:
```
apiVersion: source.toolkit.fluxcd.io/v1beta2
kind: HelmRepository
metadata:
  name: vul-operator
  namespace: flux-system
spec:
  interval: 60m
  url: https://khulnasoft-lab.github.io/helm-charts/

---
apiVersion: helm.toolkit.fluxcd.io/v2beta1
kind: HelmRelease
metadata:
  name: vul-operator
  namespace: vul-system
spec:
  chart:
    spec:
      chart: vul-operator
      sourceRef:
        kind: HelmRepository
        name: vul-operator
        namespace: flux-system
      version: 0.10.1
  interval: 60m
  values:
    vul:
      ignoreUnfixed: true
  install:
    crds: CreateReplace
    createNamespace: true
```

You can then apply the file to your Kubernetes cluster:
```
kubectl apply -f vul-operator.yaml
```

## After the installation

After the install, you want to check that the Vul operator is running in the vul-system namespace:
```
kubectl get deployment -n vul-system
```

