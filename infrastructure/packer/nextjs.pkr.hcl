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
  image  = "node:20-alpine"
  commit = true
  changes = [
    "WORKDIR /app",
    "EXPOSE ${var.port}",
    "ENTRYPOINT [\"npm\", \"run\", \"dev\"]"
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
      "git submodule update --init --recursive",
      "cd app/",
      "printf '%s\n' ${var.set_environment} >> .env",
      "yarn install",
      "npm run build"
    ]
  }
  post-processors {
    post-processor "docker-tag" {
      repository = "chess/${var.name}"
      tags       = ["${var.version}", "latest"]
    }
  }
}
