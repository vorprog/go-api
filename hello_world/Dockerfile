FROM registry.hub.docker.com/library/golang:alpine AS build

WORKDIR /src/
COPY ./source /src/
ARG BUILD_COMMIT
ARG BUILD_DATE_VERSION
RUN CGO_ENABLED=0 go build -ldflags "-X main.buildCommitLinkerFlag=$BUILD_COMMIT -X main.buildDateVersionLinkerFlag=$BUILD_DATE_VERSION" -o /bin/app

FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
