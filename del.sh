sudo docker stop web &
pid=$!
wait $pid
echo "-----"
sudo docker rm web &
pid=$!
wait $pid
sudo docker rmi asciiweb &
pid=$!
wait $pid
echo "-----"
sudo docker images &
pid=$!
wait $pid
echo "-----"
sudo docker ps -a 
echo "-----"
echo "tests terminés"