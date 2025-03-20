# Build stage (compiling dependencies)
FROM python:3.11-alpine AS builder
WORKDIR /app

# Install build dependencies
COPY requirements.txt .
RUN pip install --user --no-cache-dir -r requirements.txt

# Copy application code
COPY . .

# Final stage (minimal runtime environment)
FROM python:3.11-alpine
WORKDIR /app

# Copy only necessary parts from the builder stage
COPY --from=builder /root/.local /root/.local
COPY app.py .

# Set the path to include user-installed packages
ENV PATH=/root/.local/bin:$PATH

# Set a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

CMD ["python", "app.py"]
