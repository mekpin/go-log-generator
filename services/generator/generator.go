package generator

import (
	"encoding/json"
	"fmt"
	"go-log-generator/model"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GenerateLog(logcount int) {
	rand.Seed(time.Now().UnixNano())

	levels := []string{"INFO", "WARNING", "ERROR"}
	messages := []string{"User logged in", "Invalid credentials", "Network connection lost"}
	servicenames := []string{"go-healthcheck", "go-pingcheck", "go-remote"}
	uuid, _ := GenerateUUIDv1()
	developers := []string{"B2B", "B2C", "core-transaction"}
	routes := []string{"/api/delete", "/health/check", "/ping/check"}
	payloads, _ := GenerateRandomPayload()
	alerts := []bool{true, false}
	processnames := []string{"User go to market", "owner go to sleep", "car brakes hard"}

	for i := 0; i < logcount; i++ {
		log := model.LogFormat{
			TimeStampz:    time.Now().Format(time.RFC3339),
			Level:         levels[rand.Intn(len(levels))],
			ServiceName:   servicenames[rand.Intn(len(servicenames))],
			TraceID:       uuid,
			DeveloperTeam: developers[rand.Intn(len(developers))],
			Route:         routes[rand.Intn(len(routes))],
			Payload:       string(payloads),
			SendAlert:     alerts[rand.Intn(len(alerts))],
			ProcessName:   processnames[rand.Intn(len(processnames))],
			Message:       messages[rand.Intn(len(messages))],
		}

		jsonLog, err := json.Marshal(log)
		if err != nil {
			fmt.Println("Error marshaling log:", err)
			return
		}

		fmt.Println(string(jsonLog))
	}
}

func GenerateUUIDv1() (string, error) {
	uuidV1, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return uuidV1.String(), nil
}

func GenerateRandomPayload() ([]byte, error) {
	payload := model.Payload{
		ID:        rand.Intn(100),
		Timestamp: time.Now(),
		Message:   generateRandomString(10),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return jsonPayload, nil
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
