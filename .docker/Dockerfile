# syntax=docker/dockerfile:1
############################
# STEP 1 build executable binary
############################
FROM nafigat0r/go:1.24.3 AS builder

ARG LD_FLAGS
# https://stackoverflow.com/questions/50126741/how-to-remove-intermediate-images-from-a-build-after-the-build
LABEL stage=builder

WORKDIR /opt/ghost

COPY . .

RUN CGO_ENABLED=0 go build -a \
	-ldflags "${LD_FLAGS}" \
	-o /go/bin/ghost \
	./cmd/main.go

############################
# STEP 2 build a small image
############################
FROM scratch AS production

ARG PROJECT_REVISION
LABEL org.opencontainers.image.revision="$PROJECT_REVISION" \
	org.opencontainers.image.authors="alex@itvault.info" \
	org.opencontainers.image.vendor="Alexander Yancharuk" \
	org.opencontainers.image.licenses=MIT \
	org.opencontainers.image.source="https://github.com/nafigator/ghost/blob/${PROJECT_REVISION}/.docker/Dockerfile"

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/ghost /usr/local/sbin/ghost

WORKDIR /var/ghost

USER nobody

ENTRYPOINT ["/usr/local/sbin/ghost"]
