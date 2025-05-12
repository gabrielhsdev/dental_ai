FROM python:3.11-slim

# Set working directory
WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y \
    git \
    curl \
    libglib2.0-0 \
    libsm6 \
    libxext6 \
    libxrender-dev \
    ffmpeg \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Define the diagnostics service folder as a build argument
ARG DIAGNOSTICS_SERVICE_FOLDER

# Copy application files
COPY backend/${DIAGNOSTICS_SERVICE_FOLDER}/. .
COPY .env .

# Install Python dependencies (replace macOS-specific packages for our linux base image)
RUN pip install --upgrade pip && \
    pip install --no-cache-dir \
        tensorflow \
        ultralytics \
        opencv-python \
        huggingface-hub \
        Pillow \
        flask \
        mypy \
        types-Pillow \
        types-Flask

# Expose the port (can be overridden at runtime via ENV or CMD)
ARG DIAGNOSTICS_SERVICE_PORT=8083
EXPOSE ${DIAGNOSTICS_SERVICE_PORT}
EXPOSE 8083

# Run the application
CMD ["python", "main.py", "--port", "8083"]
