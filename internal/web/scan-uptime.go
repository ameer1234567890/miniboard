package web

import (
	// "log"
	"fmt"
	"sync"
	"time"

	// "github.com/aceberg/miniboard/internal/check"
	"github.com/aceberg/miniboard/internal/models"
	"github.com/aceberg/miniboard/internal/notify"
)

var (
	// RetriesSyncMap - count retries to send notifications
	RetriesSyncMap sync.Map
)

func appendUptimeMon(panelName string, host models.Host, notif bool) {
	mon := models.MonData{}
	mon.Panel = panelName
	mon.Host = host.Name
	mon.Addr = host.Addr
	mon.Port = host.Port
	mon.Date = time.Now().Format("2006-01-02 15:04:05")
	mon.State = host.State
	mon.Notify = notif
	UptimeMon = append(UptimeMon, mon)
}

func scanUptime(panelName string, host models.Host, oldState bool) {
	var retries int
	var notifEnabled, notif bool

	_, exists := AllLinks.Uptime.Panels[panelName]

	if AllLinks.Uptime.Enabled && exists {
		if len(AllLinks.Uptime.Panels[panelName].Notify) > 0 {
			retriesAny, ok := RetriesSyncMap.LoadOrStore(panelName+host.Name, 0)
			if ok {
				retries = retriesAny.(int)
			}
			notifEnabled = true
		}

		if oldState != host.State {

			if notifEnabled && host.State && (retries > AllLinks.Uptime.Panels[panelName].Retries) {
				notify.Notify(panelName, host.Name, "is up", AllLinks.Uptime)
				RetriesSyncMap.Store(panelName+host.Name, 0)
				notif = true
			}
			appendUptimeMon(panelName, host, notif)
		}
		if notifEnabled && !host.State {
			if retries == AllLinks.Uptime.Panels[panelName].Retries {
				msg := fmt.Sprintf("is down (retries: %d)", AllLinks.Uptime.Panels[panelName].Retries)
				notify.Notify(panelName, host.Name, msg, AllLinks.Uptime)

				appendUptimeMon(panelName, host, true)
			}
			RetriesSyncMap.Store(panelName+host.Name, retries+1)
		}
	}
}
