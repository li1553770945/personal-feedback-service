package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/li1553770945/personal-notify-service/kitex_gen/notify/notifyservice"
	"github.com/li1553770945/personal-project-service/biz/infra/config"
	"log"
)

func NewNotifyClient(config *config.Config) notifyservice.Client {
	r, err := etcd.NewEtcdResolver(config.EtcdConfig.Endpoint)
	notifyClient, err := notifyservice.NewClient(
		config.RpcConfig.NotifyServiceName,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerConfig.ServiceName}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return notifyClient
}
