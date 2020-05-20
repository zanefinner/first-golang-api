package base

import (
	"fmt"
	"html/template"
	"net/http"
)

//IndexHandle Resolves the site's root
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Website Root"))
	t, err := template.New("foo").Parse(`<!doctype HTML>
	<html>
		<head>
			<title>Hello World</title>
		</head>
		<body>
			<p>Hello, {{.}}</p>
		</body>
	</html>
		`)
	err = t.Execute(w, template.HTML(`<b>Zane</b>`))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Endpoint Hit: Website Root")
}
