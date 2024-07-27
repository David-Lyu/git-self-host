# Todo: remove go_build step and have it prebuilt already
FROM alpine:latest AS go_build
WORKDIR /app/server
COPY . .

# RUN apk add apt
RUN apk add go
RUN apk add build-base
ENV PATH=$PATH:/usr/local/go/bin
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build main.go

FROM alpine:latest AS server

WORKDIR /app

RUN apk add --update apt
RUN apk add bash
# # Need to find the file path of db

COPY ./scripts/init.sh ./init.sh
COPY ./models/database/sqlite-database.db /app/models/database/sqlite-database.db
COPY --from=go_build /app/server/main .

#ARG is build time env (image)
ARG DOCKER_ENV=DOCKER
#ENV is run time env (container)
ENV DOCKER_ENV="$DOCKER_ENV"
ENV MAIN_PATH="$MAIN_PATH"

ENTRYPOINT [ "./init.sh" ]
# EXPOSE 8080
