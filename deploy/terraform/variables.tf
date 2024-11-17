variable "postgres_password" {
  description = "The password for the postgres user"
  type      = string
  sensitive = true
}

variable "app_version" {
  description = "The version of the shortener image to deploy"
  type        = string
}
