
terraform {
  required_version = "1.7.3"
}

variable "terraform_operator_role_name" {
  description = "Terraform operator role name"
  type        = string
  default     = "TerraformOperatorRole"
}

variable "environment" {
  type        = string
  description = "Environment"
  validation {
    condition     = can(regex("^(development|stage|production|sre|cross)$", var.environment))
    error_message = "Must be one of: [development|stage|production|sre|cross]."
  }
}

variable "tenant" {
  description = "Tenant"
  type        = string
  validation {
    condition     = length(var.tenant) >= 2 && length(var.tenant) <= 3 || var.tenant == "cross"
    error_message = "Must be must be 2-3 characters in length or cross literal."
  }
}

variable "label" {
  description = "Label"
  type        = string
  validation {
    condition     = can(regex("^[a-z_0-9]+$", var.label))
    error_message = "Must be must an alpha expression."
  }
}

variable "aws_default_tags" {
  type    = map(string)
  default = {}
}

variable "k8s_default_labels" {
  type    = map(string)
  default = {}
}

variable "project" {
  type = object({
    name            = string
    owner           = string
    kind            = string
    cost_allocation = string
    scope           = string
    pretty_name     = string
  })
}

data "aws_ssm_parameter" "account_name" {
  name = "/providers/aws/organization/cross/accounts/self/account/name"
}

data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

locals {
  environment                          = var.environment
  tenant                               = var.tenant
  label                                = var.label
  project_name                         = var.project.name
  project_owner                        = var.project.owner
  project_kind                         = var.project.kind
  project_cost_allocation              = var.project.cost_allocation
  project_scope                        = var.project.scope
  project_pretty_name                  = var.project.pretty_name
  default_resource_name_prefix         = "${local.tenant}-${local.environment}-${local.project_name}"
  bucket_default_name_prefix           = "${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}-${local.tenant}-${local.environment}"
  ssm_parameter_default_project_prefix = "/${nonsensitive(data.aws_ssm_parameter.account_name.value)}/${data.aws_region.current.name}/${local.tenant}/${local.environment}/${local.project_name}"
  ssm_parameter_default_prefix         = "/${nonsensitive(data.aws_ssm_parameter.account_name.value)}/${data.aws_region.current.name}/${local.tenant}/${local.environment}"
  aws_default_tags = merge({
    Name                       = "${local.tenant}-${local.environment}-${local.project_name}"
    Environment                = local.environment
    CostAllocation             = local.project_cost_allocation
    Tenant                     = local.tenant
    ApplicationScope           = local.project_scope
    "CA:TerraformProjectOwner" = local.project_owner
    "CA:TerraformProjectName"  = local.project_name
    "CA:TerraformProjectKind"  = local.project_kind
    "CA:TerraformProjectScope" = local.project_scope
    "CA:ManagedByTerraform"    = "true"
  }, var.aws_default_tags)
  k8s_default_labels = merge({
    for k, v in local.aws_default_tags :
    format("cloudacademy.com/%s", lower(replace(trimprefix(k, "CA:"), "/(.)([A-Z])/", "$${1}-$2"))) => v
  }, var.k8s_default_labels)
  aws_default_tags_string_map = trimsuffix(<<EOT
%{for key, value in local.aws_default_tags}${key}=${value},%{endfor}
EOT
  , ",\n")
}
