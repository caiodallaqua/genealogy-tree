package rest

import (
	"errors"
	"genealogy-tree/internal/models"
)

func checkPostParentRelationshipPayload(payload models.PostParentRelationship) error {
	if payload.ChildID == payload.ParentID {
		return errors.New("Parent and child cannot have the same identifier.")
	}

	return nil
}
