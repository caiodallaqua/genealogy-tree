// Models in responses.go are not used anywhere in the code
// except for swagger auto-generated docs.
// DRY only if it helps to make the documentation clearer.
// Explicit is better than implicit, specially for docs.
package models

import "time"

// ----------------------      BASE      ----------------------

type PersonData struct {
	ID    uint32    `json:"id" example:"42"`
	Name  string    `json:"name" example:"Aristeu"`
	Birth time.Time `json:"birth" example:"2020-12-09T16:09:53+00:00"`
}

type RelationData struct {
	ParentID uint32 `json:"parent_id" example:"2"`
	ChildID  uint32 `json:"child_id" example:"4"`
}

// ----------------------      GET      ----------------------

type GetStatusRes struct {
	Data struct {
		Result string `json:"result" example:"ok"`
	}
}

type GetPersonRes struct {
	Data PersonData
}

type GetAscendantsRes struct {
	Data struct {
		Relationships []RelationData
		Ascendants    []PersonData
	}
}

type GetAscendantsAndDescendantsRes struct {
	Data struct {
		Relationships            []RelationData
		AscendantsAndDescendants []PersonData
	}
}

type GetAscendantsAndChildrenRes struct {
	Data struct {
		Relationships         []RelationData
		AscendantsAndChildren []PersonData
	}
}

// ----------------------      POST      ----------------------

type PostPersonRes struct {
	Data struct {
		ID uint32 `json:"id" example:"42"`
	}
}

type PostParentRelationshipRes struct {
	Data struct {
		Result string `json:"result" example:"ok"`
	}
}

// ----------------------      DELETE      ----------------------

type DelPersonRes struct {
	Data struct {
		Result string `json:"result" example:"ok"`
	}
}

type DelParentRelationshipRes struct {
	Data struct {
		Result string `json:"result" example:"ok"`
	}
}
