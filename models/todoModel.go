package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	Updated   int       `json:"updated"`
	Created   int       `json:"created"`
}

type Todos []Todo

var currentId int

var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
	RepoCreateTodo(Todo{Name: "Host meetup Bran Van 3000"})
	RepoCreateTodo(Todo{Name: "Host meetup Bran Van 4000"})

}

func RepoGetTodos() Todos {
	db, e := myDb.setup()
	if e != nil {
		fmt.Print(e)
	}
	rows, err := db.Query("select id, name, completed from todos")
	if err != nil {
		fmt.Print(err)
	}

	//	var results = make([]Todo, 1000)

	var results Todos

	if err != nil {
		fmt.Print(err)
	}
	i := 0
	for rows.Next() {
		var name string
		var id int
		var completed bool
		err = rows.Scan(&id, &name, &completed)
		log.Print(rows.Columns())
		log.Print(id)
		log.Print(name)
		//		todo := &Todo{Id: id, Name: name, Completed: completed}
		todo := Todo{Id: id, Name: name, Completed: completed}
		//		b, err := json.Marshal(todo)
		//		if err != nil {
		//			fmt.Println(err)
		//			return
		//		}
		//		results[i] = fmt.Sprintf("%s", string(b))
		//		results[i] := todo
		results = append(results, todo)
		i++
	}
	//	result = result[:i]

	log.Println(results)
	log.Println(todos)

	return results
}

func RepoFindTodo(id int) Todo {
	//	for _, t := range todos {
	//		if t.Id == id {
	//			return t
	//		}
	//	}
	db, e := sql.Open("mysql", "root:@tcp(localhost:3306)/accountancyApp")
	if e != nil {
		fmt.Print(e)
	}
	var todo Todo
	var name string
	//	var id int
	var completed bool
	err := db.QueryRow("select name, completed from todos where id = ?", id).Scan(&name, &completed)
	if err != nil {
		fmt.Print(err)
	}
	todo.Id = id
	todo.Name = name
	todo.Completed = completed
	log.Println(todo)
	return todo

	// return empty Todo if not found
	//	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
