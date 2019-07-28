FROM golang:1.8

RUN  apt-get -yq update && \
     apt-get -yqq install ssh

# Create ssh dir and drop aquabot key supplied by cloudbuild args
RUN mkdir /root/.ssh/
RUN cp id_rsa /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa

RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

RUN git submodule update --init --recursive
RUN /bin/bash -c "source secrets/secrets.sh"

COPY . /app
WORKDIR /app

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

