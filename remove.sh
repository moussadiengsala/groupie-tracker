docker image prune -a
docker stop $(docker ps -a -q)
docker image prune -a
docker rm -f $(docker ps -a -q)
clear
docker images
docker rmi -f $(docker images -aq)