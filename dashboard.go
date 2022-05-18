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

	tmp, _ := template.ParseFiles("www/dashboard/index.html")
	err := tmp.Execute(response, nil)
	if err != nil {
		return
	}
}
