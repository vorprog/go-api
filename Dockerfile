FROM arm64v8/golang:1.20-alpine AS build-artifact-stage
RUN apk --update add ca-certificates curl git jq

# ENV SOPS_RELEASE_URL "https://api.github.com/repos/mozilla/sops/releases/latest"
# RUN curl --silent $SOPS_RELEASE_URL | jq -r '.assets[] | select(.browser_download_url | contains(".linux")).browser_download_url' >> /tmp/sops_download_url
# RUN cat /tmp/sops_download_url
# RUN curl -L $(cat /tmp/sops_download_url) --output /bin/sops
# RUN chmod +x /bin/sops

WORKDIR /src/
COPY go.mod go.sum /src/
RUN go get -d ./...

COPY . /src/
ARG BUILD_COMMIT
RUN export CURRENT_DATE_VERSION=$(date --utc +'%Y.%m.%d.%H.%M.%S') && \
CGO_ENABLED=0 \
go build \
-ldflags "-X github.com/vorprog/go-api/util.BuildCommitLinkerFlag=$BUILD_COMMIT -X github.com/vorprog/go-api/util.BuildDateVersionLinkerFlag=$CURRENT_DATE_VERSION" \
-o /bin/app

FROM scratch
COPY --from=build-artifact-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-artifact-stage /bin/ /bin/
ENTRYPOINT ["/bin/app"]
