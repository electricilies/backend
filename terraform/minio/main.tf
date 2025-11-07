terraform {
  required_version = ">= 1.13.4"
  required_providers {
    minio = {
      source  = "aminueza/minio"
      version = ">= 3.11.3"
    }
  }
}

resource "minio_config" "notify_webhook_backend" {
  key   = "notify_webhook:backend"
  value = "endpoint=${var.backend_webhook} auth_token= queue_limit=0 queue_dir= client_cert= client_key="
}

resource "minio_s3_bucket" "electricilies" {
  acl    = "public"
  bucket = "electricilies"
}

resource "minio_s3_bucket_notification" "backend" {
  bucket = minio_s3_bucket.electricilies.bucket

  queue {
    id        = "notification-webook-backend"
    queue_arn = "arn:minio:sqs::backend:webhook"
    events = [
      "s3:ObjectCreated:*",
      "s3:ObjectRemoved:*",
    ]
  }
}
