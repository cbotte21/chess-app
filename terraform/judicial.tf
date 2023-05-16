resource "kubernetes_deployment" "judicial" {
  metadata {
    name = "judicial"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "judicial"
      }
    }

    template {
      metadata {
        labels = {
          app = "judicial"
        }
      }
      spec {
        container {
          image             = "chess/judicial"
          name              = "judicial-container"
          image_pull_policy = "Never"
          port {
            container_port = 6000
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "judicial" {
  metadata {
    name = "judicial"
  }
  spec {
    selector = {
      app = kubernetes_deployment.judicial.spec.0.template.0.metadata.0.labels.app
    }
    session_affinity = "ClientIP"
    port {
      port        = 6000
      target_port = 6000
    }

    type = "LoadBalancer"
  }
  wait_for_load_balancer = false
}
