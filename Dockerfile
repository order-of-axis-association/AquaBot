FROM golang:1.8

COPY . /app

RUN mkdir /root/.ssh

RUN ls -al /root/.ssh
RUN ls -al /app
RUN ls -al /app/secrets
RUN ls -al /

RUN mv /app/known_hosts /root/.ssh/known_hosts
RUN mv /app/id_rsa /root/.ssh/id_rsa

WORKDIR /app

RUN git submodule update --init --recursive

# Copy in secrets submodule
# This is kinda janky cuz docker removes the .git folder, meaning we can't `git submodule update`
# Instead, we have to clone AquaBot separately inside the workdir then move the secrets submodule back out. It's jank.
RUN /bin/bash -c "source /app/secrets/secrets.sh"

RUN echo "$REPO_REVISION" > /app/secrets/repo_revision

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

