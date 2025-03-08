##############################
FROM golang:1.24-alpine AS build

ARG VERSION "devel"
ARG GIT_COMMIT ""
ARG GIT_REF ""

WORKDIR /src

RUN --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  --mount=type=cache,target=/root/.cache/go-build \ 
  --mount=type=cache,target=/go/pkg \
  go mod download

RUN --mount=type=bind,source=.,target=.  \
  --mount=type=cache,target=/root/.cache/go-build \ 
  --mount=type=cache,target=/go/pkg \
  CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.version=$VERSION' -X 'main.gitCommit=$GIT_COMMIT' -X 'main.gitRef=$GIT_REF'" -o /tmp/k64dec cmd/main.go

##############################
FROM scratch

ARG VERSION

LABEL org.opencontainers.image.title="k64dec" \
  org.opencontainers.image.vendor="laghoule" \
  org.opencontainers.image.licenses="GPLv3" \
  org.opencontainers.image.version="${VERSION}" \
  org.opencontainers.image.description="Little tool to print the decoded base64 of a Kubernetes secret" \
  org.opencontainers.image.url="https://github.com/laghoule/k64dec/README.md" \
  org.opencontainers.image.source="https://github.com/laghoule/k64dec" \
  org.opencontainers.image.documentation="https://github.com/laghoule/k64dec/README.md"

USER 65534

COPY --link --from=build /tmp/k64dec /bin/k64dec

ENTRYPOINT ["/bin/k64dec"]