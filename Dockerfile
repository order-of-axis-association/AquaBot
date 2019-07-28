FROM golang:1.8

RUN  apt-get -yq update && \
     apt-get -yqq install ssh

RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

RUN /bin/bash -c "source secrets/secrets.sh"

COPY . /app
WORKDIR /app

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

