package report

import "fmt"
type Summary struct {
	TotalRequests   int
	Passed          int
	Failed          int
	BreakingChanges int
	LatencyIssues   int
}

func (s *Summary) AddResult(breaking bool, hasLatencyIssue bool) {
	s.TotalRequests++

	if breaking {
		s.Failed++
		s.BreakingChanges++
	} else {
		s.Passed++
	}

	if hasLatencyIssue {
		s.LatencyIssues++
	}
}





func (s *Summary) Print() {

	successRate := 0.0
	if s.TotalRequests > 0 {
		successRate = float64(s.Passed) / float64(s.TotalRequests) * 100
	}

	fmt.Println("=================================")
	fmt.Println("TwinFlow Replay Summary")
	fmt.Println("=================================")

	fmt.Println("Total Requests:", s.TotalRequests)
	fmt.Println("Passed:", s.Passed)
	fmt.Println("Failed:", s.Failed)
	fmt.Println("Breaking Changes:", s.BreakingChanges)
	fmt.Println("Latency Issues:", s.LatencyIssues)

	fmt.Printf("\nSuccess Rate: %.2f%%\n", successRate)

	if s.Failed > 0 {
		fmt.Println("\nSTATUS:  UNSAFE TO DEPLOY")
	} else {
		fmt.Println("\nSTATUS:  SAFE TO DEPLOY")
	}
}