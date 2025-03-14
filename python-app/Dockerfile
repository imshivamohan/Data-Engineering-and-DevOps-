# Use Ubuntu as base image
FROM ubuntu:latest

# Set non-interactive mode to avoid prompts during installation
ENV DEBIAN_FRONTEND=noninteractive

# Update package list and install Python3, Pip, and venv
RUN apt update && apt install -y python3 python3-pip python3-venv

# Set the working directory inside the container
WORKDIR /app

# Copy application files
COPY . .

# Create a virtual environment and install dependencies inside it
RUN python3 -m venv venv && \
    . venv/bin/activate && \
    pip install --no-cache-dir -r requirements.txt

# Expose the Flask port
EXPOSE 5000

# Run the Flask app inside the virtual environment
CMD ["sh", "-c", ". venv/bin/activate && python3 app.py"]

