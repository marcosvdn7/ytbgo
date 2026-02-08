package ytbgo

import "strings"

func GetContext(clientName string) ClientContext {
	for _, clientContext := range config.Clients {
		if strings.EqualFold(clientContext.ClientName, clientName) {
			return clientContext
		}
	}
	return config.Clients[0]
}
