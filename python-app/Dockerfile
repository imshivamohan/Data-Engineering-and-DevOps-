# Stage 1 - Install build dependencies using Ubuntu
FROM ubuntu:latest AS builder

# Set non-interactive mode to avoid prompts during installation
ENV DEBIAN_FRONTEND=noninteractive

# Update package list and install Python3, Pip, and venv
RUN apt update && apt install -y python3 python3-pip python3-venv

# Set the working directory inside the container
WORKDIR /app

# Copy application dependencies file
COPY requirements.txt .

# Create a virtual environment and install dependencies inside it
RUN python3 -m venv venv && \
    . venv/bin/activate && \
    pip install --no-cache-dir -r requirements.txt && \
    find /app/venv \( -type d -a -name test -o -name tests \) -o \( -type f -a -name '*.pyc' -o -name '*.pyo' \) -exec rm -rf '{}' +

# Stage 2 - Use Alpine as the final runtime image
FROM alpine:latest

# Install Python3 and required dependencies in Alpine
RUN apk add --no-cache python3 py3-pip

# Set the working directory inside the container
WORKDIR /app

# Copy the virtual environment and application files from the builder stage
COPY --from=builder /app /app
COPY app.py .

# Set environment variables
ENV PATH="/app/venv/bin:$PATH"

# Expose the Flask port
EXPOSE 5000

# Run the Flask app inside the virtual environment
CMD ["python3", "app.py"]

