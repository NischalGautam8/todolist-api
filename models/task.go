package models

type Task struct {
	Id         int    `json:"id,omitempty"`
	Title      string `json:"title"`
	Desce      string `json:"desce"`
	Completed  bool   `json:"completed"`
	Created_at string `json:"created_at,omitempty"`
}

//docker pull postgres
//docker run --name pg-container -e POSTGRES_PASSWORD=nischal -p 5432:5432 -d postgres
// docker exec -ti pg-container psql -U postgres
//\c gopgtest to connect to gopgtest db
// CREATE TABLE task(id INT,desce VARCHAR(255),completed boolean,create_at VARCHAR(255));
// CREATE TABLE "user"(id INT,username VARCHAR(255),password VARCHAR(255));
//\dt to display tables
