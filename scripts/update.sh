#!/bin/bash

if [[ `id -u` -ne 0 ]] ; then 
    echo "Update script should be ran as root, run 'sudo make update' instead."
    exit 1
fi

echo "Fetching latest code"
sudo -u $SUDO_USER -E git pull origin master

local_dir=$(readlink -e "${0%/*}/..")

echo "Stopping Arcade Tools API Service..."
systemctl stop arcade-api.service

echo "Installing latest binaries..."
mkdir -p /opt/tools
cp $local_dir/bin/* /opt/tools

echo "Starting Arcade Tools API Service..."
systemctl start arcade-api.service

echo "Done."