package storage

import (
	"github.com/GirardinClaire/vehicle-server/storage/vehiclestore"
)

type Store interface {
	Vehicle() vehiclestore.Store
}
