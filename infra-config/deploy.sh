#!/bin/bash

kind create cluster --name educhain
kubectl wait --for=condition=Ready --timeout=300s nodes --all

helm repo add cloudnative-pg https://cloudnative-pg.io/charts
helm repo update
helm upgrade --install cnpg \
    --namespace cnpg-system \
    --create-namespace \
    cloudnative-pg/cloudnative-pg

kubectl wait --for=condition=Available --timeout=300s deployment/cnpg-cloudnative-pg -n cnpg-system

kubectl apply -f database/
kubectl wait --for=condition=Ready --timeout=300s cluster/postgres

kubectl apply -f ipfs-node.yaml
kubectl wait --for=condition=Ready --timeout=300s pod/ipfs-node

kubectl apply -f geth-node.yaml
kubectl wait --for=condition=Ready --timeout=300s pod/geth-node
echo "Retriving master key from Geth node. Add it to the smart contract api config map"
kubectl exec geth-node -- /bin/sh -c "cat .ethereum/keystore/*" > key-file.json

kubectl apply -f ipfs-api.yaml
kubectl wait --for=condition=Available --timeout=300s deployment/ipfs-api

echo "Deployment completed successfully."
echo "Please deploy the smart contract api manually using the master key from master-key.json"
echo "kubectl apply -f smart-contract-api.yaml"
