echo "### Building image"
docker build -t groupie-tracker .
echo
echo "### Running image"
docker run -d -p 8000:8000 --name web groupie-tracker
echo
echo "### Images list"
docker images
echo
echo "### Containers list"
docker container ls