# idletime

## build & run
build: `docker-compose build --no-cache`<br>
run: `docker-compose up`<br>
both: `docker-compose up --build`

## access docker terminal
see all containers: `docker ps`<br>
start shell inside container: `docker exec -it container_name bash`

## postgres access for checking
log into postgres: `psql dbname -U user`<br>
(only works on our system as our db name is dbname, see docker-compose.yml)<br>
see all tables: `\dt`
