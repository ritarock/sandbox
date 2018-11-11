#!/bin/sh
sudo apt-get update
sudo apt-get -y install git
sudo git clone https://github.com/wfarr/goenv.git ~/.goenv
echo 'export PATH="$HOME/.goenv/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(goenv init -)"' >> ~/.bashrc
source ~/.bashrc
goenv install 1.9.2
goenv global 1.9.2
