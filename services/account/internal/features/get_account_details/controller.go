//nolint:godot //swagger docs
package get_account_details

import (
	"net/http"

	"github.com/ssengalanto/runic/pkg/fn"
	"github.com/ssengalanto/runic/pkg/http/json"
	"github.com/ssengalanto/runic/pkg/interfaces"
	"github.com/ssengalanto/runic/pkg/log"
)

// Controller represents a controller that handles business logic
// and serves as an entry point for handling incoming requests.
type Controller struct {
	slog     interfaces.Logger
	mediator interfaces.Mediator
}

// NewController creates a new instance of the Controller.
func NewController(slog interfaces.Logger, mediator interfaces.Mediator) *Controller {
	return &Controller{slog: slog, mediator: mediator}
}

// Handle retrieves the account details by ID.
// @Tags account
// @Summary Get Account Details
// @Description Retrieves the account details for a specific account ID.
// This API endpoint fetches account information, including owner details.
// @Accept json
// @Produce json
// @Param id path string true "Account ID" example("0b6ecded-fa9d-4b39-a309-9ef501de15f4")
// @Success 200 {object} AccountUserResponse "Returns the account details."
// @Failure 404 {object} HTTPError "The specified record does not exist."
// @Failure 500 {object} HTTPError "An error occurred while processing the request."
// @Router /api/account/{id} [get]
func (d *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := fn.ParseUUIDFromURLParam("id", w, r)
	if err != nil {
		d.slog.Error("invalid uuid", log.Field("error", err))
		return
	}

	q := NewQuery(id)

	result, err := d.mediator.Send(ctx, q)
	if err != nil {
		json.Error(w, err)
		return
	}

	res := result.(Response) //nolint:errcheck //intentional panic
	json.Success(w, http.StatusOK, res)
}
