# Base image
FROM envoyproxy/envoy:v1.22.0

# Copy the Envoy configuration file
COPY envoy.yaml /etc/envoy/envoy.yaml

# Expose Envoy's port
EXPOSE 10000

CMD ["envoy", "-c", "/etc/envoy/envoy.yaml", "--service-cluster", "api-gateway"]
