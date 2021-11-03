#!/bin/bash
CONTAINER="X-X-X"
DOCKER_REPO="iiicondor/$CONTAINER"
VERSION="X.X.X.X"
MESSAGE="New Version X.X.X.X"

# docker build -t $DOCKER_REPO:dev .
# docker push $DOCKER_REPO:dev
# echo "[`date "+%Y-%m-%d %H:%M:%S"`] dev => {$MESSAGE}" >> ImageInfo.txt

docker build -t $DOCKER_REPO:$VERSION .
docker push $DOCKER_REPO:$VERSION
echo "[`date "+%Y-%m-%d %H:%M:%S"`] $VERSION => {$MESSAGE}" >> ImageInfo.txt


docker rmi -f $(docker images | grep $DOCKER_REPO | awk '{print $3}')
docker image prune -f