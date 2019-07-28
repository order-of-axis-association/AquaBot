FROM golang:1.8

RUN  apt-get -yq update && \
     apt-get -yqq install ssh

RUN /bin/bash -c "source secrets/secrets.sh"

COPY . /app
WORKDIR /app

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

