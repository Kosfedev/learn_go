package env

import (
	"fmt"
	"net"
	"os"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type grpcConfig struct {
	Host string
	Port string
}

func NewGRPCConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", grpcHostEnvName)
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", grpcHostEnvName)
	}

	return &grpcConfig{
		Host: host,
		Port: port,
	}, nil
}

func (cfg grpcConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}
