FROM golang:1.8

ARG REPO_REVISION="N/A"
ENV COMMIT_SHA=$REPO_REVISION

RUN echo "SHA IS $COMMIT_SHA"

COPY . /app

RUN cp -r /app/AquaBot/secrets/* /app/secrets
RUN rm -rf /app/AquaBot

RUN mkdir /root/.ssh
RUN cp /app/secrets/id_rsa /root/.ssh/
RUN cp /app/secrets/known_hosts /root/.ssh/

WORKDIR /app

# Copy in secrets submodule
# This is the fucking worst. From what I'm able to tell, cloudbuild auto-triggered builds
# don't actually have a way of ignoring... or explicitly including, files into the build.
# .gcloudignore seems to only be used with `gcloud` commands such as `gcloud submit`.
# As such I think, at least until google adds an alt, that I'm forced to clone repo in a cloudbuild
# step and then copy the files into the docker container workspace.

# Ugh.
RUN /bin/bash -c "source /app/secrets/secrets.sh"

RUN go get -d -v ./...
RUN make build -C /app

ENTRYPOINT ["/app/entrypoint.sh"]

