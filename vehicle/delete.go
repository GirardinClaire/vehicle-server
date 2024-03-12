package vehicle

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GirardinClaire/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")

	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		fmt.Println("An error has occured", err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	test, err := d.store.Vehicle().Delete(r.Context(), id)
	if err != nil {
		fmt.Println("An error has occured", err)
		return
	}

	if test {
		rw.WriteHeader(http.StatusNoContent)
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}
}
