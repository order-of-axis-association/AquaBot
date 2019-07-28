FROM golang:1.8

COPY . /app

RUN ls -al /app/AquaBot/
RUN ls -al /app/AquaBot/secrets/

# Look, I cannot for the life of me figure out how to make cloudbuild
# clone the /app/secrets/ directory correctly. It seems to get rid of /app/.git so I can't
# run submodule init/updates :-/
RUN /bin/bash -c "source /app/AquaBot/secrets/secrets.sh"

RUN mkdir /root/.ssh
RUN cp /app/AquaBot/secrets/id_rsa /root/.ssh/

WORKDIR /app

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

