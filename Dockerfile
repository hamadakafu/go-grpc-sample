FROM golang:1.12-alpine3.9 as builder
WORKDIR /workspace

RUN apk update && apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app server/server.go

RUN ls -l -a


FROM alpine:3.9 
WORKDIR /workspace

COPY --from=builder /workspace/app app

CMD [ "./app" ]
