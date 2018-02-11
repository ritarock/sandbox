FROM alpine

RUN apk add --no-cache nodejs

WORKDIR /app

ADD . /app

RUN npm install

EXPOSE 8080

CMD ["npm", "run", "dev"]
