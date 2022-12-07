FROM golang:1.19-alpine AS dep
WORKDIR /src/
COPY . .
RUN cd cmd \
go get -d -v

FROM dep AS build
ARG VERSION "devel"
ARG GIT_COMMIT ""
ARG GIT_REF ""
WORKDIR /src/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'github.com/laghoule/k64dec/main.version=$VERSION' -X 'github.com/laghoule/k64dec/main.gitCommit=$GIT_COMMIT' -X 'github.com/laghoule/k64dec/main.gitRef=$GIT_REF'" -o k64dec cmd/main.go

FROM alpine:3.17
LABEL org.opencontainers.image.source https://github.com/laghoule/k64dec
COPY --from=build /src/k64dec /usr/bin/
ENTRYPOINT ["/usr/bin/k64dec"]
CMD ["-h"]