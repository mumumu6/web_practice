# Use an official Python runtime as the base image
FROM golang:1.22

# Set the working directory in the container
WORKDIR /app

# Install the dependencies
RUN go install github.com/air-verse/air@latest

# Expose the port on which the application will run
EXPOSE 8081

# Define the command to run the application
CMD ["air", "-c", ".air.toml"]