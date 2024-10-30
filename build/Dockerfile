FROM golang:1.23-alpine AS dev
RUN apk add --no-cache git make
RUN adduser -D appuser
COPY . /src/
WORKDIR /src

ENV GO111MODULE=on
RUN --mount=type=cache,sharing=locked,id=gomod,target=/go/pkg/mod/cache \
    --mount=type=cache,sharing=locked,id=goroot,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux make build

FROM alpine:3.20
COPY --from=dev /src/out/scrapper /
ENTRYPOINT ["/scrapper"]
