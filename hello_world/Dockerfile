FROM registry.hub.docker.com/library/golang:alpine AS build

WORKDIR /src/
COPY ./source /src/
ARG GIT_COMMIT_HASH
ARG GIT_COMMIT_TIMESTAMP
RUN CGO_ENABLED=0 \
  && go build \
  -ldflags "-X main.compiledGitCommitHash=$GIT_COMMIT_HASH" \
  -ldflags "-X main.compiledGitCommitTimestamp=$GIT_COMMIT_TIMESTAMP" \
  -o /bin/app

FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
