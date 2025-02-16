FROM golang:1.22 as builder

WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/. .

RUN CGO_ENABLED=0 go build -o main -a -ldflags '-extldflags "-static"' .

FROM alpine:latest  


RUN adduser -D user
USER user
WORKDIR /home/user
COPY --from=builder /app/main /home/user/
RUN mkdir notes
EXPOSE 5000
CMD ["./main"]