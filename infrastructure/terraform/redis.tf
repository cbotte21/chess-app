resource "kubernetes_deployment" "redis" {
  metadata {
    name = "redis"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "redis"
      }
    }

    template {
      metadata {
        labels = {
          app = "redis"
        }
      }
      spec {
        container {
          image = "redis/redis-stack-server:latest"
          name  = "redis-container"
          port {
            container_port = 6379
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "redis" {
  metadata {
    name = "redis"
  }
  spec {
    selector = {
      app = kubernetes_deployment.redis.spec.0.template.0.metadata.0.labels.app
    }
    session_affinity = "ClientIP"
    port {
      port        = 6379
      target_port = 6379
    }
  }
}
