resource "aws_ecr_repository" "this" {
  #checkov:skip=CKV_AWS_51
  name                 = "cloudacademy/${local.project_name}"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
  encryption_configuration {
    encryption_type = "KMS"
  }
}

resource "aws_ecr_repository_policy" "cross_account_policy" {
  repository = aws_ecr_repository.this.name
  policy = data.aws_iam_policy_document.cross_account_policy.json
}
