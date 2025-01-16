#!/bin/bash

# set variables
PACKAGE="./proxyserver"        
OUTPUT_DIR="./output"         
AAR_NAME="proxyserver.aar"    

# output directory
mkdir -p $OUTPUT_DIR

# generate a single AAR supporting all architectures
echo "Generating a combined AAR for all architectures..."
gomobile bind -target=android -o "$OUTPUT_DIR/$AAR_NAME" $PACKAGE

if [ $? -eq 0 ]; then
    echo "Successfully generated the combined AAR: $OUTPUT_DIR/$AAR_NAME"
else
    echo "Failed to generate the AAR."
    exit 1
fi
