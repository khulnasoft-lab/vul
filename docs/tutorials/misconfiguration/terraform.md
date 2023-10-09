# Scanning Terraform files with Vul

This tutorial is focused on ways Vul can scan Terraform IaC configuration files. 

A video tutorial on Terraform Misconfiguration scans can be found on the [Khulnasoft Open Source YouTube account.](https://youtu.be/BWp5JLXkbBc)

**A note to tfsecurity users** 
We have been consolidating all of our scanning-related efforts in one place, and that is Vul. You can read more on the decision in the [tfsecurity discussions.](https://github.com/khulnasoft-lab/tfsecurity/discussions/1994)

## Vul Config Command 

Terraform configuration scanning is available as part of the `vul config` command. This command scans all configuration files for misconfiguration issues. You can find the details within [misconfiguration scans in the Vul documentation.](https://khulnasoft-lab.github.io/vul/latest/docs/misconfiguration/scanning/) 

Command structure: 
``` 
vul config <any flags you want to use> <file or directory that you would like to scan> 
``` 

The `vul config` command can scan Terraform configuration, CloudFormation, Dockerfile, Kubernetes manifests, and Helm Charts for misconfiguration. Vul will compare the configuration found in the file with a set of best practices.  

- If the configuration is following best practices, the check will pass,  
- If the configuration does not define the resource of some configuration, Vul will assume the default configuration for the resource creation is used. In this case, the check might fail.
- If the configuration that has been defined does not follow best practices, the check will fail.  

### Prerequisites 
Install Vul on your local machines. The documentation provides several [different installation options.](https://khulnasoft-lab.github.io/vul/latest/getting-started/installation/) 
This tutorial will use this example [Terraform tutorial](https://github.com/Cloud-Native-Security/vul-demo/tree/main/bad_iac/terraform) for terraform misconfiguration scanning with Vul. 

Git clone the tutorial and cd into the directory: 
``` 
git clone git@github.com:Cloud-Native-Security/vul-demo.git
cd bad_iac/terraform
``` 
In this case, the folder only containes Terraform configuration files. However, you could scan a directory that contains several different configurations e.g. Kubernetes YAML manifests, Dockerfile, and Terraform. Vul will then detect the different configuration files and apply the right rules automatically. 

## Different types of `vul config` scans 

Below are several examples of how the vul config scan can be used. 

General Terraform scan with vul: 
``` 
vul config <specify the directory> 
``` 

So if we are already in the directory that we want to scan: 
``` 
vul config ./ 
``` 
### Specify the scan format 
The `--format` flag changes the way that Vul displays the scan result: 

JSON: 
```
vul config -f json terraform-infra 
``` 

Sarif: 
``` 
vul config -f sarif terraform-infra 
``` 

### Specifying the output location 

The `--output` flag specifies the file location in which the scan result should be saved: 

JSON: 
``` 
vul config -f json -o example.json terraform-infra 
``` 

Sarif: 
``` 
vul config -f sarif -o example.sarif terraform-infra 
``` 

### Filtering by severity 

If you are presented with lots and lots of misconfiguration across different files, you might want to filter or the misconfiguration with the highest severity: 

``` 
vul config --severity CRITICAL, MEDIUM terraform-infra 
``` 

### Passing tf.tfvars files into `vul config` scans 

You can pass terraform values to Vul to override default values found in the Terraform HCL code. More information are provided [in the documentation.](https://khulnasoft-lab.github.io/vul/latest/docs/misconfiguration/options/values/) 

``` 
vul conf --tf-vars terraform.tfvars ./
``` 
### Custom Checks 

We have lots of examples in the [documentation](https://khulnasoft-lab.github.io/vul/latest/docs/misconfiguration/custom/) on how you can write and pass custom Rego policies into terraform misconfiguration scans. 

## Secret and vulnerability scans

The `vul config` command does not perform secrete and vulnerability checks out of the box. However, you can specify as part of your `vul fs` scan that you would like to scan you terraform files for exposed secrets and misconfiguraction through the following flags: 

```
vul fs --scanners secret,config ./
```

The `vul config` command is a sub-command of the `vul fs` command. You can learn more about this command in the [documentation.](https://khulnasoft-lab.github.io/vul/latest/docs/target/filesystem/) 

## Scanning Terraform Plan files

Instead of scanning your different Terraform resources individually, you could also scan your terraform plan output before it is deployed for misconfiguration. This will give you insights into any misconfiguration of your resources as they would become deployed. [Here](https://khulnasoft-lab.github.io/vul/latest/docs/scanner/misconfiguration/custom/examples/#terraform-plan) is the link to the documentation.

First, create a terraform plan and save it to a file:
```
terraform plan --out tfplan.binary
```

Next, convert the file into json format:
```
terraform show -json tfplan.binary > tfplan.json
```

Lastly, scan the file with the `vul config` command:
```
vul config ./tfplan.json
```

Note that you need to be able to create a terraform init and plan without any errors. 

## Using Vul in your CI/CD pipeline 
Similar to tfsec, Vul can be used either on local developer machines or integrated into your CI/CD pipeline. There are several steps available for different pipelines, including GitHub Actions, Circle CI, GitLab, Travis and more in the tutorials section of the documentation: [https://khulnasoft-lab.github.io/vul/latest/tutorials/integrations/](https://khulnasoft-lab.github.io/vul/latest/tutorials/integrations/) 

 