// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/virtual_service.proto

// Configuration affecting traffic routing. Here are a few terms useful to define
// in the context of traffic routing.
//
// `Service` a unit of application behavior bound to a unique name in a
// service registry. Services consist of multiple network *endpoints*
// implemented by workload instances running on pods, containers, VMs etc.
//
// `Service versions (a.k.a. subsets)` - In a continuous deployment
// scenario, for a given service, there can be distinct subsets of
// instances running different variants of the application binary. These
// variants are not necessarily different API versions. They could be
// iterative changes to the same service, deployed in different
// environments (prod, staging, dev, etc.). Common scenarios where this
// occurs include A/B testing, canary rollouts, etc. The choice of a
// particular version can be decided based on various criterion (headers,
// url, etc.) and/or by weights assigned to each version. Each service has
// a default version consisting of all its instances.
//
// `Source` - A downstream client calling a service.
//
// `Host` - The address used by a client when attempting to connect to a
// service.
//
// `Access model` - Applications address only the destination service
// (Host) without knowledge of individual service versions (subsets). The
// actual choice of the version is determined by the proxy/sidecar, enabling the
// application code to decouple itself from the evolution of dependent
// services.
//
// A `VirtualService` defines a set of traffic routing rules to apply when a host is
// addressed. Each routing rule defines matching criteria for traffic of a specific
// protocol. If the traffic is matched, then it is sent to a named destination service
// (or subset/version of it) defined in the registry.
//
// The source of traffic can also be matched in a routing rule. This allows routing
// to be customized for specific client contexts.
//
// The following example on Kubernetes, routes all HTTP traffic by default to
// pods of the reviews service with label "version: v1". In addition,
// HTTP requests with path starting with /wpcatalog/ or /consumercatalog/ will
// be rewritten to /newcatalog and sent to pods with label "version: v2".
//
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: VirtualService
// metadata:
//   name: reviews-route
// spec:
//   hosts:
//   - reviews.prod.svc.cluster.local
//   http:
//   - name: "reviews-v2-routes"
//     match:
//     - uri:
//         prefix: "/wpcatalog"
//     - uri:
//         prefix: "/consumercatalog"
//     rewrite:
//       uri: "/newcatalog"
//     route:
//     - destination:
//         host: reviews.prod.svc.cluster.local
//         subset: v2
//   - name: "reviews-v1-route"
//     route:
//     - destination:
//         host: reviews.prod.svc.cluster.local
//         subset: v1
// ```
//
// A subset/version of a route destination is identified with a reference
// to a named service subset which must be declared in a corresponding
// `DestinationRule`.
//
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: reviews-destination
// spec:
//   host: reviews.prod.svc.cluster.local
//   subsets:
//   - name: v1
//     labels:
//       version: v1
//   - name: v2
//     labels:
//       version: v2
// ```
//

package v1alpha3

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler supporting oneof fields for StringMatch
func (this *StringMatch) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for StringMatch
func (this *StringMatch) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for HTTPFaultInjection_Delay
func (this *HTTPFaultInjection_Delay) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for HTTPFaultInjection_Delay
func (this *HTTPFaultInjection_Delay) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for HTTPFaultInjection_Abort
func (this *HTTPFaultInjection_Abort) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for HTTPFaultInjection_Abort
func (this *HTTPFaultInjection_Abort) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler supporting oneof fields for PortSelector
func (this *PortSelector) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for PortSelector
func (this *PortSelector) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	VirtualServiceMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	VirtualServiceUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)