package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleCreateForm() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		html := `
<h1>What is the domain for the new site?</h1>
<form action="create-site" method="post">
		From:
		<input type="text" name="skeleton" value="something.com" /><br />
		New Domain:
		<input type="text" name="domain" /><br /><br />
		<input type="submit" value="Create" />
</form>
		`
		w.Write([]byte(html))
	}
}
