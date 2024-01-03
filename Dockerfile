FROM golang:1.21-alpine AS dep
WORKDIR /src/
COPY . .
RUN cd cmd \
	go get -d -v

FROM dep AS build
ARG VERSION "devel"
ARG GIT_COMMIT ""
ARG GIT_REF ""
WORKDIR /src/
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.version=$VERSION' -X 'main.gitCommit=$GIT_COMMIT' -X 'main.gitRef=$GIT_REF'" -o k64dec cmd/main.go

FROM alpine:3.19
LABEL org.opencontainers.image.source https://github.com/laghoule/k64dec
COPY --from=build /src/k64dec /usr/bin/
USER nobody
ENTRYPOINT ["/usr/bin/k64dec"]
