FROM registry.hub.docker.com/library/golang:alpine AS build
RUN apk add git

WORKDIR /src/
COPY ./source /src/
ARG BUILD_COMMIT
RUN go get -d ./...
RUN export CURRENT_DATE_VERSION=$(date +'%Y.%m.%d.%H.%M.%S') && \
CGO_ENABLED=0 \
go build \
-ldflags "-X server.buildCommitLinkerFlag=$BUILD_COMMIT -X server.buildDateVersionLinkerFlag=$CURRENT_DATE_VERSION" \
-o /bin/app

FROM scratch
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
