package graph_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/opentickettech/graph-exercise/graph"
)

// newGraph returns a new graph used for testing.
func newGraph() *graph.Node {
	joost := &graph.Node{Name: "joost"}
	chris := &graph.Node{Name: "chris"}
	mart := &graph.Node{Name: "mart"}
	marta := &graph.Node{Name: "marta"}
	peter := &graph.Node{Name: "peter"}
	jos := &graph.Node{Name: "jos"}
	dennis := &graph.Node{Name: "dennis"}

	mart.AddChild(peter)
	joost.AddChild(chris)
	dennis.AddChild(peter)
	chris.AddChild(mart)
	chris.AddChild(peter)
	dennis.AddChild(marta)
	joost.AddChild(peter)
	joost.AddChild(jos)
	jos.AddChild(dennis)

	// Joost is the root-node since he is the CEO.
	return joost
}

func TestOrdered(t *testing.T) {
	sorted := newGraph().Ordered()
	fmt.Println(sorted)

	reportMap := map[string][]string{
		"joost":  {"chris", "peter", "jos"},
		"chris":  {"peter", "mart"},
		"mart":   {"peter"},
		"marta":  {},
		"peter":  {},
		"jos":    {"dennis"},
		"dennis": {"marta", "peter"},
	}

	// Check whether no report has a lower index than their manager.
	for i, manager := range sorted {
		for _, report := range reportMap[manager] {
			if i >= slices.Index(sorted, report) {
				fmt.Println("Report", report, "occurs before manager", manager)
				t.Fail()
			}
		}
	}

	// Check for duplicates
	for _, manager := range sorted {
		if _, ok := reportMap[manager]; !ok {
			// Already deleted
			t.FailNow()
		}

		delete(reportMap, manager)
	}

	// Check ordered-list completion. I.e., are all persons present in the list.
	if len(reportMap) > 0 {
		t.Fail()
	}
}
