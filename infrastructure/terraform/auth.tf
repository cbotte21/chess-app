resource "kubernetes_deployment" "auth" {
  metadata {
    name = "auth"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "auth"
      }
    }

    template {
      metadata {
        labels = {
          app = "auth"
        }
      }
      spec {
        container {
          image             = "chess/auth"
          name              = "auth-container"
          image_pull_policy = "Never"
          port {
            container_port = 5000
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "auth" {
  metadata {
    name = "auth"
  }
  spec {
    selector = {
      app = kubernetes_deployment.auth.spec.0.template.0.metadata.0.labels.app
    }
    session_affinity = "ClientIP"
    port {
      port        = 5000
      target_port = 5000
    }

    type = "LoadBalancer"
  }
  wait_for_load_balancer = false
}
