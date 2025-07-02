package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"smart-contract-api/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRegisterProvider(t *testing.T) {
	var times []time.Duration
	for i := 0; i < 50; i++ {
		provider := models.Provider{
			Email:    fmt.Sprintf("test-%d@gmail.com", i),
			Password: "testPassword123",
			OrganizationInfo: models.OrganizationInfo{
				Name:    fmt.Sprintf("TestOrg-%d", i),
				Address: "Bucharest",
			},
		}
		providerBytes, _ := json.MarshalIndent(provider, "", "  ")

		now := time.Now()
		resp, err := http.Post(
			"http://172.18.0.2:30081/providers", "application/json",
			bytes.NewBuffer(providerBytes),
		)
		times = append(times, time.Since(now))
		if !assert.NoError(t, err) {
			t.Fail()
		}
		if !assert.Equal(t, http.StatusCreated, resp.StatusCode) {
			t.Fail()
		}
	}

	sum := 0.0
	for i := 0; i < len(times); i++ {
		sum += times[i].Seconds()
	}
	average := sum / float64(len(times))
	t.Log("average time", average)
}

func TestEnrollStudents(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAyNzgzMjcsImlhdCI6MTc1MDE5MTkyNywiaXNzIjoiRWR1Q2hhaW4iLCJwaWQiOjI1OSwidHRsIjoxfQ.MQIKTeXqDWnq1GozAZdEgHcHINz6wN4sdvkQTr-wa6c"

	var times []time.Duration
	for i := 0; i < 50; i++ {
		var students []models.Student

		for j := 0; j < 6000; j++ {
			id, _ := uuid.NewRandom()
			student := models.Student{
				ExternalID:    id.String(),
				WalletAddress: "0x8aADBc68b378C61af7b6f6E7BA0b95e85a7f4ED5",
			}
			students = append(students, student)
		}

		payload, _ := json.MarshalIndent(students, "", "  ")

		req, err := http.NewRequest(
			http.MethodPost, "http://172.18.0.2:30081/students",
			bytes.NewBuffer(payload),
		)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Add("Authorization", "Bearer "+token)
		req.Header.Add("Content-Type", "application/json")

		now := time.Now()
		resp, err := http.DefaultClient.Do(req)
		times = append(times, time.Since(now))
		if !assert.NoError(t, err) {
			t.Fail()
		}
		if !assert.Equal(t, http.StatusCreated, resp.StatusCode) {
			t.Fail()
		}
	}

	sum := 0.0
	for i := 0; i < len(times); i++ {
		sum += times[i].Seconds()
	}
	average := sum / float64(len(times))
	t.Log("average time", average)
}

func TestRetrieveAllStudents(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTAyNzgzMjcsImlhdCI6MTc1MDE5MTkyNywiaXNzIjoiRWR1Q2hhaW4iLCJwaWQiOjI1OSwidHRsIjoxfQ.MQIKTeXqDWnq1GozAZdEgHcHINz6wN4sdvkQTr-wa6c"

	var times []time.Duration
	for i := 0; i < 50; i++ {
		req, err := http.NewRequest(
			http.MethodGet, "http://172.18.0.2:30081/students",
			nil,
		)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Add("Authorization", "Bearer "+token)

		now := time.Now()
		resp, err := http.DefaultClient.Do(req)
		times = append(times, time.Since(now))
		if !assert.NoError(t, err) {
			t.Fail()
		}
		if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
			t.Fail()
		}
	}

	sum := 0.0
	for i := 0; i < len(times); i++ {
		sum += times[i].Seconds()
	}
	average := sum / float64(len(times))
	t.Log("average time", average)
}
