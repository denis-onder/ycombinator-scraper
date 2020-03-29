#!/bin/bash

go build .
sudo mv ./ycombinator-scraper /usr/local/bin

echo "Moved the binary to /usr/local/bin"
echo "You can now use ycombinator-scraper from anywhere in your terminal!"