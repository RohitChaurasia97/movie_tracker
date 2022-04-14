FROM golang:1.17.5-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && apk add --no-cache bash git openssh make autoconf gcc libc-dev sudo procps curl

# Set the Current Working Directory inside the container
WORKDIR /var/www/movie_tracker

COPY go.mod go.sum ./

RUN go mod download -x

COPY . /var/www/movie_tracker/

RUN go build -v

# Expose port 9000 to the outside world
EXPOSE 9000 

CMD ["./movie_tracker","start"]

