# Introduction

This project handles the creation of the resources required to run the Waptap application on a Kubernetes cluster.
Optionally, a Digital Ocean Kubernetes Service or a Linode Kubernetes Engine cluster can be created as well. Locally
there is no high quality support for any Kubernetes Engine setup using Terraform. For local Kubernetes, you should setup
the Kubernetes Engine by yourself as detailed below.

# Requirements

- Kubectl >= 1.25 and <= 1.27. Follow instructions at: https://kubernetes.io/docs/tasks/tools/ . Remember to not install
  1.28 or higher.
- Terraform >= 0.13: https://developer.hashicorp.com/terraform/downloads
- DNS configuration

# DigitalOcean Kubernetes (DOKS)

## Additional requirements

Those are the additional requirements while using LKE.

- `doctl`: Follow instructions at: https://docs.digitalocean.com/reference/doctl/how-to/install/
- after installation authorize and switch to context:

```shell
$ doctl auth init --context waptap-test
Please authenticate doctl for use with your DigitalOcean account. You can generate a token in the control panel at https://cloud.digitalocean.com/account/api/tokens

❯ Enter your access token:  ●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●●

Validating token... ✔
$ doctl auth switch --context waptap-test
$ doctl auth list
default
waptap-test (current)
```

## Steps

1. A DigitalOcean Token is required. It is passed to Terraform through the environment
   variable `TF_VAR_digitalocean_token`. A suggestion is to save the token as the contents of the
   file `$HOME/.config/doctl/aparlay.token` and export
   with `export TF_VAR_digitalocean_token=$(cat $HOME/.config/doctl/aparlay.token)`.
1. A CloudFlare API Token is required. It is passed to Terraform through the environment
   variable `TF_VAR_cloudflare_api_token`. A suggestion is to save the token as the contents of the
   file `$HOME/.config/cloudflare.token` and export
   with `export TF_VAR_cloudflare_api_token=$(cat $HOME/.config/cloudflare.token)`.
1. `make digitalocean`: this will create the DOKS cluster

## Dashboard

A DigitalOcean Kubernetes (DOKS) dashboard is available at: [[https://cloud.digitalocean.com/kubernetes/clusters]].
Details can be seen by clicking the cluster name. Within this dashboard, there is an external link for the standard
Kubernetes Dashboard, which is focused on the cluster Kubernetes resources. Hardware resources are managed on the DOKS
dashboard.

## Cluster access

1. Running `make setup_doks_access` will save the Kubernetes configuration to your local machine
1. Optionally you can run `make manually_setup_doks_access` if you don't want to override your default configuration

## Example: loading a database file

The easiest for loading a database file is to copy the file into a MongoDB instance and load it from there. If you
followed the previous step on cluster access, there is no need to specify the namespace (`-n waptap`).

The following command copies a file into some of the MongoDB instances:
`kubectl cp -c mongod ./tmp/db.gz mongodb-waptap-rs0-1:/tmp/`
The file is copied into the `mongod` container of the rs0-1 replicaset.

The following commands loads this file into MongoDB:
`kubectl  exec mongodb-waptap-rs0-1 -it -- mongorestore $(kubectl get secret mongodb-waptap-secrets --template="{{.data.MONGODB_DB_DSN_DATABASE_ADMIN}}" | base64 -d) --archive=/tmp/db.gz --gzip --drop`
The command above executes the `mongorestore` command on the same replicaset we copied the file to. The inner shell
command is fetching the full DB_DSN for the database admin user, which is stored as a Kubernetes secret.
You can run the following shell and copy the DSN if you prefer:
`kubectl get secret mongodb-waptap-secrets --template="{{.data.MONGODB_DB_DSN_DATABASE_ADMIN}}" | base64 -d`

After the double dashes (`--`) you can use any MongoDB command. For convenience you can also use `bash` and run the
commands from within the container. I usually prefer the former approach because of the shell history.

# Linode Kubernetes Engine (LKE)

## Additional requirements

Those are the additional requirements while using LKE.

- `jq`. Follow instructions at: https://stedolan.github.io/jq/download/
  . `jq` is used to assist us into configuring access to the newly created cluster.
- `linode-cli`: Follow instructions at: https://www.linode.com/docs/products/tools/cli/guides/install/

## Steps

1. A Linode Token is required. It is passed to Terraform through the environment variable `TF_VAR_linode_token`. A
   suggestion is to save the token as the contents of the file `$HOME/.linode/aparlay.token` and export
   with `export TF_VAR_linode_token=$(cat $HOME/.linode/aparlay.token)`.
2. `make linode`: this will create the LKE cluster

**Note:** Linode Kubernetes Engine sometimes returns EOF on Kubernetes API calls. If that's the case, re-run the command
above.

## Optional steps

1. At the end of the previous command, there will be an optional steps written to the standard output. It is required to
   run it in a terminal where you want to use kubectl for accessing the cluster. It is as
   follows: `export KUBECONFIG=$(pwd)/kubeconfig.linode`, assuming you are at the top level folder for this project.
2. Running `make setup_lke_access` will output the command again, for assisting in your local configuration setup.
3. You can also save the file contents at `$HOME/kube/config` if it does not exist, or edit this file do add the
   contents otherwise.

# Local Kubernetes

The waptap setup can be performed locally or on Linode Kubernetes Engine. If setting up locally, a Kubernetes must be
installed locally. The only fully tested solution at this point is MicroK8s.
MicroK8s: https://microk8s.io/

Another solid option is https://minikube.sigs.k8s.io/docs/start/

Note: I have tested using Kind and considered using Flexkube given some level of support for Kubernetes Terraform
Providers, however they do not fit the minimum requirements.

## MicroK8s

The required modules are required:

sudo microk8s enable dns metallb rbac storage ingress

- dns
- ingress
- metallb
- rbac
- storage
- registry

MetalLB requires an IP range, I picked, however you might need to chose something else 172.25.0.1-172.25.0.255.

For Docker Registry usage, some Docker setups will refuse to use the insecure registry, if pushing to it fails after
enabling the server, you must edit your /etc/docker/daemon.json to add:

```
{
  "insecure-registries" : ["localhost:32000"]
}
```

then restart Docker with:
`sudo systemctl restart docker`

## DNS Config [Optional]

This section is needed only for debugging purposes or for some very specific development needs. For most use cases, you
can skip this section.

resolved.conf stub example:

```
[Resolve]
Cache=yes
DNS=10.152.183.10
Domains=waptap.svc.cluster.local svc.cluster.local cluster.local
```

Please confirm the DNS address first with: ``

If you are already overriding the default DNS providers at the system level, you might want to combine the entries, as
in:

```
[Resolve]
Cache=yes
DNS=10.152.183.10 1.1.1.1 1.0.0.1
Domains=waptap.svc.cluster.local svc.cluster.local cluster.local
```

Place file into `/etc/systemd/resolved.conf.d/`, creating the folder as needed. Check config with:

```
systemd-analyze cat-config systemd/resolved.conf
```

the file contents should appear at the bottom.

## How to point the waptap domains the cluster

### Using /etc/hosts

As usual, you can edit your `/etc/hosts` file adding a line pointing to the IP address exposed to your machine.

If using MetalLB + HAProxy, with the settings above, your IP address is likely going to be 172.25.0.1. You can get the
IP address with:
`kubectl -n ingress-controller get service haproxy-ingress`

An example, would be:
`172.25.0.1 waptap.test app.waptap.test admin.waptap.test s.waptap.test a.waptap.test api.waptap.test`

## Get logs of pods

as some containers stored in app pod and logs are in stderr you need to run this command to get all containers pods

`kubctl logs -f POD_NAME --all-containers=true --max-log-requests=10`
