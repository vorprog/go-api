FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY ./source /src/
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch
ENV PORT 8080
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
EXPOSE 8080