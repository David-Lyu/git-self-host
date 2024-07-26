#!/bin/bash

#Check os system
cat /etc/os-release;
if [ -f /etc/os-release ]; then
    . /etc/os-release
    OS=$NAME
    VER=$VERSION_ID
else
    echo "We do not support that os";
fi

# Install git server
# Run Go server? Here?
# We need atleast cmake file here to read main?
#
