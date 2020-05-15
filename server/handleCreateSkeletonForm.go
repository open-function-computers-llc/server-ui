package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) handleCreateSkeletonForm() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// TODO: generate the list of sites via DB or something
		html := `
<h1>What is the domain you want to turn into a skeleton?</h1>
<form action="create-skeleton" method="post">
		Site:
		<select name="domain">
			<option value="newcoolthing.com">newcoolthing.com</option>
			<option value="othercoolthing.com">othercoolthing.com</option>
		</select><br /><br />
		<input type="submit" value="Create Skeleton" />
</form>
		`
		w.Write([]byte(html))
	}
}
