#!/bin/bash

# Array of Go project directories
goProjects=("../vsys-empms-db" "../vsys-empms-rest" "../vsys-empms-web")
# Uncomment the line below to build only the web project
# goProjects=("../vsys-empms-web")

# Save the current directory to return after processing each project
original_dir=$(pwd)

for project_dir in "${goProjects[@]}"
do
   # Display project directory in uppercase
   echo "-------------------- ${project_dir} ------------------" | tr '[a-z]' '[A-Z]'

   # Navigate to the project directory
   cd "$project_dir" || { echo "Failed to enter directory: $project_dir"; exit 1; }

   # Clean up any previous build artifacts
   go clean

   # Ensure dependencies are tidy and up-to-date
   go mod tidy

   # Build the Go project for Linux, with CGO disabled
   CGO_ENABLED=0 GOOS=linux go build || { echo "Build failed in directory: $project_dir"; exit 1; }

   # Return to the original directory
   cd "$original_dir"
done

echo "Build completed for all projects."
