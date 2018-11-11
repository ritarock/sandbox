#!/bin/sh
sudo apt update
sudo apt install software-properties-common
yes '' | sudo apt-add-repository ppa:ansible/ansible
sudo apt update
sudo apt -y install ansible
