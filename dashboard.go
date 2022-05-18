package ion

import (
	"html/template"
	"net/http"
)

type Dashboard struct{}

func (d *Dashboard) Index(response http.ResponseWriter, request *http.Request) {

	session, _ := store.Get(request, "ion_auth_cookie")

	if c, ok := session.Values["is_authenticated"].(bool); !ok || !c {
		http.Redirect(response, request, "/admin/login", http.StatusForbidden)

		return
	}

	tmp, _ := template.ParseFiles("web-fragments/dashboard/index.html")
	err := tmp.Execute(response, nil)
	if err != nil {
		return
	}
}

func (d *Dashboard) Logout(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "ion_auth_cookie")

	session.Values["is_authenticated"] = false
	err := session.Save(request, response)
	if err != nil {
		return
	}

	http.Redirect(response, request, "/admin/login", http.StatusOK)
}
