package replay

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Saad7890-web/twinflow/internal/diff"
	"github.com/Saad7890-web/twinflow/internal/report"
	"github.com/Saad7890-web/twinflow/pkg/models"
)

type Engine struct {
	target string
	client *http.Client
}

func NewEngine(target string) *Engine {
	return &Engine{
		target: target,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (e *Engine) Replay(captureFile string) error {

	file, err := os.Open(captureFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	
	summary := &report.Summary{}

	for scanner.Scan() {

		line := scanner.Bytes()

		var record models.TrafficRecord

		err := json.Unmarshal(line, &record)
		if err != nil {
			fmt.Println("Invalid record:", err)
			continue
		}

		
		result, hasLatencyIssue, err := e.replayRequest(record)
		if err != nil {
			fmt.Println("Replay error:", err)
			continue
		}

		summary.AddResult(result.Breaking, hasLatencyIssue)
	}

	
	summary.Print()

	return nil
}

func (e *Engine) replayRequest(record models.TrafficRecord) (diff.Result, bool, error) {

	url := e.target + record.Path

	req, err := http.NewRequest(
		record.Method,
		url,
		bytes.NewBuffer([]byte(record.RequestBody)),
	)
	if err != nil {
		return diff.Result{}, false, err
	}

	for k, v := range record.Headers {
		req.Header.Set(k, v)
	}

	start := time.Now()

	resp, err := e.client.Do(req)
	if err != nil {
		return diff.Result{}, false, err
	}
	defer resp.Body.Close() 

	duration := time.Since(start)

	bodyBytes, _ := io.ReadAll(resp.Body)

	record.ReplayStatus = resp.StatusCode
	record.ReplayBody = string(bodyBytes)
	record.ReplayLatency = duration.Milliseconds()

	
	result := diff.Compare(record)

	
	hasLatencyIssue := record.ReplayLatency > record.LatencyMs*2

	
	fmt.Println("-----")
	fmt.Println("REQUEST:", record.Method, record.Path)

	if result.Breaking {
		fmt.Println("BREAKING CHANGE")
	} else {
		fmt.Println("OK")
	}

	for _, msg := range result.Messages {
		fmt.Println("-", msg)
	}

	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Latency:", duration)

	return result, hasLatencyIssue, nil
}