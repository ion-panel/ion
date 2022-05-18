package ion

import (
	"github.com/gorilla/sessions"
)

var key = make([]byte, 64)
var store = sessions.NewCookieStore(key)
