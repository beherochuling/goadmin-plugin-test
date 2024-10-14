package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/beesoft/goadmin-plugin-test/guard"
	"html/template"
)

func (h *Handler) Example(ctx *context.Context) {
	var param = guard.GetExampleParam(ctx)
	h.HTML(ctx, types.Panel{
		Title:       "Example",
		Description: "example",
		Content:     template.HTML(param.Param),
	})
}
