package config

import (
	"fmt"

	"github.com/zufardhiyaulhaq/istio-ratelimit-operator/api/v1alpha1"
	istioClientNetworking "istio.io/client-go/pkg/apis/networking/v1alpha3"
)

type GlobalRateLimitConfigEnvoyFilterFactory interface {
	Build() (*istioClientNetworking.EnvoyFilter, error)
}

func GetGlobalRateLimitConfigEnvoyFilterFactory(version string, config v1alpha1.GlobalRateLimitConfig) (GlobalRateLimitConfigEnvoyFilterFactory, error) {
	if version == "1.9" {
		return NewGlobalRateLimitConfig1_9EnvoyFilterBuilder(config), nil
	}

	return nil, fmt.Errorf("version not supported")
}
