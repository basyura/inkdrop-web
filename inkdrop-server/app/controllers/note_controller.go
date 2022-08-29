package controllers

import (
	"fmt"
	"inkdrop-server/app/inkdrop"

	"github.com/revel/revel"
)

type NoteController struct {
	*revel.Controller
}

func (c NoteController) Index() revel.Result {

	b, err := inkdrop.Get("notes")
	if err != nil {
		return c.RenderJSON(err)
	}

	return c.RenderText(string(b))
}

func (c NoteController) Show(id string) revel.Result {
	// note:qPIR-2cx5
	b, err := inkdrop.Get("note:" + id)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	return c.RenderText(string(b))
}
