provider "aws" {
  assume_role {
    role_arn = "arn:aws:iam::${var.target_account_id}:role/${var.terraform_operator_role_name}"
  }
  default_tags {
    tags = local.aws_default_tags
  }
}
