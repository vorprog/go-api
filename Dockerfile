FROM registry.hub.docker.com/library/golang:alpine AS build
RUN apk --update add ca-certificates git

WORKDIR /src/
COPY . /src/
ARG BUILD_COMMIT
RUN go get -d ./...
RUN export CURRENT_DATE_VERSION=$(date +'%Y.%m.%d.%H.%M.%S') && \
CGO_ENABLED=0 \
go build \
-ldflags "-X github.com/vorprog/go-api/util.BuildCommitLinkerFlag=$BUILD_COMMIT -X github.com/vorprog/go-api/util.BuildDateVersionLinkerFlag=$CURRENT_DATE_VERSION" \
-o /bin/app

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
