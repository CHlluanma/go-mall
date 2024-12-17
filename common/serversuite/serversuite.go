package serversuite

import (
	"github.com/chhz0/go-mall-kitex/common/mtl"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegisterAddr       string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
	}

	r, err := consul.NewConsulRegister(s.RegisterAddr)
	if err != nil {
		klog.Fatal(err)
	}

	_ = provider.NewOpenTelemetryProvider(provider.WithSdkTracerProvider(mtl.TracerProvider),
		provider.WithEnableMetrics(false),
	)

	opts = append(opts,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithTracer(prometheus.NewServerTracer("",
			"",
			prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry)),
		),
	)

	opts = append(opts,
		server.WithRegistry(r),
	)

	return opts
}
