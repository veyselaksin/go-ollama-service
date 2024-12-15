#nginx/Dockerfile
FROM nginx:latest

# Create directory for certificates
RUN mkdir -p /etc/nginx/certs

# Copy nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Copy SSL certificates (you'll need to create these)
COPY certs/server.crt /etc/nginx/certs/
COPY certs/server.key /etc/nginx/certs/
COPY certs/ca.crt /etc/nginx/certs/

EXPOSE 443