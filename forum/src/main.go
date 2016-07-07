package main

import (
	"github.com/kataras/iris"
	"html/template"
)

func main() {
	db := dbInit()

	a := iris.New()

	a.Config.Render.Template.Layout = "main.html"
	a.Config.Render.Template.IsDevelopment = true
	a.Config.Render.Template.HTMLTemplate.Funcs = template.FuncMap{}
	a.Static("/public", "./public", 1)

	a.Get("/", func(c *iris.Context) {
		sections := []Section{}
		db.Select(&sections, "SELECT * FROM 'sections' ORDER BY 'ID'")
		c.Render("index.html", map[string]interface{}{"sections": sections})
	})

	a.Get("/section/:section", func(c *iris.Context) {
		section, _ := c.ParamInt("section")

		themes := []Theme{}
		db.Select(&themes, "SELECT * FROM 'themes' WHERE SECTION_ID=? ORDER BY 'ID'", section)
		c.Render("section.html", map[string]interface{}{"themes": themes})
	})

	a.Get("/theme/:theme", func(c *iris.Context) {
		theme, _ := c.ParamInt("theme")

		messages := []Message{}
		db.Select(&messages, "SELECT * FROM 'messages' WHERE THEME_ID=? ORDER BY 'ID'", theme)
		c.Render("theme.html", map[string]interface{}{"messages": messages})
	})

	a.Listen(":8080")
}
