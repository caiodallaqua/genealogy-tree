package models

import "time"

type PersonData struct {
	ID    uint32    `json:"id" example:"42"`
	Name  string    `json:"name" example:"Aristeu"`
	Birth time.Time `json:"birth" example:"2020-12-09T16:09:53+00:00"`
}

type RelationData struct {
	ParentID uint32 `json:"parent_id" example:"2"`
	ChildID  uint32 `json:"child_id" example:"4"`
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

type PostPersonRes struct {
	Data struct {
		ID uint32 `json:"id" example:"42"`
	}
}

type PostParentRelationshipRes struct {
	Data struct {
		string `example:"ok"`
	}
}

type DelPersonRes struct {
	Data struct {
		string `example:"ok"`
	}
}

type DelParentRelationshipRes struct {
	Data struct {
		string `example:"ok"`
	}
}
