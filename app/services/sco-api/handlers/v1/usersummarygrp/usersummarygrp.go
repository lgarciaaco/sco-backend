// Package usersummarygrp maintains the group of handlers for user summary access.
package usersummarygrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sco1237896/sco-backend/business/core/usersummary"
	"github.com/sco1237896/sco-backend/business/web/v1/paging"
	"github.com/sco1237896/sco-backend/foundation/web"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	summary *usersummary.Core
}

// New constructs a handlers for route access.
func New(summary *usersummary.Core) *Handlers {
	return &Handlers{
		summary: summary,
	}
}

// Query returns a list of user summary data with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	page, err := paging.ParseRequest(r)
	if err != nil {
		return err
	}

	filter, err := parseFilter(r)
	if err != nil {
		return err
	}

	orderBy, err := parseOrder(r)
	if err != nil {
		return err
	}

	smms, err := h.summary.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	total, err := h.summary.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(toAppUsersSummary(smms), total, page.Number, page.RowsPerPage), http.StatusOK)
}
