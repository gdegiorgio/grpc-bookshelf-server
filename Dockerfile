FROM golang:1.22-alpine

ARG pretty=false
ARG HOST=localhost
ARG PORT=4500

RUN apk add --no-cache make

COPY .  .
RUN go mod download
RUN make build

CMD ["make", "run"]
