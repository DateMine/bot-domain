package config

import (
	"fmt"
	"net"
	"os"
)

const grpcHostSenderEnvName = "GRPC_HOST_SENDER"
const grpcPortSenderEnvName = "GRPC_PORT_SENDER"

type GRPCConfig interface {
	GRPCAddress() string
}

type grpcConfig struct {
	grpcHostSender string
	grpcPortSender string
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostSenderEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", grpcHostSenderEnvName)
	}
	port := os.Getenv(grpcPortSenderEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", grpcPortSenderEnvName)
	}
	return &grpcConfig{
		grpcHostSender: host,
		grpcPortSender: port,
	}, nil
}

func (c grpcConfig) GRPCAddress() string {
	return net.JoinHostPort(c.grpcHostSender, c.grpcPortSender)
}
