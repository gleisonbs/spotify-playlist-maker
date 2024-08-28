resource "aws_subnet" "subnet" {
  cidr_block        = cidrsubnet(aws_vpc.spotify_utils.cidr_block, 3, 1)
  vpc_id            = aws_vpc.spotify_utils.id
  availability_zone = var.availability_zone
}