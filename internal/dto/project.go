package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type (
	// HeaderAttrID ...
	HeaderAttrID string

	// HeaderAttribute ...
	HeaderAttribute struct {
		ID   HeaderAttrID `json:"attrId" bson:"attrId"`
		Name string       `json:"attrName" bson:"attrName"`
	}

	// ProjectWork ...
	ProjectWork struct {
		WbsCode       string   `json:"wbs_code" bson:"wbsCode"`
		Name          string   `json:"name" bson:"name"`
		Length        int      `json:"length" bson:"length"`
		PreviousWorks []string `json:"prev_works" bson:"prevWorks"`
	}

	// Project ...
	Project struct {
		ID              bson.ObjectID     `json:"id" bson:"_id"`
		Title           string            `json:"title" bson:"title"`
		StartDate       time.Time         `json:"projectStartDate" bson:"projectStartDate"`
		DateFormat      string            `json:"dateDisplayTemplate" bson:"dateDisplayTemplate"`
		IsSuppressZeros bool              `json:"isSuppressZeros" bson:"isSuppressZeros"`
		HeaderAttrs     []HeaderAttribute `json:"projectHeaderAttributes" bson:"projectHeaderAttributes"`
		Works           []ProjectWork     `json:"projectWorksList" bson:"projectWorksList"`
	}
)

const (
	// HeaderAttrCode ...
	HeaderAttrCode = HeaderAttrID("wbs_code")
	//HeaderAttrName ...
	HeaderAttrName = HeaderAttrID("work_name")
	// HeaderAttrLength ...
	HeaderAttrLength = HeaderAttrID("length")
	// HeaderAttrPercentComplete ...
	HeaderAttrPercentComplete = HeaderAttrID("percent_complete")
	// HeaderAttrStartDate ...
	HeaderAttrStartDate = HeaderAttrID("start_date")
	// HeaderAttrFinishDate ...
	HeaderAttrFinishDate = HeaderAttrID("finish_date")
	// HeaderAttrPrevWorks ...
	HeaderAttrPrevWorks = HeaderAttrID("prev_works")
)

// ProjectTemplate ...
var ProjectTemplate = Project{
	Title:           "Sample Project",
	DateFormat:      "yy/MM/dd",
	IsSuppressZeros: true,
	HeaderAttrs: []HeaderAttribute{
		{ID: HeaderAttrCode, Name: "WBS"},
		{ID: HeaderAttrName, Name: "Work"},
		{ID: HeaderAttrLength, Name: "Len. (d)"},
		{ID: HeaderAttrPercentComplete, Name: "% done"},
		{ID: HeaderAttrStartDate, Name: "Start at"},
		{ID: HeaderAttrFinishDate, Name: "Finish at"},
		{ID: HeaderAttrPrevWorks, Name: "Prev."},
	},
	Works: []ProjectWork{
		{
			WbsCode: "",
			Name:    "Project \"Creation of a billing information system\" (DEV-26-10)",
		},
		{
			WbsCode: "1",
			Name:    "Control milestones",
		},
		{
			WbsCode: "1.1",
			Name:    "Project started",
			Length:  1,
		},
		{
			WbsCode:       "1.2",
			Name:          "",
			Length:        1,
			PreviousWorks: []string{"1.1"},
		},
		{
			WbsCode: "2",
			Name:    "Analysis",
		},
		{
			WbsCode: "2.1",
			Name:    "Requirements scope definition",
		},
		{
			WbsCode: "2.1.1",
			Name:    "Business process analysis AS-IS",
			Length:  6,
		},
	},
}
