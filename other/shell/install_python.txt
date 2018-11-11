#!/bin/sh
sudo apt-get update
sudo apt-get install -y git
sudo apt-get install -y build-essential libncursesw5-dev libgdbm-dev libc6-dev zlib1g-dev libsqlite3-dev tk-dev libssl-dev openssl libbz2-dev libreadline-dev
git clone https://github.com/yyuu/pyenv.git ~/.pyenv
echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.profile
echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.profile
echo 'eval "$(pyenv init -)"' >> ~/.profile
source ~/.profile
pyenv install 3.6.3
pyenv global 3.6.3
