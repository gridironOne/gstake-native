ARG IMG_TAG=latest

# Compile the gaiad binary
FROM golang:1.19-alpine AS builder
WORKDIR /go/src/github.com/gridironOne/gstake-native
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES
RUN CGO_ENABLED=0 make install

# Add to a distroless container
FROM gcr.io/distroless/cc:$IMG_TAG
ARG IMG_TAG
COPY --from=builder /go/bin/gstaked /usr/local/bin/gstaked
EXPOSE 26656 26657 1317 9090

ENTRYPOINT ["gstaked", "start"]
