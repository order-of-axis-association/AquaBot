FROM golang:1.8


RUN /bin/bash -c "source /root/.ssh/secrets/secrets.sh"

COPY . /app
WORKDIR /app

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

