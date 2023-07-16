FROM mysql:8.0.23

COPY ./mysql-confd/locale.gen /etc/locale.gen

RUN sed -i 's@archive.ubuntu.com@ftp.jaist.ac.jp/pub/Linux@g' /etc/apt/sources.list

RUN set -ex && \
    apt-get update -qq && \
    apt-get install -y --no-install-recommends locales && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    locale-gen ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8
