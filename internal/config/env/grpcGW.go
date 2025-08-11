package env

import (
	"fmt"
	"net"
	"os"
)

const (
	grpcGWHostEnvName = "GRPC_GW_HOST"
	grpcGWPortEnvName = "GRPC_GW_PORT"
)

type grpcGWConfig struct {
	Host string
	Port string
}

func NewGRPCGWConfig() (*grpcGWConfig, error) {
	host := os.Getenv(grpcGWHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", grpcHostEnvName)
	}

	port := os.Getenv(grpcGWPortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", grpcHostEnvName)
	}

	return &grpcGWConfig{
		Host: host,
		Port: port,
	}, nil
}

func (cfg grpcGWConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}
