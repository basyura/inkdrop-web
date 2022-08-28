package controllers

import (
	"fmt"
	"inkdrop-server/app/inkdrop/proxy"

	"github.com/revel/revel"
)

type NoteController struct {
	*revel.Controller
}

func (c NoteController) Index() revel.Result {

	b, err := proxy.Get("/notes")
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	return c.RenderJSON(string(b))
}

func (c NoteController) Show(id string) revel.Result {
	// note:qPIR-2cx5
	b, err := proxy.Get("/note:" + id)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	return c.RenderJSON(string(b))

}
