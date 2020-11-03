FROM registry.hub.docker.com/library/golang:alpine AS build

WORKDIR /src/
COPY ./source /src/
ARG BUILD_COMMIT
RUN CGO_ENABLED=0 go build -ldflags "-X main.buildCommitLinkerFlag=$BUILD_COMMIT -X main.buildDateVersionLinkerFlag=$(date +'%Y.%m.%d.%H.%M.%S')" -o /bin/app

FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
