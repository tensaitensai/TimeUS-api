docker/run: 
	docker-compose up -d

docker/down: 
	docker-compose down
	@rm -rf db/data

docker/mysql:
	mysql -h 127.0.0.1 -P 3306 -utestuser -ptestpass -P 3306 
	
	