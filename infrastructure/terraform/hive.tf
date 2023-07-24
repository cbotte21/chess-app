resource "kubernetes_deployment" "hive" {
  metadata {
    name = "hive"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "hive"
      }
    }

    template {
      metadata {
        labels = {
          app = "hive"
        }
      }
      spec {
        container {
          image             = "chess/hive"
          name              = "hive-container"
          image_pull_policy = "Never"
          port {
            container_port = 6001
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "hive" {
  metadata {
    name = "hive"
  }
  spec {
    selector = {
      app = kubernetes_deployment.hive.spec.0.template.0.metadata.0.labels.app
    }
    session_affinity = "ClientIP"
    port {
      port        = 6001
      target_port = 6001
    }

    type = "LoadBalancer"
  }
  wait_for_load_balancer = false
}
