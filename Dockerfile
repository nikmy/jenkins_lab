FROM golang:1.21.4-alpine

WORKDIR /app

ADD . .

RUN go build ./cmd/
RUN rm go.mod

EXPOSE 80
ENTRYPOINT [ "./main", "-port=80" ]