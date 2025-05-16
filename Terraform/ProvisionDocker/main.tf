provider "docker" {}

resource "docker_image" "nginx" {
	name = "nginx:latest"
	keep_locally = false
}

resource "docker_container" "nginx" {
	image = docker_image.nginx.image_id
	name  = "terraform-nginx"
	ports {
		internal = 80
		external = 9100
	}
}