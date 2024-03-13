package vehicle

import (
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
		http.Error(rw, "Failed to parse ID", http.StatusInternalServerError)
		return
	}

	test, err := d.store.Vehicle().Delete(r.Context(), id)
	if err != nil {
		http.Error(rw, "Failed to delete vehicle for the id", http.StatusInternalServerError)
		return
	}

	if test {
		rw.WriteHeader(http.StatusNoContent)
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}
}
