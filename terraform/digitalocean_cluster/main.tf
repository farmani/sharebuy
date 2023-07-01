terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "2.28.1"
    }
  }
}

provider "digitalocean" {
  token = var.digitalocean_token
}

variable "digitalocean_token" {
  description = "Digitalocean API token"
}

variable "label" {
  description = "Cluster unique label"
  default = "waptap-test"
}

variable "region" {
  description = "Cluster region"
  default = "nyc3"
}

variable "tags" {
  description = "List of tags to apply to the cluster"
  type = list(string)
  default = ["testing"]
}

variable "k8s_version" {
  description = "Kubernetes version to use for the cluster"
  default = "1.26.5-do.0"
}

variable "pools" {
  description = "The Node Pools for the cluster"
  type = list(object({
    name = string
    size = string
    node_count = number
    labels = map(string)
  }))
  default = [{
    name = "workers"
    size  = "s-6vcpu-16gb"
    node_count = 3
    labels = {
      # TODO: Those labels are required for rook-ceph
      # My suggestion for production is to use another nodegroup
      # just for rook-ceph, however it is not strictly required
      "app" = "rook-ceph-osd",
      "app" = "rook-ceph-osd-prepare"
    }
  }]
}


resource "digitalocean_kubernetes_cluster" "default" {
  name       = var.label
  version = var.k8s_version
  region      = var.region
  tags        = var.tags

  dynamic node_pool {
    for_each = var.pools
    content {
      name = node_pool.value["name"]
      size = node_pool.value["size"]
      node_count = node_pool.value["node_count"]
      labels = node_pool.value["labels"]
    }
  }
}

resource "digitalocean_firewall" "my_firewall" {
  depends_on = [digitalocean_kubernetes_cluster.default]

  name = "waptap-k8s"

  inbound_rule {
    protocol = "tcp"
    port_range = "80"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }
  inbound_rule {
    protocol = "tcp"
    port_range = "443"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }
}


output "kubeconfig" {
  value = digitalocean_kubernetes_cluster.default.kube_config
  sensitive = true
}

output "status" {
  value = digitalocean_kubernetes_cluster.default.status
}

output "id" {
  value = digitalocean_kubernetes_cluster.default.id
}

output "pool" {
  value = digitalocean_kubernetes_cluster.default.node_pool
}
