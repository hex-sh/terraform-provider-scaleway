variable "scaleway_auth_token" {}
variable "scaleway_organization" {}

provider "scaleway" {
  token = "${var.scaleway_auth_token}"
  organization = "${var.scaleway_organization}"
}

resource "scaleway_ip" "nginx_ip" {}

resource "scaleway_volume" "storage" {
  size = "50GB"
}

resource "scaleway_server" "hex" {
  name = "Nginx"
  image = "87256bb0-f531-4367-bc93-be04581dfb67"
  ip = "${scaleway_ip.nginx_ip.id}"
  volumes = {
    "1" = "${scaleway_volume.storage.id}"
  }
}



