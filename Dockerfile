FROM golang:latest

RUN mkdir -p /home/app

WORKDIR /home/app

COPY .  /home/app
# Build the Go application
RUN go build -o myapp .

# Expose the port your application will run on
EXPOSE 8080

# Specify the command to run on container start
CMD ["./myapp"]

