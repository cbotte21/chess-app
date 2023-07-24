resource "kubernetes_deployment" "nextjs" {
  metadata {
    name = "nextjs-client"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "client"
      }
    }

    template {
      metadata {
        labels = {
          app = "client"
        }
      }
      spec {
        container {
          image             = "chess/nextjs"
          name              = "nextjs-container"
          image_pull_policy = "Never"
          port {
            container_port = 3000
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "nextjs" {
  metadata {
    name = "nextjs-client"
  }
  spec {
    selector = {
      app = kubernetes_deployment.nextjs.spec.0.template.0.metadata.0.labels.app
    }
    session_affinity = "ClientIP"
    port {
      node_port   = 30000
      port        = 3000
      target_port = 3000
    }

    type = "LoadBalancer"
  }
  wait_for_load_balancer = false
}
