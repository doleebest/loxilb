#!/bin/bash

echo "#########################################"
echo "Removing testbed"
echo "#########################################"

if [ "$1" ]; then
  KUBECONFIG="$1"
fi

source ../common.sh

sudo kubectl $KUBECONFIG delete -f nginx-svc-lb1.yml >> /dev/null 2>&1
sudo kubectl $KUBECONFIG delete -f sctp-svc-lb.yml >> /dev/null 2>&1
sudo kubectl $KUBECONFIG delete -f kube-loxilb.yml >> /dev/null 2>&1
sudo kubectl $KUBECONFIG delete -f nginx-svc-lb.yml >> /dev/null 2>&1
sudo kubectl $KUBECONFIG delete -f nginx.yml >> /dev/null 2>&1
sudo kubectl $KUBECONFIG delete -f https://github.com/loxilb-io/loxi-ccm/raw/master/manifests/loxi-ccm-k3s.yaml >> /dev/null 2>&1

disconnect_docker_hosts user r1
disconnect_docker_hosts user r3
disconnect_docker_hosts r1 sw1
disconnect_docker_hosts r3 sw1 
disconnect_docker_hosts sw1 llb1
disconnect_docker_hosts sw1 llb2
disconnect_docker_hosts llb1 r2
disconnect_docker_hosts llb2 r2
disconnect_docker_hosts r2 ep1
disconnect_docker_hosts r2 ep2
disconnect_docker_hosts r2 ep3

delete_docker_host ka_llb1
delete_docker_host ka_llb2
delete_docker_host llb1
delete_docker_host llb2
delete_docker_host user
delete_docker_host r1
delete_docker_host r2
delete_docker_host r3
delete_docker_host sw1
delete_docker_host ep1
delete_docker_host ep2
delete_docker_host ep3
sudo ip link del esysllb1
sudo ip link del esysllb2

# If k3s setup exists, remove it
if [[ -f "/usr/local/bin/k3s-uninstall.sh" ]]; then
  /usr/local/bin/k3s-uninstall.sh
fi

sudo apt-get remove bird2 --yes
#docker image rm loxilb-io/sctp-server
#docker images -a | grep "loxilb-io" | awk '{print $3}' | xargs docker rmi
echo "#########################################"
echo "Removed testbed"
echo "#########################################"
