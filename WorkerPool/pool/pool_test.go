package pool

import (
	"testing"
)

func TestProcess(t *testing.T) {
	jobs := []Job{
		{ID: 1, Multiply: 10},
		{ID: 2, Multiply: 10},
		{ID: 3, Multiply: 10},
	}

	results := Process(2, jobs)

	if len(results) != len(jobs) {
		t.Fatalf("expected %d results, got %d", len(jobs), len(results))
	}

	resultMap := make(map[int]int)
	for _, r := range results {
		resultMap[r.JobID] = r.Value
	}

	expected := map[int]int {
		1:10,
		2:20,
		3:30,
	}

	for id, expValue := range expected {
		actualValue, exists := resultMap[id]
		if !exists {
			t.Errorf("missing result for JobID %d", id)
			continue
		}
		if actualValue != expValue {
			t.Errorf("JobID %d: expected %d, got %d", id, expValue, actualValue)
		}
	}
}