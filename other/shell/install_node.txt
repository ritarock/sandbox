#!/bin/sh
sudo apt-get update
sudo apt-get install -y build-essential libssl-dev
sudo apt-get install -y curl
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.1/install.sh | bash
source ~/.bashrc
nvm ls-remote
nvm install v8.9.4
node -v
