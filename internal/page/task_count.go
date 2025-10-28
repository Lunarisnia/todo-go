package page

import (
	"html/template"
	"net/http"
)

type TaskCountPage struct {
	templ string
}

func NewTaskCountPage() TaskCountPage {
	taskCountPage := TaskCountPage{
		templ: `
<!DOCTYPE html>
<html lang="en">
	<head>
		<title>ToDo App</title>
	</head>
	<body>
		<h1 style="color: red;">Active Task: {{.Count}}!</h1>
	</body>
</html>
`,
	}
	return taskCountPage
}

type TaskCountPageContent struct {
	Count int
}

func (t TaskCountPage) Execute(w http.ResponseWriter, content TaskCountPageContent) error {
	tmp, err := template.New("Webpage").Parse(t.templ)
	if err != nil {
		return err
	}

	if err := tmp.Execute(w, content); err != nil {
		return err
	}

	return nil
}
