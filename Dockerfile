FROM golang:1.21
LABEL org.opencontainers.image.source=https://github.com/ASullivan219/newsletter


# Copy Files to docker container /app
# Download all dependencies
WORKDIR /app
COPY . .
RUN go mod download

# Work in the /cmd directory
# build main.go
# Run the project exposing port 8080
WORKDIR /app/cmd
RUN go build -o .
EXPOSE 8080
CMD ["./cmd"]


