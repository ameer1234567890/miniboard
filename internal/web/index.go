package web

import (
	"net/http"
	"strconv"

	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/yaml"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig

	reload := r.URL.Query().Get("reload")
	if reload == "yes" {
		AllLinks = yaml.Read(AppConfig.YamlPath)
		assignAllIDs() // assign-IDs.go

		close(AppConfig.Quit)
		AppConfig.Quit = make(chan bool)

		go scanPorts(AppConfig.Quit) // scan.go

		http.Redirect(w, r, "/", 302)
	}

	tabStr := r.URL.Query().Get("tab")
	var tab int
	if tabStr == "" {
		tab = 0
	} else {
		tab, _ = strconv.Atoi(tabStr)
	}

	guiData.CurrentTab = AllLinks.Tabs[tab].Name
	guiData.Links = AllLinks

	guiData.Panels = make(map[int]models.Panel)
	for i, name := range AllLinks.Tabs[tab].Panels {
		guiData.Panels[i] = AllLinks.Panels[name]
	}

	execTemplate(w, "index", guiData)
}
