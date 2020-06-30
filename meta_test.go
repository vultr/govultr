package govultr

import (
	"encoding/json"
	"testing"
)

var metaBytes = []byte(`
	{
        "total": 11,
        "links": {
            "next": "bmV4dF9fMTMxOTgxNQ==",
            "prev": ""
        }
	}
`)

func TestMeta(t *testing.T) {
	var meta *Meta

	if err := json.Unmarshal(metaBytes, &meta); err != nil {
		t.Fatal(err)
	}

	if meta.Total != 11 {
		t.Fatal("Total did not equal 11")
	}

	if meta.Links.Next != "bmV4dF9fMTMxOTgxNQ==" {
		t.Fatal("Next cursor did not equal bmV4dF9fMTMxOTgxNQ==")
	}

	if meta.Links.Prev != "" {
		t.Fatal("Previous cursor was not empty")
	}
}
