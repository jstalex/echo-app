FROM golang:alpine AS build

WORKDIR /app

ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o main cmd/main.go

EXPOSE 80

FROM scratch

COPY --from=build app/main /bin/main

ENTRYPOINT ["/bin/main"]

