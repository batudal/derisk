#!/bin/bash

git pull
echo "Pulled latest changes from git"
make build
echo "Built app"
cd app && docker build -f app.dockerfile -t batudal/de-risk-app:0.0.1 . && docker push batudal/de-risk-app:0.0.1
echo "Built and pushed app image"
docker service update --image batudal/de-risk-app:0.0.1 app_de-risk
echo "Updated app service"
