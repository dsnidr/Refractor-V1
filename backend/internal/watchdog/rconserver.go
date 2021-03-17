package watchdog

import (
	"github.com/sniddunc/refractor/pkg/log"
	"github.com/sniddunc/refractor/refractor"
	"net"
	"time"
)

// StartRCONServerWatchdog starts a watchdog for RCON clients. If a server exists and does not have an RCON client, this
// watchdog will create one for it. If a server goes offline, this watchdog is also responsible for reconnecting.
func StartRCONServerWatchdog(rconService refractor.RCONService, serverService refractor.ServerService, log log.Logger) {
	for {
		// Run every 15 seconds
		time.Sleep(time.Second * 15)

		rconClients := rconService.GetClients()
		allServerData, _ := serverService.GetAllServerData()

		if len(rconClients) != len(allServerData) {
			for _, serverData := range allServerData {
				client := rconClients[serverData.ServerID]

				// Check if an RCON client exists for this server. If one does not, create one
				if client == nil {
					server, _ := serverService.GetServerByID(serverData.ServerID)
					if server == nil {
						log.Warn("Watchdog could not create a new RCON client for server ID: %d. Server could not fetched by id.", serverData.ServerID)
						continue
					}

					if err := rconService.CreateClient(server); err != nil {
						switch errType := err.(type) {
						case *net.OpError:
							// If this error is not a dial error, we want to log it. If it is a dial error, we just assume
							// that the server is offline so there is no need to spam the logs with this info.
							if errType.Op != "dial" {
								log.Warn("Watchdog RCON client connection error: %v", err)
							}
							continue
						default:
							log.Error("Watchdog could not create a new RCON client for server ID: %d. Error: %v", serverData.ServerID, err)
							continue
						}
					}
				}
			}
		}
	}
}
