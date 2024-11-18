resource "kubernetes_secret" "postgres_secret" {
  metadata {
    name      = "postgres-password"
    namespace = kubernetes_namespace.ya_paraktikum_go_cloud_native.metadata[0].name
  }
  data = {
    pg_password = var.postgres_password
  }
  type = "Opaque"
}
