package main

import "fmt"

func main() {
	user := make(map[string]map[string]string, 10)
	modifyUser(user, "tom")
	for _, v := range user {
		fmt.Println(v)
	}

}

func modifyUser(user map[string]map[string]string, username string) {
	if user[username] != nil {
		user[username]["psw"] = "888888"
	} else {
		user[username] = make(map[string]string, 2)
		user[username]["psw"] = "888888"
	}
}