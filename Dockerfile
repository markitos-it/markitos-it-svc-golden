FROM golang:1.25-alpine AS builder

WORKDIR /build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

# Instala dependencias de build (incluye bash) + verifica que bash quedó instalado
RUN apk add --no-cache make git protobuf bash \
 && echo "== tool versions ==" \
 && go version \
 && protoc --version \
 && echo "== bash check ==" \
 && which bash \
 && ls -la /bin/bash \
 && bash --version

# Instala plugins de protoc para Go y verifica que existen en PATH
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
 && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
 && echo "== protoc plugins check ==" \
 && which protoc-gen-go \
 && which protoc-gen-go-grpc

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copia el repo
COPY . .

# Debug: ver qué hay en proto antes de generar
RUN echo "== proto before make proto ==" && ls -la proto

# Genera .pb.go / .pb.grpc.go
RUN make proto

# Debug: confirmar que se generaron los .go en proto/
RUN echo "== proto after make proto ==" && ls -la proto

# Build
RUN go build -v -x -o app ./cmd/app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /build/app .

EXPOSE 3000

CMD ["./app"]