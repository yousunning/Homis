package main

import "fmt"

var users []User // 배열 user형으로 userd에게 선언

func main() {
	//jsonUsers := "hello world"
	users = []User{ // users 변수에  user 배열 할당
		{
			Name:   "name1",
			Age:    1,
			Gender: "M",
		},
		{
			Name:   "name2",
			Age:    2,
			Gender: "F",
		},
		{
			Name:   "name3",
			Age:    3,
			Gender: "M",
		},
	}
	Delete("name2") // name2 를 지울거야
}

type User struct {
	Name   string
	Age    int
	Gender string
}

func Delete(name string) []User {
	var result []User
	for i, user := range users {
		if user.Name == name {
			continue
		}
		result = append(result, users[i:][0])
	}
	fmt.Print(result)
	return result
}
