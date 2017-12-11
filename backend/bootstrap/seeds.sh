docker-compose exec go  bash -c "soda create -e development"
docker-compose exec go  bash -c "soda migrate"
docker-compose exec go  bash -c "cd bootstrap && go build && ./bootstrap"