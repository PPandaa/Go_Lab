#!/bin/bash
CONTAINER="GoLab"
DOCKER_REPO="iiicondor/$CONTAINER"
VERSION="1.0.0.0"
MESSAGE="New Version 1.0.0.0"

# docker build -t $DOCKER_REPO:dev .
# docker push $DOCKER_REPO:dev
# echo "[`date "+%Y-%m-%d %H:%M:%S"`] dev => {$MESSAGE}" >> ImageInfo.txt

# docker build -t $DOCKER_REPO:$VERSION .
# docker push $DOCKER_REPO:$VERSION
# echo "[`date "+%Y-%m-%d %H:%M:%S"`] $VERSION => {$MESSAGE}" >> ImageInfo.txt


docker rmi -f $(docker images | grep $DOCKER_REPO | awk '{print $3}')
docker image prune -f