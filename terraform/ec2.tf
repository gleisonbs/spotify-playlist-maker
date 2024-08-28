data "aws_ami" "ubuntu_ami" {
  most_recent = true

  filter {
    name   = "name"
    values = ["*ubuntu-noble-24.04-amd64-server*"]
  }
}

resource "aws_instance" "spotify_playlist_maker" {
  # ami                         = "ami-0e86e20dae9224db8"
  ami                         = "ami-066784287e358dad1"
  instance_type               = "t2.micro"
  associate_public_ip_address = true
  security_groups             = ["${aws_security_group.security.id}"]
  subnet_id                   = aws_subnet.subnet.id
  key_name                    = var.key_pair_name

  tags = {
    Name = "spotify-playlist-maker"
  }

  user_data = file("./setup.sh")
}

output "instance_ip" {
  description = "The public IP address for SSH access"
  value       = aws_instance.spotify_playlist_maker.public_ip
}
