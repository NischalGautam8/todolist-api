1. Pull the official postgres docker image `docker pull postgres`
2. Run it `docker run --name pg-container -e POSTGRES_PASSWORD=nischal -p 5432:5432 -d postgres`
3. `docker exec -ti pg-container psql -U postgres`
4. connect to gopgtest db `\c gopgtest`
5. Run the sql command to create two tables user and task `CREATE TABLE task(id INT,desce VARCHAR(255),completed boolean,create_at VARCHAR(255)); CREATE TABLE "user"(id INT,username VARCHAR(255),password VARCHAR(255));`
6. Run the main file `go run main.go` from root directory
7. try accessing /register / login /tasks
