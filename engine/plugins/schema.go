package plugins

import (
	"github.com/AstroSynapseAI/app-service/sdk/crud/database"
)

type Plugins interface {
	LoadConfig(db *database.Database) error
	OpenConnection(db *database.Database) error
}

type Plugin interface {
	OpenConnection(db *database.Database)
}
