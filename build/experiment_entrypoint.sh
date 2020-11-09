#!/bin/bash

set -e

# Install litmus
if [ "$INSTALL_LITMUS" == "true" ];then
    #running install-litmus go binary 
    ./install-litmus
fi

#execute desired chaosexperiment
if [ ! -z "$EXPERIMENT_NAME" ];then
    #running experiment go binary 
    ./$EXPERIMENT_NAME
else
    echo "No experiment to run. Please setup EXPERIMENT_NAME env to run an experiment"
fi
