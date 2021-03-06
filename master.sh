#!/bin/bash
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF
yum install maven git -y
setenforce 0
sed -i 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config
swapoff -a
echo net.bridge.bridge-nf-call-ip6tables = 1 >> /etc/sysctl.conf
echo net.bridge.bridge-nf-call-iptables = 1 >> /etc/sysctl.conf
modprobe br_netfilter
sysctl -p
yum install -y docker kubelet kubeadm kubectl kubernetes-cni
systemctl enable docker && systemctl start docker
systemctl enable kubelet && systemctl start kubelet
kubeadm init
#--pod-network-cidr=10.244.0.0/16
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
kubectl taint nodes --all node-role.kubernetes.io/master-
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml
rm -rf user-service
git clone https://github.com/agbankar/user-service.git
cd user-service && mvn clean install
docker build -t shahpriti919/user-service .
docker login -u="shahpriti919" -p="$urekha12"
docker push shahpriti919/user-service
kubectl delete deployment user-deployment && kubectl delete svc user-service-svc
kubectl apply -f user-deployment.yml && kubectl apply -f user-service.yml
kubectl get deployments
kubectl get svc
kubeadm token create --print-join-command
