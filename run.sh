sudo docker image build -f Dockerfile -t asciiweb . &
pid=$!
wait $pid
echo "-----"
sudo docker images &
pid=$!
wait $pid
echo "-----"
sudo docker container run -p 8080:8080 --detach --name web asciiweb &
pid=$!
wait $pid
echo "-----"
sudo docker ps -a &
pid=$!
wait $pid
echo "-----"
#sudo docker exec -it web /bin/bash