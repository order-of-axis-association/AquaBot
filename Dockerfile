FROM golang:1.8

RUN ls -al .
RUN ls -al /root/.ssh
RUN ls -al /root/.ssh/secrets

RUN /bin/bash -c "source /root/.ssh/secrets/secrets.sh"

COPY . /app
WORKDIR /app

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

