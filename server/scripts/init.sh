#!/bin/bash

OS=""
# Check if it is from docker container
if [[ $DOCKER_ENV != "DOCKER" ]]; then
    # Check if it is a linux distro, and if it is debian
    if [ -f /etc/os-release ]; then
        OS=$(. /etc/os-release; echo "$NAME")
        if [[ $OS != "Debian GNU/Linux" ]]; then
            echo "We do not support the linux distro";
            exit 1;
        fi
    else
        echo "We do not support that os";
        exit 1;
    fi
fi

#Install dependencies if needed
# Need to check for dependencies
if [[ $OS == "Debian GNU/Linux" ]]; then
    apt install git;
elif [[ DOCKER_ENV == "DOCKER" ]]; then
    apk add git;
fi

#Run the go backend
if [[ DOCKER_ENV == "DOCKER" ]]; then
    ./main
else
    #Grabs path of script and goes back a level where main should live
    $(dirname $0)/../main
fi
