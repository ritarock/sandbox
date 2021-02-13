#!/bin/sh

rm -rf src
mkdir src

cat <<EOF > src/Gemfile
source 'https://rubygems.org'
gem 'rails', '5.1.6'
EOF
touch src/Gemfile.lock

rm -rf db_volume
mkdir db_volume

docker-compose run web rails new . --force --database=mysql --skip-bundle --skip-coffee

rm -rf src/.git
docker-compose build

sed -i -e "s/password:/password: root/" src/config/database.yml
sed -i -e "s/host: localhost/host: db/" src/config/database.yml

docker-compose run web rails db:create

docker-compose up
