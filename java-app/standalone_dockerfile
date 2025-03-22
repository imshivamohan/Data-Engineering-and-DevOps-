# Use lightweight Alpine as the base image
FROM alpine:latest

# Install OpenJDK 11 and Maven
RUN apk add --no-cache openjdk11 maven

# Set the working directory inside the container
WORKDIR /app

# Copy the project files into the container
COPY . .

# Build the Java application using Maven
RUN mvn clean package

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["java", "-jar", "target/demo-app-1.0.0.jar"]

