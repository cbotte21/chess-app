#!/usr/bin/env bash

cd docker-compose
eval $(minikube docker-env)
docker compose build
#kubectl --rollout restart
