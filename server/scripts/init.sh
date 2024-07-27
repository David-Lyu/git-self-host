#!/bin/bash

OS=""
# Check if it is from docker container
if [ DOCKER_ENV != "DOCKER" ]; then
    # Check if it is a linux distro, and if it is debian
    if [[ ! -f /etc/os-release ]] || [[ $(grep '^(NAME)=' /etc/os-release) == "Debian GNU/Linux" ]]; then
        echo "We do not support that os";
        exit 1;
    else
        OS=$(grep '^(NAME)=' /etc/os-release)
        echo $OS
        echo $(grep '^(NAME)=' /etc/os-release)
    fi
fi

echo $OS


# Install git server
# Run Go server? Here?
# We need atleast cmake file here to read main?
#
