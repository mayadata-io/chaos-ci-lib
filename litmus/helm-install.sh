#!/bin/bash

set -e

#Provide default value of APP_NS
APP_NS=${APP_NS:="default"}
IMAGE_PULL_POLICY=${IMAGE_PULL_POLICY:="Always"}
EXPERIMENT_IMAGE=${EXPERIMENT_IMAGE:="litmuschaos/ansible-runner"}
EXPERIMENT_IMAGE_TAG=${EXPERIMENT_IMAGE_TAG:="1.4.1"}
CONTAINER_KILL_HELPER_IMAGE=${CONTAINER_KILL_HELPER_IMAGE:="gaiaadm/pumba"}
CONTAINER_KILL_HELPER_TAG=${CONTAINER_KILL_HELPER_TAG:="0.6.5"}
LIB=${LIB:="pumba"}

#Add chaos helm repository
helm repo add k8s-chaos https://litmuschaos.github.io/chaos-helm/
helm repo list
helm search repo k8s-chaos

#Install the kubernetes chaos experiments
helm install k8s k8s-chaos/kubernetes-chaos --set image.litmus.pullPolicy=${IMAGE_PULL_POLICY} \
--set image.litmus.repository=${EXPERIMENT_IMAGE} --set image.litmus.tag=${EXPERIMENT_IMAGE_TAG} \
--set image.pumba.repository=${CONTAINER_KILL_HELPER_IMAGE} --set image.pumba.libName=${LIB} \
--set image.pumba.tag=${CONTAINER_KILL_HELPER_TAG}  --namespace=${APP_NS}

#Checking the installation 
kubectl get chaosexperiments -n ${APP_NS}
