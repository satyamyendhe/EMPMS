# Use an empty base image
FROM scratch

# Copy the statically compiled Go binary into the image
COPY vsys.empms.rest /vsys.empms.rest

# Expose port 7100 for the application
EXPOSE 7100

# Set the default command to run the binary
CMD ["/vsys.empms.rest"]
