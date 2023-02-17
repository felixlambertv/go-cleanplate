FROM golang:1.19-alpine as base
WORKDIR /app

ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    git \
    && update-ca-certificates

# ----------- DEV DOCKERFILE ------------
FROM base as dev

ENV GO111MODULE="on"

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/vektra/mockery/v2@latest

EXPOSE 8000
EXPOSE 2345

# Run the air command in the directory where our code will live
ENTRYPOINT ["air"]

# ----------- PROD DOCKERFILE -----------

FROM base as prod

COPY . .

RUN go mod download

RUN go build -mod=readonly -v -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=prod /app/main/ .

EXPOSE 8000

CMD ["./cmd/app/main"]

