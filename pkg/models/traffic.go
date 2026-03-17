package models

type TrafficRecord struct {
	Timestamp   string            `json:"timestamp"`
	Method      string            `json:"method"`
	Path        string            `json:"path"`
	Headers     map[string]string `json:"headers"`
	RequestBody string            `json:"request_body"`

	ResponseStatus int    `json:"response_status"`
	ResponseBody   string `json:"response_body"`

	LatencyMs int64 `json:"latency_ms"`
}