package hallpers

import (
	"html/template"
	"os"
)

func CreateHTML (Data any,PathStartFile ,PathEndile string)string{
	temp, _ := template.ParseFiles(PathStartFile)

	f, _ := os.Create(PathEndile)
	temp.Execute(f, Data)

	return PathEndile
}