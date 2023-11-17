variable "name" {
  type    = string
  default = "unspecified"
}

variable "set_environment" {
  type = string
  default = ""
}

variable "version" {
  type = string
  default = "unknown"
}

variable "port" {
  type = string
  default = "5000"
}

packer {
  required_plugins {
    docker = {
      version = ">= 0.0.7"
      source  = "github.com/hashicorp/docker"
    }
  }
}

source "docker" "ubuntu" {
  image  = "golang:1.21.0-alpine3.18"
  commit = true
  changes = [
    "WORKDIR /go/app/cmd",
    "EXPOSE ${var.port}",
    "ENTRYPOINT /go/app/cmd/cmd"
  ]
}

build {
  name = "chess/${var.name}"
  sources = [
    "source.docker.ubuntu"
  ]
  provisioner "shell" {
    inline = [
      "apk add git",
      "git clone https://github.com/cbotte21/${var.name} app",
      "cd app/cmd",
      "go get",
      "go build",
    ]
  }
  provisioner "shell" {
    inline = [
      "cd /go/app/cmd",
      "printf '%s\n' ${var.set_environment} >> .env"
    ]
  }
  post-processors {
    post-processor "docker-tag" {
      repository = "chess/${var.name}"
      tags       = ["${var.version}", "latest"]
    }
  }
}
