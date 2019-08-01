FROM golang:1.8

RUN ls -al .

COPY . /app
COPY .git/ /app/.git/

RUN ls -al /app

# Copy in secrets submodule
# This is kinda janky cuz docker removes the .git folder, meaning we can't `git submodule update`
# Instead, we have to clone AquaBot separately inside the workdir then move the secrets submodule back out. It's jank.
RUN cp -r /app/AquaBot/secrets/* /app/secrets
RUN rm -rf /app/AquaBot

RUN /bin/bash -c "source /app/secrets/secrets.sh"

RUN mkdir /root/.ssh
RUN cp /app/secrets/id_rsa /root/.ssh/
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

WORKDIR /app

RUN echo "$REPO_REVISION" > /app/secrets/repo_revision

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

