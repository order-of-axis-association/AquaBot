FROM golang:1.8

COPY . /app

RUN pwd
RUN ls -al .
RUN ls -al src
RUN ls -al bin
RUN ls -al /
RUN ls -al /app
RUN ls -al /root/.ssh
RUN ls -al /root/.ssh/secrets

RUN /bin/bash -c "source /root/.ssh/secrets/secrets.sh"

WORKDIR /app

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

