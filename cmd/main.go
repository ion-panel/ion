package main

import (
	"fmt"
	"ion"
	"net/http"
)

var bootStr = `
  _____            
 |_   _|           
   | |  ___  _ __  
   | | / _ \| '_ \ 
  _| || (_) | | | |
 |_____\___/|_| |_|
`

func main() {
	conf := &ion.Config{}

	if !conf.Exists() {
		fmt.Println("ion-config.json file does not exist. Creating...")
		conf.Create()

		fmt.Println("Created config file. Please restart Ion.")
		return
	}

	acc := &ion.AccountController{}
	dashboard := &ion.Dashboard{}

	fmt.Println(bootStr)

	ss := http.FileServer(http.Dir("./static"))
	http.Handle("/", ss)
	fmt.Println("CSS Handeling server started: ")

	// endpoints
	http.HandleFunc("/admin/login", acc.Login)
	http.HandleFunc("/dashboard/logout/", acc.Logout)
	http.HandleFunc("/dashboard/index/", dashboard.Index)

	fmt.Println("Handlers initialized.")

	fmt.Println("Instance running on: " + conf.GetHost())

	err := http.ListenAndServe(conf.GetHost(), nil)
	if err != nil {
		return
	}

}
