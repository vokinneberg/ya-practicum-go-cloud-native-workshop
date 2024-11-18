# main.tf
resource "kubernetes_namespace" "ya_paraktikum_go_cloud_native" {
  metadata {
    name = "ya-paraktikum-go-cloud-native"
  }
}

resource "kubernetes_config_map" "shortener_config" {
  metadata {
    name      = "shortener-config"
    namespace = kubernetes_namespace.ya_paraktikum_go_cloud_native.metadata[0].name
  }

  data = {
    APP_NAME        = "shortener"
    BASE_URL        = "http://shortener.minikube"
    DATABASE_DSN    = "postgres://postgres:${var.postgres_password}@${kubernetes_service.postgres_service.metadata[0].name}:${kubernetes_service.postgres_service.spec[0].port[0].port}/shortener?sslmode=disable"
    ENVIRONMENT     = "development"
    LOG_LEVEL       = "info"
    RUN_MIGRATIONS  = "false"
    SERVER_ADDRESS  = ":80"
  }
}

resource "kubernetes_deployment" "shortener" {
  depends_on = [kubernetes_service.postgres_service]
  metadata {
    name      = "shortener"
    namespace = kubernetes_namespace.ya_paraktikum_go_cloud_native.metadata[0].name
    labels = {
      app = "shortener"
    }
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "shortener"
      }
    }

    template {
      metadata {
        labels = {
          app = "shortener"
        }
      }

      spec {
        container {
          name  = "shortener-container"
          image = "ya-paraktikum-go-cloud-native-workshop-shortener:${var.app_version}"
          port {
            container_port = 80
          }

          env_from {
            config_map_ref {
              name = kubernetes_config_map.shortener_config.metadata[0].name
            }
          }

          resources {
            requests = {
              cpu    = "100m"
              memory = "128Mi"
            }
            limits = {
              cpu    = "500m"
              memory = "512Mi"
            }
          }

          # Liveness probe
          liveness_probe {
            http_get {
              path = "/healthz"
              port = 80

              http_header {
                name = "User-Agent"
                value = "kube-liveness-probe"
              }
            }
            initial_delay_seconds = 10
            period_seconds = 5
          }

          # Readiness probe
          readiness_probe {
            http_get {
              path = "/readyz"
              port = 80

              http_header {
                name = "User-Agent"
                value = "kube-readiness-probe"
              }
            }
            initial_delay_seconds = 5
            period_seconds = 5
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "shortener_service" {
  metadata {
    name      = "shortener-service"
    namespace = kubernetes_namespace.ya_paraktikum_go_cloud_native.metadata[0].name
  }
  spec {
    selector = {
      app = "shortener"
    }
    port {
      port        = 80
      target_port = 80
    }
    type = "ClusterIP"
  }
}

resource "kubernetes_ingress_v1" "shortener_ingress" {
  metadata {
    name      = "shortener-ingress"
    namespace = kubernetes_namespace.ya_paraktikum_go_cloud_native.metadata[0].name
  }
  spec {
    ingress_class_name = "nginx"
    rule {
      host = "shortener.minikube" # Custom hostname for accessing the service
      http {
        path {
          path = "/*"
          backend {
            service {
              name = kubernetes_service.shortener_service.metadata[0].name
              port {
                number = 80
              }
            }
          }
        }
      }
    }
  }
}
