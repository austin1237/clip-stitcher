resource "aws_dynamodb_table" "failed-messages" {
  name           = "FailedMessages"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "QueName"
  range_key      = "MessageID"

  attribute {
    name = "QueName"
    type = "S"
  }

  attribute {
    name = "MessageID"
    type = "S"
  }
}