FROM golang:1.8

RUN  apt-get -yq update && \
     apt-get -yqq install ssh

COPY . /app
WORKDIR /app

# Create ssh dir and drop aquabot key supplied by cloudbuild args
RUN mkdir /root/.ssh/
RUN echo "${AQUABOT_SECRETS_KEY}" > /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa

RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN echo "AQUABOT SECRETS CONTENT:"
RUN echo "${AQUABOT_SECRETS_KEY}"

RUN go get -d -v ./...
RUN /bin/bash -c "source secrets/secrets.sh"

RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

