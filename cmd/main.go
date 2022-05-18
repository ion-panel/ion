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
	fmt.Println("Instance running on:  :2000\n")

	//http.HandleFunc("/admin/index", acc.Index)
	http.HandleFunc("/admin/login", acc.Login)

	http.HandleFunc("/dashboard/index/", dashboard.Index)
	http.HandleFunc("/dashboard/logout/", dashboard.Logout)

	err := http.ListenAndServe(":2000", nil)
	if err != nil {
		return
	}

}
