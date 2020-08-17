package server

import "testing"

func TestWeCanParseDataFromAnHTMLString(t *testing.T) {
	basicTemplate := `nothing fancy`
	output, data := parseVarsFromView(basicTemplate)

	if output != basicTemplate {
		t.Error("A basic template should match the output.")
	}
	if len(data) != 0 {
		t.Error("A basic template should not have generated any view data.")
	}

	fancyTemplate := `=====
Foo: bar
=====
stuff!`
	output, data = parseVarsFromView(fancyTemplate)
	if output != "stuff!" {
		t.Errorf("When parsing a template with vars, we should ignore all the lines surrounded by the delimiter. \nExpected: %v\n  Actual: %v", "stuff!", output)
	}
	if data["Foo"] != "bar" {
		t.Errorf("When parsing a template with vars, the top area of the template should parse into key value pairs of type string. \nData: %v", data)
	}
}
