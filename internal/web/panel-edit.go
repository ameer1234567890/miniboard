package web

import (
	// "log"
	"net/http"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
)

func panelEditHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.CurrentTab = "Edit Panel"
	guiData.Links.Tabs = AllLinks.Tabs

	edit := r.URL.Query().Get("edit")
	del := r.URL.Query().Get("del")

	if edit != "" {
		guiData.Links.Panels = make(map[string]models.Panel)
		guiData.Links.Panels[edit] = AllLinks.Panels[edit]

		execTemplate(w, "panel-edit", guiData)
	} else if del != "" {
		delete(AllLinks.Panels, del)
		// delete from tabs?
		yaml.Write(AppConfig.YamlPath, AllLinks)

		http.Redirect(w, r, "/panels/", 302)
	}
}