#!/bin/bash
set -e  # Exit on any error

echo "Starting Pelias placeholder data extraction..."
npm run extract

echo "Extraction completed. Starting build process..."
npm run build

echo "Build completed successfully!"