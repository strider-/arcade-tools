#!/bin/bash

if [[ `id -u` -ne 0 ]] ; then 
    echo "Install script should be ran as root, run 'sudo make install' instead."
    exit 1
fi

local_dir=$(readlink -e "${0%/*}/..")
script_dir=$(readlink -e "${0%/*}")

echo "Installing arcade tools..."
mkdir -p /opt/tools
cp $local_dir/bin/* /opt/tools

echo "Installing relay power service..."
cp -u $script_dir/relay-power.service /etc/systemd/system
cp -u $script_dir/powerup.sh /opt/tools
chmod +x /opt/tools/powerup.sh
systemctl enable relay-power.service

echo "Installing Arcade Tools API service..."
cp -u $script_dir/arcade-api.service /etc/systemd/system
systemctl enable arcade-api.service
systemctl start arcade-api.service

echo "Done"