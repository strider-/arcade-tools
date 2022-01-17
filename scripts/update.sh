#!/bin/bash

local_dir=$(readlink -e "${0%/*}/..")

echo "Removing local changes & pulling down latest code..."
git reset HEAD --hard
git pull origin master

echo "Stopping Arcade Tools API Service..."
systemctl stop arcade-api.service

echo "Installing latest binaries..."
mkdir -p /opt/tools
cp $local_dir/bin/* /opt/tools

echo "Starting Arcade Tools API Service..."
systemctl start arcade-api.service

echo "Done."