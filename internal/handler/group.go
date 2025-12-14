package handler

import (
	"encoding/json"
	"net/http"

	"RazdelyCheck/internal/dto"
	"RazdelyCheck/internal/service"
	"RazdelyCheck/internal/util"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type GroupHandler struct {
	service *service.GroupService
}

func NewGroupHandler(service *service.GroupService) *GroupHandler {
	return &GroupHandler{service: service}
}

// POST /groups
func (h *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var g dto.Group
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	if err := h.service.Create(&g); err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusCreated, g)
}

// GET /groups/{id}
func (h *GroupHandler) GetGroup(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	groupID, err := uuid.Parse(idStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid groupID"})
		return
	}

	group, err := h.service.GetByID(groupID)
	if err != nil {
		util.RespondJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusOK, group)
}

// GET /groups
func (h *GroupHandler) ListGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := h.service.List()
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusOK, groups)
}

// PUT /groups/{id}
func (h *GroupHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	groupID, err := uuid.Parse(idStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid groupID"})
		return
	}

	var g dto.Group
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}
	g.ID = groupID

	if err := h.service.Update(&g); err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusOK, g)
}

// DELETE /groups/{id}
func (h *GroupHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	groupID, err := uuid.Parse(idStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid groupID"})
		return
	}

	if err := h.service.Delete(groupID); err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

// POST /groups/{id}/users
func (h *GroupHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	groupID, err := uuid.Parse(idStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid groupID"})
		return
	}

	var body struct {
		UserID uuid.UUID `json:"userId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	if err := h.service.AddUserToGroup(body.UserID, groupID); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// DELETE /groups/{id}/users/{userID}
func (h *GroupHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid groupID"})
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid userID"})
		return
	}

	if err := h.service.RemoveUserFromGroup(userID, groupID); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// GET /groups/user/{userID}
func (h *GroupHandler) ListGroupsByUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid userID"})
		return
	}

	groups, err := h.service.ListGroupsByUser(userID)
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusOK, groups)
}

// GET /groups/{id}/users
func (h *GroupHandler) ListUsersByGroup(w http.ResponseWriter, r *http.Request) {
	groupIDStr := chi.URLParam(r, "id")
	groupID, err := uuid.Parse(groupIDStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid groupID"})
		return
	}

	users, err := h.service.ListUsersByGroup(groupID)
	if err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusOK, users)
}
