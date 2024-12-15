#!/bin/bash

# Create directory for certificates
mkdir -p nginx/certs

# Generate CA private key and certificate
openssl genrsa -out nginx/certs/ca.key 4096
openssl req -x509 -new -nodes -key nginx/certs/ca.key -sha256 -days 3650 -out nginx/certs/ca.crt \
  -subj "/C=US/ST=State/L=City/O=Organization/CN=My CA"

# Generate server private key and CSR
openssl genrsa -out nginx/certs/server.key 2048
openssl req -new -key nginx/certs/server.key -out nginx/certs/server.csr \
  -subj "/C=US/ST=State/L=City/O=Organization/CN=localhost"

# Sign server certificate with CA
openssl x509 -req -in nginx/certs/server.csr -CA nginx/certs/ca.crt -CAkey nginx/certs/ca.key \
  -CAcreateserial -out nginx/certs/server.crt -days 365 -sha256

# Generate client private key and CSR
openssl genrsa -out nginx/certs/client.key 2048
openssl req -new -key nginx/certs/client.key -out nginx/certs/client.csr \
  -subj "/C=US/ST=State/L=City/O=Organization/CN=client"

# Sign client certificate with CA
openssl x509 -req -in nginx/certs/client.csr -CA nginx/certs/ca.crt -CAkey nginx/certs/ca.key \
  -CAcreateserial -out nginx/certs/client.crt -days 365 -sha256

# Set proper permissions
chmod 600 nginx/certs/*.key