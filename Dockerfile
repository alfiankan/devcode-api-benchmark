FROM golang:1.18beta1-alpine as build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o devcode


FROM alpine:latest
RUN apk add dumb-init
COPY --from=build /app/devcode ./

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ./devcode

EXPOSE 3030
