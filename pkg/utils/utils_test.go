package utils

import (
	"testing"

	"github.com/Ricardxdev/payarc-sdk-go/pkg/extra"
)

type queryInput struct {
	Limit    int             `query:"limit,omitempty"`
	Page     int             `query:"page,omitempty"`
	StatusIn []string        `query:"status[in],omitempty"`
	SortedBy extra.SortOrder `query:"sortedBy,omitempty"`
	Excel    *bool           `query:"excel,omitempty"`
	Gte      string          `query:"created_at[gte],omitempty"`
	Skipped  string          `query:"-"`
	NoTag    string
}

func TestStructToQuery(t *testing.T) {
	excel := true
	params := StructToQuery(queryInput{
		Limit:    10,
		StatusIn: []string{"active", "trial"},
		SortedBy: extra.SortOrderDesc,
		Excel:    &excel,
		Gte:      "2024/01/01",
	})

	if params["limit"] != "10" {
		t.Errorf("expected limit 10, got %q", params["limit"])
	}
	if _, ok := params["page"]; ok {
		t.Errorf("expected page to be omitted, got %q", params["page"])
	}
	if params["status[in]"] != "active,trial" {
		t.Errorf("expected status[in] 'active,trial', got %q", params["status[in]"])
	}
	if params["sortedBy"] != "desc" {
		t.Errorf("expected sortedBy desc, got %q", params["sortedBy"])
	}
	if params["excel"] != "1" {
		t.Errorf("expected excel 1, got %q", params["excel"])
	}
	if params["created_at[gte]"] != "2024/01/01" {
		t.Errorf("expected created_at[gte] 2024/01/01, got %q", params["created_at[gte]"])
	}
	if _, ok := params["Skipped"]; ok {
		t.Error("expected Skipped to be omitted")
	}
	if _, ok := params["NoTag"]; ok {
		t.Error("expected NoTag to be omitted")
	}
}

func TestStructToQueryNilPointer(t *testing.T) {
	params := StructToQuery((*queryInput)(nil))
	if len(params) != 0 {
		t.Errorf("expected empty params, got %v", params)
	}
}
