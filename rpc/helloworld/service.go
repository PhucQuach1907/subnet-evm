package helloworld

import (
	"context"
)

// Service implements the Hello World RPC service
type Service struct{}

// NewService creates a new Hello World service
func NewService() *Service {
	return &Service{}
}

// HelloWorld returns a greeting message
func (s *Service) HelloWorld(ctx context.Context) (string, error) {
	return "Hello World", nil
} 