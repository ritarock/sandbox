FROM ruby:2.6.5

RUN apt-get update -qq && \
    apt-get install -y build-essential && \
    apt-get install -y libpq-dev && \
    apt-get install -y default-mysql-client && \
    apt-get install -y nodejs

RUN mkdir /myapp
ENV APP_ROOT /myapp
WORKDIR $APP_ROOT
ADD ./src/Gemfile $APP_ROOT/Gemfile
ADD ./src/Gemfile.lock $APP_ROOT/Gemfile.lock

RUN bundle install
ADD ./src/ $APP_ROOT
