FROM registry.hub.docker.com/library/golang:alpine AS build

WORKDIR /src/
COPY ./source /src/
ARG GIT_CURRENT_COMMIT_HASH
RUN CGO_ENABLED=0 go build -ldflags "-X main.gitCommitHash=$GIT_CURRENT_COMMIT_HASH" -o /bin/app

FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
