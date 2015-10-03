variable "scaleway_auth_token" {}
variable "scaleway_organization" {}

provider "scaleway" {
  token = "${var.scaleway_auth_token}"
  organization = "${var.scaleway_organization}"
}

resource "scaleway_server" "hex" {
  name = "Nginx"
  image = "16bf98c1-1a50-4212-8a2c-4b2e5001837c"
  size = "50gb"
}

resource "scaleway_ip"  "nginx_ip" {
  server = "${scaleway_server.hex.id}"
}



