#Internet Gateway
resource "aws_internet_gateway" "spotify_utils_gw" {
  vpc_id = aws_vpc.spotify_utils.id
}