FROM registry.hub.docker.com/library/golang:alpine AS build

WORKDIR /src/
COPY ./source /src/
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" /bin/wasm_exec.js
ARG BUILD_COMMIT
RUN GOOS=js GOARCH=wasm CGO_ENABLED=0 go build -ldflags "-X main.buildCommitLinkerFlag=$BUILD_COMMIT -X main.buildDateVersionLinkerFlag=$(date +'%Y.%m.%d.%H.%M.%S')" -o /bin/app.wasm

FROM nginx:latest
COPY ./index.html /usr/share/nginx/html/index.html
COPY --from=build /bin/wasm_exec.js /usr/share/nginx/html/wasm_exec.js
COPY --from=build /bin/app.wasm /usr/share/nginx/html/app.wasm
