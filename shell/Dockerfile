FROM registry.hub.docker.com/library/golang:alpine AS build

WORKDIR /src/
COPY ./source /src/
RUN CGO_ENABLED=0 go build -o /bin/shell
# TODO: specify architecture for go build

FROM registry.hub.docker.com/library/golang:alpine AS print_path_binary

WORKDIR /src/
COPY ./print_path /src/
RUN CGO_ENABLED=0 go build -o /bin/print_path

FROM scratch
ENV PATH="/bin:/usr/bin:/usr/local/bin"
ENV HOME=/home
ENV PS1="> "
COPY --from=build /bin/shell /bin/shell
COPY --from=print_path_binary /bin/print_path /bin/ppath
ENTRYPOINT ["/bin/shell"]
