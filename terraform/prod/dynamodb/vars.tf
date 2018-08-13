variable "table_name" {
  description = "Name of the table"
}

variable "hash_key" {
  description = "The attribute to use as the hash partition key"
}

variable "range_key" {
  description = "The attribute to use as the range sort key."
}

variable "read_capacity" {
  description = "One read capacity unit represents one strongly consistent read per second, or two eventually consistent reads per second, for an item up to 4 KB in size"
  default = 1
}

variable "write_capacity" {
  description = "One write capacity unit represents one write per second for an item up to 1 KB in size."
  default = 1
}


