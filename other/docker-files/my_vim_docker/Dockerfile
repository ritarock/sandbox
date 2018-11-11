FROM alpine:latest

WORKDIR /root

ADD package_manager.py /root
ADD package.json /root
COPY vim.txt /root/.vimrc

RUN set -x && \
    apk add --no-cache vim && \
    apk add --no-cache python3 && \
    apk add --no-cache git && \
    mkdir -p .vim/pack/mypackage/start/ && \
    mkdir -p .vim/pack/mypackage/opt/myplugin/ && \
    python3 package_manager.py && \
    apk del --purge python3 && \
    apk del --purge git
