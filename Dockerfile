FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /lite-social-presence-system
EXPOSE 80 81
CMD ["/lite-social-presence-system"]