FROM golang:1.8

COPY . /app
WORKDIR /app

# Create ssh dir and drop aquabot key supplied by cloudbuild args
RUN mkdir /root/.ssh/
RUN echo "${AQUABOT_SECRETS_KEY}" > /root/.ssh/id_rsa

RUN go get -d -v ./...
RUN /bin/bash -c "source secrets/secrets.sh"

RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

