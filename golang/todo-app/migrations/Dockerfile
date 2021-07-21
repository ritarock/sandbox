# use https://github.com/ufoscout/docker-compose-wait & https://github.com/golang-migrate/migrate
## Use whatever base image
FROM migrate/migrate

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

## Add your application to the docker image
ADD ./migrations /migrations

## Launch the wait tool and then your application
ENTRYPOINT /wait && migrate -path /migrations -database "mysql://user:password@tcp(db:3306)/app" up
