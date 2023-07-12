FROM golang:1.20

ARG BIN

ENV BIN=${BIN}
ENV PORT=${PORT}

# Create a location in the container for the source code.
RUN mkdir -p /app

# Copy the module files first and then download the dependencies. If this
# doesn't change, we won't need to do this again in future builds.
COPY go.* /app/

WORKDIR /app
RUN go mod download
RUN go mod verify

# Copy the source code into the container.
COPY pkg pkg
COPY services/"$BIN" services/"$BIN"

# Install air for development
RUN go install github.com/cosmtrek/air@latest

# Expose port
EXPOSE $PORT

# Start the application.
CMD ["air", "-c", "./services/account/.air.toml"]