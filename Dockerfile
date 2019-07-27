FROM golang:1.8

COPY . /app
WORKDIR /app

RUN go get -d -v ./...
RUN /bin/bash -c "source secrets/secrets.sh"

RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

