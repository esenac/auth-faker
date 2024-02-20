data "aws_ssm_parameter" "platform_development_account_id" {
  name = "/providers/aws/organization/main/accounts/platform-development/account/id"
}

data "aws_iam_policy_document" "cross_account_policy" {
  statement {
    sid    = "CrossAccountPolicy"
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = [
        nonsensitive(data.aws_ssm_parameter.platform_development_account_id.value),
      ]
    }

    actions = [
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability"
    ]
  }
}
