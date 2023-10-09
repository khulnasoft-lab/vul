# Kubernetes Scanning Tutorial

## Prerequisites 

To test the following commands yourself, make sure that you’re connected to a Kubernetes cluster. A simple kind, a Docker-Desktop or microk8s cluster will do. In our case, we’ll use a one-node kind cluster.  
 
Pro tip: The output of the commands will be even more interesting if you have some workloads running in your cluster. 

## Cluster Scanning

Vul K8s is great to get an overview of all the vulnerabilities and misconfiguration issues or to scan specific workloads that are running in your cluster. You would want to use the Vul K8s command either on your own local cluster or in your CI/CD pipeline post deployments.  

The `vul k8s` command is part of the Vul CLI. 

With the following command, we can scan our entire Kubernetes cluster for vulnerabilities and get a summary of the scan: 

```
vul k8s --report=summary cluster
```

To get detailed information for all your resources, just replace ‘summary’ with ‘all’: 

```
vul k8s --report=all cluster
```

However, we recommend displaying all information only in case you scan a specific namespace or resource since you can get overwhelmed with additional details. 

Furthermore, we can specify the namespace that Vul is supposed to scan to focus on specific resources in the scan result: 

```
vul k8s -n kube-system --report=summary cluster
```

Again, if you’d like to receive additional details, use the ‘--report=all’ flag: 

```
vul k8s -n kube-system --report=all cluster
```

Like with scanning for vulnerabilities, we can also filter in-cluster security issues by severity of the vulnerabilities: 

```
vul k8s --severity=CRITICAL --report=summary cluster
```

Note that you can use any of the Vul flags on the Vul K8s command. 

With the Vul K8s command, you can also scan specific workloads that are running within your cluster, such as our deployment: 

```
vul k8s --namespace  app --report=summary deployments/react-application
```

## Vul Operator 

The Vul K8s command is an imperative model to scan resources. We wouldn’t want to manually scan each resource across different environments. The larger the cluster and the more workloads are running in it, the more error-prone this process would become. With the Vul Operator, we can automate the scanning process after the deployment.  

The Vul Operator follows the Kubernetes Operator Model. Operators automate human actions, and the result of the task is saved as custom resource definitions (CRDs) within your cluster. 

This has several benefits: 

- Vul Operator is installed CRDs in our cluster. As a result, all our resources, including our security scanner and its scan results, are Kubernetes resources. This makes it much easier to integrate the Vul Operator directly into our existing processes, such as connecting Vul with Prometheus, a monitoring system. 

- The Vul Operator will automatically scan your resources every six hours. You can set up automatic alerting in case new critical security issues are discovered. 

- The CRDs can be both machine and human-readable depending on which applications consume the CRDs. This allows for more versatile applications of the Vul operator. 

 
There are several ways that you can install the Vul Operator in your cluster. In this guide, we’re going to use the Helm installation based on the [following documentation.](../../docs/target/kubernetes.md#vul-operator)

Please follow the Vul Operator documentation for further information on:

- [Installation of the Vul Operator](https://khulnasoft-lab.github.io/vul-operator/latest/getting-started/installation/)
- [Getting started guide](https://khulnasoft-lab.github.io/vul-operator/latest/getting-started/quick-start/)



 

