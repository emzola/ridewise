package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Registry interface {
	Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error
	Deregister(ctx context.Context, serviceName string, instanceID string) error
	ReportHealthyState(instanceID string, serviceName string) error
	ServiceAddresses(ctx context.Context, serviceName string) ([]string, error)
}

var ErrNotFound = errors.New("no service addresses found")

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
