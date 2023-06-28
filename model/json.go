package model

import "time"

type LogFormat struct {
	TimeStampz    string `json:"timestampz"`
	Level         string `json:"level"`
	ServiceName   string `json:"service_name"`
	TraceID       string `json:"trace_id"`
	DeveloperTeam string `json:"developer_team"`
	Route         string `json:"route"`
	Payload       string `json:"payload"`
	SendAlert     bool   `json:"send_alert"`
	ProcessName   string `json:"process_name"`
	Message       string `json:"message"`
}

type Payload struct {
	ID        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}
