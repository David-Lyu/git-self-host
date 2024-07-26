FROM alpine:latest as go_build
WORKDIR /app/server
COPY . .

# RUN apk add apt
RUN apk add go
RUN apk add build-base
ENV PATH=$PATH:/usr/local/go/bin
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build main.go

FROM alpine:latest as server

WORKDIR /app

RUN apk add apt
RUN apk add bash
# # Need to find the file path of db

COPY ./scripts/init.sh ./init.sh
COPY ./models/database/sqlite-database.db /app/models/database/sqlite-database.db
COPY --from=go_build /app/server/main .

RUN ./init.sh
# EXPOSE 8080

CMD ["./main"]
