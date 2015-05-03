sample POST query

curl -H "Content-Type: application/json" -d '{"name":"New Todo 123"}' http://localhost:8080/todos


sample GET query

curl -H "Content-Type: application/json" -g http://localhost:8080/todos


sample GET by id

curl -H "Content-Type: application/json" -g http://localhost:8080/todos/12