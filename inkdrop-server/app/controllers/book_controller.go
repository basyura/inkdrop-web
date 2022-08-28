package controllers

import (
	"fmt"
	"inkdrop-server/app/inkdrop"

	"github.com/revel/revel"
)

type BookController struct {
	*revel.Controller
}

func (c BookController) Index() revel.Result {
	b, err := inkdrop.Get("/books")
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	return c.RenderText(string(b))
}
