terraform {
  required_providers {
    yunjigjl = {
      source = "yunji/yunjigjl"
    }
  }
}

resource "yunjigjl_demo" "test" {
  instance_name= "guanguan"
  disk_size = 100
  tags = "test"
}