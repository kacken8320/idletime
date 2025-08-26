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

## to-do
only server:<br>
[] http for early web-functionalities (testing purposes)<br>
[] dummy users (user-table, userIds in category, NOT activity) (further testing)<br>
[] proper users (sign-in, no security yet)<br>
[] tlr stuff<br>
[] think of more to-do stuff<br>