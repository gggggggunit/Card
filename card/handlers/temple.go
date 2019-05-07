package handlers

import "html/template"

var tpl *template.Template

func init() {

	var err error //где находится штмл
	tpl, err = template.ParseFiles(
		"templates/index.html",
		"templates/room.html",

	)

	if err != nil {
		panic(err)
	}
}
