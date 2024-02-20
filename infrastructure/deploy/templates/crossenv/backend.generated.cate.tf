
terraform {
  backend "s3" {
    bucket         = "713209205358-us-west-2-cross-cross-tfstate"
    key            = "<key will be generated at deployment time>"
    region         = "us-west-2"
    dynamodb_table = "tools-us-west-2-cross-cross-tflock"
    encrypt        = true
    assume_role = {
      role_arn = "arn:aws:iam::713209205358:role/TerraformStateRole"
    }
  }
}
