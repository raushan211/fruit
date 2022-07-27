package main

import "fruit/server"

func main() {

	createDBConnection()
	//defer DB.Close()
	//Data = make(map[string]User)
	server.Server()
}
