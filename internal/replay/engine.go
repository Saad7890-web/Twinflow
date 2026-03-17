package replay

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

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

	for scanner.Scan() {

		line := scanner.Bytes()

		var record models.TrafficRecord

		err := json.Unmarshal(line, &record)
		if err != nil {
			fmt.Println("Invalid record:", err)
			continue
		}

		err = e.replayRequest(record)
		if err != nil {
			fmt.Println("Replay error:", err)
		}
	}

	return nil
}


func (e *Engine) replayRequest(record models.TrafficRecord) error {

	url := e.target + record.Path

	req, err := http.NewRequest(
		record.Method,
		url,
		bytes.NewBuffer([]byte(record.RequestBody)),
	)

	if err != nil {
		return err
	}

	for k, v := range record.Headers {
		req.Header.Set(k, v)
	}

	start := time.Now()

	resp, err := e.client.Do(req)
	if err != nil {
		return err
	}

	duration := time.Since(start)

	fmt.Println("Replayed:", record.Method, record.Path)
	fmt.Println("Status:", resp.StatusCode)
	fmt.Println("Latency:", duration)

	return nil
}