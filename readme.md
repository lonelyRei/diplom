#### поднять контейнер: 
docker run --name=db -e POSTGRES_PASSWORD='qwe
rty' -p 5436:5432 -d --rm postgres

#### проверить состояние:
docker ps

#### создать миграцию:
migrate create -ext sql -dir .\scema\ -seq ini
t

#### выполнить миграцию:
migrate -path ./scema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up


#### очистить миграцию
docker exec -it e245c75aafca /bin/bash
psql -U postgres - запускает в консоль
\d - чекнуть все таблицы
select * from schema_migrations;
update schema_migrations set version='000001',dirty=false;

#### подцепиться к базе:
docker exec -it e245c75aafca /bin/bash
psql -U postgres - запускает в консоль