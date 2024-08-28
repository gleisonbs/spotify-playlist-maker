#!/bin/bash
yum update
yum install git -y
su ec2-user -c 'cd ~; \
    git clone https://github.com/gleisonbs/spotify-playlist-maker.git; \
    chmod +x ~/spotify-playlist-maker/backend/spotify-utils; \
    ~/spotify-playlist-maker/backend/./spotify-utils'