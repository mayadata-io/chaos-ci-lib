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
    exit 1
fi

# Uninstall litmus
if [ "$UNINSTALL_LITMUS" == "true" ];then
    #running uninstall-litmus go binary 
    ./uninstall-litmus
fi
