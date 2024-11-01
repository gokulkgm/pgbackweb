package layout

import (
	"github.com/eduardolat/pgbackweb/internal/view/reqctx"
	"github.com/eduardolat/pgbackweb/internal/view/web/component"
	"github.com/eduardolat/pgbackweb/internal/view/web/htmx"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type DashboardParams struct {
	Title string
	Body  []gomponents.Node
}

func Dashboard(reqCtx reqctx.Ctx, params DashboardParams) gomponents.Node {
	title := "PG Back Web"
	if params.Title != "" {
		title = params.Title + " - " + title
	}

	if reqCtx.IsHTMXBoosted {
		body := append(params.Body, html.TitleEl(gomponents.Text(title)))
		return component.RenderableGroup(body)
	}

	return components.HTML5(components.HTML5Props{
		Language: "en",
		Title:    title,
		Head: []gomponents.Node{
			head(),
		},
		Body: []gomponents.Node{
			htmx.HxIndicator("#header-indicator"),
			components.Classes{
				"w-screen h-screen bg-base-200":      true,
				"flex justify-start overflow-hidden": true,
				"mobile:flex-col mobile:overflow-auto": true, // Add media queries for mobile screens
			},
			dashboardAside(),
			html.Div(
				html.Class("flex-grow overflow-y-auto mobile:overflow-visible"), // Adjust layout to be more flexible and responsive
				dashboardHeader(),
				html.Main(
					html.ID("dashboard-main"),
					html.Class("p-4 mobile:p-2"), // Adjust layout to be more flexible and responsive
					gomponents.Group(params.Body),
				),
			),
		},
	})
}
