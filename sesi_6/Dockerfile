# Stage 1 build
FROM golang:1.22 AS builder

# Set the current working directory inside the container
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

# Copy the source from the current directory to the Working Directory inside the container
COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o myapp ./cmd/gin/main.go

FROM scratch

COPY --from=builder /app/myapp /myapp

ENTRYPOINT [ "/myapp" ]

EXPOSE 5000