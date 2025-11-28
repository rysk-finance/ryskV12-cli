#!/bin/bash

# Define the target platforms
platforms=(
  "linux/amd64"
  "linux/arm64"
  "darwin/arm64"
)

# Build for each platform
for platform in "${platforms[@]}"; do
  os=$(echo "$platform" | cut -d'/' -f1)
  arch=$(echo "$platform" | cut -d'/' -f2)
  output="ryskV12-$os-$arch"

  if [[ "$os" == "windows" ]]; then
    output="$output.exe"
  fi

  echo "Building for $os/$arch..."
  GOOS="$os" GOARCH="$arch" go build -ldflags="-X 'main.Version=3.0.0'" -o "$output"
  if [ $? -ne 0 ]; then
    echo "Failed to build for $os/$arch"
    exit 1
  fi
  echo "Build for $os/$arch successful: $output"
done

echo "All builds completed."