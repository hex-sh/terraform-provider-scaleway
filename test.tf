variable "scaleway_auth_token" {}
variable "scaleway_organization" {}

provider "scaleway" {
  token = "${var.scaleway_auth_token}"
  organization = "${var.scaleway_organization}"
}

resource "scaleway_server" "hex" {
  name = "Nginx"
  image = "dd17877a-4dc8-4d94-82f5-05019b7f6d67"
  size = "50gb"
  type = "C1"
}




