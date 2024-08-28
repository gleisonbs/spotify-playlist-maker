resource "aws_route_table" "route_table_spotify_utils" {
  vpc_id = aws_vpc.spotify_utils.id


  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.spotify_utils_gw.id
  }
}

resource "aws_route_table_association" "subnet_association" {
  subnet_id      = aws_subnet.subnet.id
  route_table_id = aws_route_table.route_table_spotify_utils.id
}