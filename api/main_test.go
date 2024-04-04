package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {
	ctx := context.Background()

	// Create a sample APIGatewayProxyRequest
	request := events.APIGatewayProxyRequest{}

	// Call the handler function
	response, err := handleRequest(ctx, request)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Validate the response
	if response.StatusCode != 200 {
		t.Errorf("expected status code 200, got %d", response.StatusCode)
	}

	var respBody events.APIGatewayProxyResponse
	if err := json.Unmarshal([]byte(response.Body), &respBody); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}

	// Check the message in the response body
	expectedMessage := "Hello From API!"
	var responseBody map[string]interface{}
	if err := json.Unmarshal([]byte(respBody.Body), &responseBody); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}

	if message, ok := responseBody["Message"].(string); !ok || message != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, message)
	}
}
