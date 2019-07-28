FROM golang:1.8

COPY . /app

RUN ls -al /app/AquaBot/
RUN ls -al /app/AquaBot/secrets/

RUN cp /app/AquaBot/secrets/ /app/secrets/
RUN rm -rf /app/AquaBot

# Look, I cannot for the life of me figure out how to make cloudbuild
# clone the /app/secrets/ directory correctly. It seems to get rid of /app/.git so I can't
# run submodule init/updates :-/
RUN /bin/bash -c "source /app/secrets/secrets.sh"

RUN mkdir /root/.ssh
RUN cp /app/secrets/id_rsa /root/.ssh/
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

WORKDIR /app

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

