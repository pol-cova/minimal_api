package mapi

import (
	"html/template"
)

func (c *Context) RenderTemplate(templateName string, data interface{}) {
	tmpl, err := template.New(templateName).ParseFiles(templateName)
	if err != nil {
		c.InternalServerError("Failed to render template")
		return
	}
	c.SetHeader("Content-Type", "text/html")
	tmpl.Execute(c, data)
}
