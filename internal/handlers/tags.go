package handlers

import (
	"net/http"

	db "galihwicaksono90/musikmarching-be/internal/storage/persistence"
)

type Tags struct {
	Instruments []db.Instrument `json:"instruments"`
	Categories  []db.Category   `json:"categories"`
	Allocations []db.Allocation `json:"allocations"`
}

func (h *Handler) HandleGetScoreTags(w http.ResponseWriter, r *http.Request) {
	categories, _ := h.category.GetAll()
	instruments, _ := h.instrument.GetAll()
	allocations, _ := h.allocation.GetAll()

	var tags Tags

	tags.Instruments = instruments
	tags.Categories = categories
	tags.Allocations = allocations

	h.handleResponse(w, http.StatusOK, http.StatusText(http.StatusOK), tags)
}
