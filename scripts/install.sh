#!/bin/bash

local_dir=$(readlink -e "${0%/*}/..")
script_dir=$(readlink -e "${0%/*}")

echo "Installing arcade tools..."
mkdir -p /opt/tools
cp $local_dir/bin/* /opt/tools

echo "Installing relay power service..."
cp -u -p $script_dir/relay-power.service /etc/systemd/system
cp -u -p $script_dir/powerup.sh /opt/tools
systemctl enable relay-power.service

echo "Installing Arcade Tools API service..."
cp -u -p $script_dir/arcade-api.service /etc/systemd/system
systemctl enable arcade-api.service
systemctl start arcade-api.service

echo "Done"