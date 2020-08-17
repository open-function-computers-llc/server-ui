package server

import (
	"html/template"
	"net/http"
	"strings"
)

func (s *Server) renderView(file string, w http.ResponseWriter, data *map[string]interface{}) {
	v, viewData, err := s.view(file)
	t, _ := template.New("foo").Parse(v)

	var d map[string]interface{}
	if data == nil {
		d = make(map[string]interface{})
	} else {
		d = *data
	}

	// append values pulled from the view file
	for k, v := range viewData {
		d[k] = v
	}

	// global things that are available for every view
	d["RoutePrefix"] = s.routePrefix

	// append errors to data
	if err != nil {
		d["StatusCode"] = 500
	}

	t.Execute(w, d)
}

func (s *Server) view(file string) (string, map[string]string, error) {
	header, _ := s.assets.String("partials/head.tpl")
	footer, _ := s.assets.String("partials/foot.tpl")
	viewData := make(map[string]string)

	templateString, err := s.assets.String(file + ".tpl")
	if err != nil {
		s.LogError("Requested view file " + file + ".tpl was not found.")
		templateString, _ = s.assets.String("error.tpl")
		viewData["Pagetitle"] = "Error"
		viewData["BodyClasses"] = "error-page"
		return header + templateString + footer, viewData, err
	}
	templateString, viewData = parseVarsFromView(templateString)
	return header + templateString + strings.TrimSpace(footer), viewData, nil
}
