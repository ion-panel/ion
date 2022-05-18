package ion

import (
	"fmt"
	"html/template"
	"net/http"
)

type AccountController struct{}

// Index serves the index of the ION panel, containing the login form.
func (a *AccountController) Index(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("www/login/index.html")

	err := tmp.Execute(response, nil)
	if err != nil {
		return
	}
}

// Login serves the login form, complete with session handling for a single admin user.
func (a *AccountController) Login(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		return
	}

	config := &Config{}
	confPwd := config.RetrieveAdminPassword()
	confUser := config.RetrieveAdminUsername()

	username := request.Form.Get("username")
	password := request.Form.Get("password")

	if username == confUser && password == confPwd &&
		len(confUser) != 0 &&
		len(confPwd) != 0 {
		fmt.Println("LOG: Login accepted: { username: " + confUser + "} " + "{ password: " + confPwd + " }")

		session, _ := store.Get(request, "ion_auth_cookie")
		session.Values["is_authenticated"] = true
		err := session.Save(request, response)
		if err != nil {
			return
		}

		http.Redirect(response, request, "/dashboard/index/", http.StatusSeeOther)
	} else {
		fmt.Println("LOG: Login denied")

		data := map[string]interface{}{
			"err": "invalid",
		}
		tmp, _ := template.ParseFiles("www/login/index.html")
		err := tmp.Execute(response, data)
		if err != nil {
			return
		}
	}
}

func (a *AccountController) Logout(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "ion_auth_cookie")

	session.Values["is_authenticated"] = false
	err := session.Save(request, response)
	if err != nil {
		return
	}

	http.Redirect(response, request, "/admin/login", http.StatusOK)
}
