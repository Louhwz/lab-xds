package main

import (
	"log"
	"time"

	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func BuildCluster() []cache.Resource {
	var clusterName1 = "service_bbc"
	log.Println(">>>>>>>>>>>>>>>>>>> creating cluster ", clusterName1)

	h := &core.Address{Address: &core.Address_SocketAddress{
		SocketAddress: &core.SocketAddress{
			Address:  "127.0.0.1",
			Protocol: core.SocketAddress_TCP,
			PortSpecifier: &core.SocketAddress_PortValue{
				PortValue: uint32(18080),
			},
		},
	}}

	hLightest := &core.Address{Address: &core.Address_SocketAddress{SocketAddress: &core.SocketAddress{
		Protocol: core.SocketAddress_TCP,
		Address:  "127.0.0.1",
		PortSpecifier: &core.SocketAddress_PortValue{
			PortValue: uint32(18081),
		},
	}}}

	return []cache.Resource{
		&v2.Cluster{
			Name:                 clusterName1,
			ConnectTimeout:       ptypes.DurationProto(2 * time.Second),
			ClusterDiscoveryType: &v2.Cluster_Type{Type: v2.Cluster_STRICT_DNS},
			LbPolicy:             v2.Cluster_ROUND_ROBIN,
			LoadAssignment: &v2.ClusterLoadAssignment{
				ClusterName: clusterName1,
				Endpoints: []*endpoint.LocalityLbEndpoints{
					{
						LbEndpoints: []*endpoint.LbEndpoint{
							{
								HostIdentifier: &endpoint.LbEndpoint_Endpoint{
									Endpoint: &endpoint.Endpoint{
										Address: h,
									},
								},
								LoadBalancingWeight: &wrappers.UInt32Value{
									Value: 1000,
								},
							},
							{
								HostIdentifier: &endpoint.LbEndpoint_Endpoint{
									Endpoint: &endpoint.Endpoint{
										Address: hLightest,
									},
								},
								LoadBalancingWeight: &wrappers.UInt32Value{
									Value: 1,
								},
							},
						},
					},
				},
			},
		},
	}
}
