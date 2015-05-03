sample POST query

`curl -H "Content-Type: application/json" -d '{"name":"New Todo 123"}' http://localhost:8080/todos`


sample GET query

curl -H "Content-Type: application/json" -g http://localhost:8080/todos


sample GET by id

`curl -H "Content-Type: application/json" -g http://localhost:8080/todos/12`


## Assumptions

- The business is operating in one country and selling internationaly
- For the goods or services we sell there is only one tax rate in each country
- For each financial year exchange rates in our system are fixed
- An account can only exist in one country
- An accountholder manages the accounts for one country
- Sale and handover are independent