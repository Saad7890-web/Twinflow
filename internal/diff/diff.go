package diff

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Saad7890-web/twinflow/pkg/models"
)

type Result struct {
	Path      string
	Breaking  bool
	Messages  []string
}

func Compare(record models.TrafficRecord) Result {

	result := Result{
		Path: record.Path,
	}

	// 1. Status Code Check
	if record.ResponseStatus != record.ReplayStatus {
		result.Breaking = true
		result.Messages = append(result.Messages,
			fmt.Sprintf("Status changed: %d → %d",
				record.ResponseStatus,
				record.ReplayStatus))
	}

	// 2. Body Comparison (JSON aware)
	oldJSON := parseJSON(record.ResponseBody)
	newJSON := parseJSON(record.ReplayBody)

	if oldJSON != nil && newJSON != nil {
		compareJSON(oldJSON, newJSON, &result)
	} else {
		if record.ResponseBody != record.ReplayBody {
			result.Messages = append(result.Messages, "Response body changed")
		}
	}

	// 3. Latency Check
	if record.ReplayLatency > record.LatencyMs*2 {
		result.Messages = append(result.Messages,
			fmt.Sprintf("Latency increased: %dms → %dms",
				record.LatencyMs,
				record.ReplayLatency))
	}

	return result
}

func parseJSON(data string) map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil
	}
	return result
}


func compareJSON(old, new map[string]interface{}, result *Result) {

	// Check removed fields
	for key := range old {
		if _, exists := new[key]; !exists {
			result.Breaking = true
			result.Messages = append(result.Messages,
				fmt.Sprintf("Field removed: %s", key))
		}
	}

	// Check added fields
	for key := range new {
		if _, exists := old[key]; !exists {
			result.Messages = append(result.Messages,
				fmt.Sprintf("Field added: %s", key))
		}
	}

	// Check value changes
	for key, oldVal := range old {
		if newVal, exists := new[key]; exists {
			if !reflect.DeepEqual(oldVal, newVal) {
				result.Messages = append(result.Messages,
					fmt.Sprintf("Field changed: %s", key))
			}
		}
	}
}