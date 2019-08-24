FROM ruby

RUN set -x && \
  apt-get update && \
  apt-get install -y nodejs && \
  gem install rails -v 5.1.6

RUN mkdir /myapp
ENV APP_ROOT /myapp
WORKDIR $APP_ROOT

ADD ./src/Gemfile $APP_ROOT/Gemfile
ADD ./src/Gemfile.lock $APP_ROOT/myapp/Gemfile.lock

RUN bundle install
COPY ./src $APP_ROOT

