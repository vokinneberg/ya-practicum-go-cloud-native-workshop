// migrations.tf

resource "kubernetes_config_map" "shortener_db_migrations_config" {
  metadata {
    name      = "shortener-db-migrations-config"
    namespace = kubernetes_namespace.ya_paraktikum_go_cloud_native.metadata[0].name
  }

  data = {
    POSTGRES_USER     = "postgres"
    POSTGRES_PASSWORD = var.postgres_password
    POSTGRES_HOST     = kubernetes_service.postgres_service.metadata[0].name
    POSTGRES_PORT     = kubernetes_service.postgres_service.spec[0].port[0].port
    POSTGRES_DB       = "shortener"
  }
}

resource "kubernetes_job" "shortener_db_migrations" {
  metadata {
    name      = "shortener-db-migrations"
    namespace = kubernetes_namespace.ya_paraktikum_go_cloud_native.metadata[0].name
  }

  spec {
    template {
      metadata {
        labels = {
          app = "shortener-db-migrations"
        }
      }

      spec {
        container {
          name  = "migrations-container"
          image = "ya-paraktikum-go-cloud-native-workshop-shortener-db-migrations:${var.app_version}"

          env_from {
            config_map_ref {
              name = kubernetes_config_map.shortener_db_migrations_config.metadata[0].name
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
        }

        restart_policy = "OnFailure"
      }
    }
  }
}