#!/bin/bash

target=$1
local_dir=$(readlink -e "${0%/*}/..")
dest=/var/www/arcade-fe

if [ -z "$target" ]; then
    echo "Missing <user@host> deployment target"
    exit 1
fi

echo "Building the front-end..."
cd $local_dir/ui
npm run build

echo "Moving current deployment to previous..."
ssh -T $target << EOF
    [ -d $dest/public.prev ] && rm -rf $dest/public.prev
    [ -d $dest/public ] && mv $dest/public $dest/public.prev
    exit
EOF

echo "Deploying to server..."
scp -prq ./build/ $target:$dest/public

echo "Done"