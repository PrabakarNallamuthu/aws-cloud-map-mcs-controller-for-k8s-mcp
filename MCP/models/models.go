package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetNamespaceRequest represents the GetNamespaceRequest schema from the OpenAPI specification
type GetNamespaceRequest struct {
	Id interface{} `json:"Id"`
}

// ListServicesRequest represents the ListServicesRequest schema from the OpenAPI specification
type ListServicesRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListInstancesRequest represents the ListInstancesRequest schema from the OpenAPI specification
type ListInstancesRequest struct {
	Serviceid interface{} `json:"ServiceId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Instance represents the Instance schema from the OpenAPI specification
type Instance struct {
	Attributes interface{} `json:"Attributes,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Id interface{} `json:"Id"`
}

// GetInstancesHealthStatusResponse represents the GetInstancesHealthStatusResponse schema from the OpenAPI specification
type GetInstancesHealthStatusResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DeleteServiceRequest represents the DeleteServiceRequest schema from the OpenAPI specification
type DeleteServiceRequest struct {
	Id interface{} `json:"Id"`
}

// RegisterInstanceResponse represents the RegisterInstanceResponse schema from the OpenAPI specification
type RegisterInstanceResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// PublicDnsNamespaceProperties represents the PublicDnsNamespaceProperties schema from the OpenAPI specification
type PublicDnsNamespaceProperties struct {
	Dnsproperties interface{} `json:"DnsProperties"`
}

// ListNamespacesResponse represents the ListNamespacesResponse schema from the OpenAPI specification
type ListNamespacesResponse struct {
	Namespaces interface{} `json:"Namespaces,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// UpdateServiceRequest represents the UpdateServiceRequest schema from the OpenAPI specification
type UpdateServiceRequest struct {
	Id interface{} `json:"Id"`
	Service interface{} `json:"Service"`
}

// PrivateDnsNamespaceProperties represents the PrivateDnsNamespaceProperties schema from the OpenAPI specification
type PrivateDnsNamespaceProperties struct {
	Dnsproperties interface{} `json:"DnsProperties"`
}

// OperationFilter represents the OperationFilter schema from the OpenAPI specification
type OperationFilter struct {
	Condition interface{} `json:"Condition,omitempty"`
	Name interface{} `json:"Name"`
	Values interface{} `json:"Values"`
}

// HealthCheckCustomConfig represents the HealthCheckCustomConfig schema from the OpenAPI specification
type HealthCheckCustomConfig struct {
	Failurethreshold interface{} `json:"FailureThreshold,omitempty"`
}

// InstanceHealthStatusMap represents the InstanceHealthStatusMap schema from the OpenAPI specification
type InstanceHealthStatusMap struct {
}

// ServiceChange represents the ServiceChange schema from the OpenAPI specification
type ServiceChange struct {
	Description interface{} `json:"Description,omitempty"`
	Dnsconfig interface{} `json:"DnsConfig,omitempty"`
	Healthcheckconfig interface{} `json:"HealthCheckConfig,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Tags interface{} `json:"Tags"`
}

// HttpInstanceSummary represents the HttpInstanceSummary schema from the OpenAPI specification
type HttpInstanceSummary struct {
	Servicename interface{} `json:"ServiceName,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Healthstatus interface{} `json:"HealthStatus,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Namespacename interface{} `json:"NamespaceName,omitempty"`
}

// HttpProperties represents the HttpProperties schema from the OpenAPI specification
type HttpProperties struct {
	Httpname interface{} `json:"HttpName,omitempty"`
}

// OperationTargetsMap represents the OperationTargetsMap schema from the OpenAPI specification
type OperationTargetsMap struct {
}

// DeleteServiceResponse represents the DeleteServiceResponse schema from the OpenAPI specification
type DeleteServiceResponse struct {
}

// CreateServiceRequest represents the CreateServiceRequest schema from the OpenAPI specification
type CreateServiceRequest struct {
	Namespaceid interface{} `json:"NamespaceId,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Dnsconfig interface{} `json:"DnsConfig,omitempty"`
	Healthcheckcustomconfig interface{} `json:"HealthCheckCustomConfig,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Healthcheckconfig interface{} `json:"HealthCheckConfig,omitempty"`
	Name interface{} `json:"Name"`
}

// DeregisterInstanceRequest represents the DeregisterInstanceRequest schema from the OpenAPI specification
type DeregisterInstanceRequest struct {
	Serviceid interface{} `json:"ServiceId"`
	Instanceid interface{} `json:"InstanceId"`
}

// ListInstancesResponse represents the ListInstancesResponse schema from the OpenAPI specification
type ListInstancesResponse struct {
	Instances interface{} `json:"Instances,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetOperationResponse represents the GetOperationResponse schema from the OpenAPI specification
type GetOperationResponse struct {
	Operation interface{} `json:"Operation,omitempty"`
}

// CreateHttpNamespaceResponse represents the CreateHttpNamespaceResponse schema from the OpenAPI specification
type CreateHttpNamespaceResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// DeleteNamespaceResponse represents the DeleteNamespaceResponse schema from the OpenAPI specification
type DeleteNamespaceResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// CreatePublicDnsNamespaceRequest represents the CreatePublicDnsNamespaceRequest schema from the OpenAPI specification
type CreatePublicDnsNamespaceRequest struct {
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
	Properties interface{} `json:"Properties,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListOperationsResponse represents the ListOperationsResponse schema from the OpenAPI specification
type ListOperationsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Operations interface{} `json:"Operations,omitempty"`
}

// DnsProperties represents the DnsProperties schema from the OpenAPI specification
type DnsProperties struct {
	Soa interface{} `json:"SOA,omitempty"`
	Hostedzoneid interface{} `json:"HostedZoneId,omitempty"`
}

// NamespaceFilter represents the NamespaceFilter schema from the OpenAPI specification
type NamespaceFilter struct {
	Name interface{} `json:"Name"`
	Values interface{} `json:"Values"`
	Condition interface{} `json:"Condition,omitempty"`
}

// ServiceSummary represents the ServiceSummary schema from the OpenAPI specification
type ServiceSummary struct {
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Healthcheckcustomconfig interface{} `json:"HealthCheckCustomConfig,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Healthcheckconfig interface{} `json:"HealthCheckConfig,omitempty"`
	Dnsconfig interface{} `json:"DnsConfig,omitempty"`
}

// UpdatePublicDnsNamespaceRequest represents the UpdatePublicDnsNamespaceRequest schema from the OpenAPI specification
type UpdatePublicDnsNamespaceRequest struct {
	Id interface{} `json:"Id"`
	Namespace interface{} `json:"Namespace"`
	Updaterrequestid interface{} `json:"UpdaterRequestId,omitempty"`
}

// Attributes represents the Attributes schema from the OpenAPI specification
type Attributes struct {
}

// UpdateInstanceCustomHealthStatusRequest represents the UpdateInstanceCustomHealthStatusRequest schema from the OpenAPI specification
type UpdateInstanceCustomHealthStatusRequest struct {
	Instanceid interface{} `json:"InstanceId"`
	Serviceid interface{} `json:"ServiceId"`
	Status interface{} `json:"Status"`
}

// UpdateHttpNamespaceRequest represents the UpdateHttpNamespaceRequest schema from the OpenAPI specification
type UpdateHttpNamespaceRequest struct {
	Updaterrequestid interface{} `json:"UpdaterRequestId,omitempty"`
	Id interface{} `json:"Id"`
	Namespace interface{} `json:"Namespace"`
}

// GetInstanceRequest represents the GetInstanceRequest schema from the OpenAPI specification
type GetInstanceRequest struct {
	Instanceid interface{} `json:"InstanceId"`
	Serviceid interface{} `json:"ServiceId"`
}

// HttpNamespaceChange represents the HttpNamespaceChange schema from the OpenAPI specification
type HttpNamespaceChange struct {
	Description interface{} `json:"Description"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// PublicDnsPropertiesMutable represents the PublicDnsPropertiesMutable schema from the OpenAPI specification
type PublicDnsPropertiesMutable struct {
	Soa interface{} `json:"SOA"`
}

// Operation represents the Operation schema from the OpenAPI specification
type Operation struct {
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Updatedate interface{} `json:"UpdateDate,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
}

// DnsConfigChange represents the DnsConfigChange schema from the OpenAPI specification
type DnsConfigChange struct {
	Dnsrecords interface{} `json:"DnsRecords"`
}

// SOA represents the SOA schema from the OpenAPI specification
type SOA struct {
	Ttl interface{} `json:"TTL"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Tagkeys interface{} `json:"TagKeys"`
	Resourcearn interface{} `json:"ResourceARN"`
}

// HealthCheckConfig represents the HealthCheckConfig schema from the OpenAPI specification
type HealthCheckConfig struct {
	Failurethreshold interface{} `json:"FailureThreshold,omitempty"`
	Resourcepath interface{} `json:"ResourcePath,omitempty"`
	TypeField interface{} `json:"Type"`
}

// UpdateServiceResponse represents the UpdateServiceResponse schema from the OpenAPI specification
type UpdateServiceResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// DiscoverInstancesRequest represents the DiscoverInstancesRequest schema from the OpenAPI specification
type DiscoverInstancesRequest struct {
	Namespacename interface{} `json:"NamespaceName"`
	Optionalparameters interface{} `json:"OptionalParameters,omitempty"`
	Queryparameters interface{} `json:"QueryParameters,omitempty"`
	Servicename interface{} `json:"ServiceName"`
	Healthstatus interface{} `json:"HealthStatus,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// DeregisterInstanceResponse represents the DeregisterInstanceResponse schema from the OpenAPI specification
type DeregisterInstanceResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// OperationSummary represents the OperationSummary schema from the OpenAPI specification
type OperationSummary struct {
	Id interface{} `json:"Id,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// PublicDnsNamespaceChange represents the PublicDnsNamespaceChange schema from the OpenAPI specification
type PublicDnsNamespaceChange struct {
	Description interface{} `json:"Description,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
}

// CreateServiceResponse represents the CreateServiceResponse schema from the OpenAPI specification
type CreateServiceResponse struct {
	Service interface{} `json:"Service,omitempty"`
}

// PrivateDnsPropertiesMutable represents the PrivateDnsPropertiesMutable schema from the OpenAPI specification
type PrivateDnsPropertiesMutable struct {
	Soa interface{} `json:"SOA"`
}

// GetInstanceResponse represents the GetInstanceResponse schema from the OpenAPI specification
type GetInstanceResponse struct {
	Instance interface{} `json:"Instance,omitempty"`
}

// GetServiceRequest represents the GetServiceRequest schema from the OpenAPI specification
type GetServiceRequest struct {
	Id interface{} `json:"Id"`
}

// CreatePrivateDnsNamespaceResponse represents the CreatePrivateDnsNamespaceResponse schema from the OpenAPI specification
type CreatePrivateDnsNamespaceResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// InstanceSummary represents the InstanceSummary schema from the OpenAPI specification
type InstanceSummary struct {
	Id interface{} `json:"Id,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
}

// CreatePrivateDnsNamespaceRequest represents the CreatePrivateDnsNamespaceRequest schema from the OpenAPI specification
type CreatePrivateDnsNamespaceRequest struct {
	Name interface{} `json:"Name"`
	Properties interface{} `json:"Properties,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Vpc interface{} `json:"Vpc"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// PublicDnsNamespacePropertiesChange represents the PublicDnsNamespacePropertiesChange schema from the OpenAPI specification
type PublicDnsNamespacePropertiesChange struct {
	Dnsproperties interface{} `json:"DnsProperties"`
}

// GetNamespaceResponse represents the GetNamespaceResponse schema from the OpenAPI specification
type GetNamespaceResponse struct {
	Namespace interface{} `json:"Namespace,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// UpdateHttpNamespaceResponse represents the UpdateHttpNamespaceResponse schema from the OpenAPI specification
type UpdateHttpNamespaceResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// UpdatePrivateDnsNamespaceResponse represents the UpdatePrivateDnsNamespaceResponse schema from the OpenAPI specification
type UpdatePrivateDnsNamespaceResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// PublicDnsPropertiesMutableChange represents the PublicDnsPropertiesMutableChange schema from the OpenAPI specification
type PublicDnsPropertiesMutableChange struct {
	Soa interface{} `json:"SOA"`
}

// Namespace represents the Namespace schema from the OpenAPI specification
type Namespace struct {
	TypeField interface{} `json:"Type,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
	Servicecount interface{} `json:"ServiceCount,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// RegisterInstanceRequest represents the RegisterInstanceRequest schema from the OpenAPI specification
type RegisterInstanceRequest struct {
	Instanceid interface{} `json:"InstanceId"`
	Serviceid interface{} `json:"ServiceId"`
	Attributes interface{} `json:"Attributes"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
}

// GetOperationRequest represents the GetOperationRequest schema from the OpenAPI specification
type GetOperationRequest struct {
	Operationid interface{} `json:"OperationId"`
}

// UpdatePrivateDnsNamespaceRequest represents the UpdatePrivateDnsNamespaceRequest schema from the OpenAPI specification
type UpdatePrivateDnsNamespaceRequest struct {
	Id interface{} `json:"Id"`
	Namespace interface{} `json:"Namespace"`
	Updaterrequestid interface{} `json:"UpdaterRequestId,omitempty"`
}

// SOAChange represents the SOAChange schema from the OpenAPI specification
type SOAChange struct {
	Ttl interface{} `json:"TTL"`
}

// CreatePublicDnsNamespaceResponse represents the CreatePublicDnsNamespaceResponse schema from the OpenAPI specification
type CreatePublicDnsNamespaceResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// PrivateDnsNamespacePropertiesChange represents the PrivateDnsNamespacePropertiesChange schema from the OpenAPI specification
type PrivateDnsNamespacePropertiesChange struct {
	Dnsproperties interface{} `json:"DnsProperties"`
}

// DnsRecord represents the DnsRecord schema from the OpenAPI specification
type DnsRecord struct {
	Ttl interface{} `json:"TTL"`
	TypeField interface{} `json:"Type"`
}

// GetInstancesHealthStatusRequest represents the GetInstancesHealthStatusRequest schema from the OpenAPI specification
type GetInstancesHealthStatusRequest struct {
	Serviceid interface{} `json:"ServiceId"`
	Instances interface{} `json:"Instances,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DnsConfig represents the DnsConfig schema from the OpenAPI specification
type DnsConfig struct {
	Dnsrecords interface{} `json:"DnsRecords"`
	Namespaceid interface{} `json:"NamespaceId,omitempty"`
	Routingpolicy interface{} `json:"RoutingPolicy,omitempty"`
}

// NamespaceProperties represents the NamespaceProperties schema from the OpenAPI specification
type NamespaceProperties struct {
	Httpproperties interface{} `json:"HttpProperties,omitempty"`
	Dnsproperties interface{} `json:"DnsProperties,omitempty"`
}

// ListNamespacesRequest represents the ListNamespacesRequest schema from the OpenAPI specification
type ListNamespacesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// ListOperationsRequest represents the ListOperationsRequest schema from the OpenAPI specification
type ListOperationsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// PrivateDnsNamespaceChange represents the PrivateDnsNamespaceChange schema from the OpenAPI specification
type PrivateDnsNamespaceChange struct {
	Description interface{} `json:"Description,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
}

// DeleteNamespaceRequest represents the DeleteNamespaceRequest schema from the OpenAPI specification
type DeleteNamespaceRequest struct {
	Id interface{} `json:"Id"`
}

// DiscoverInstancesResponse represents the DiscoverInstancesResponse schema from the OpenAPI specification
type DiscoverInstancesResponse struct {
	Instances interface{} `json:"Instances,omitempty"`
}

// ServiceFilter represents the ServiceFilter schema from the OpenAPI specification
type ServiceFilter struct {
	Values interface{} `json:"Values"`
	Condition interface{} `json:"Condition,omitempty"`
	Name interface{} `json:"Name"`
}

// CreateHttpNamespaceRequest represents the CreateHttpNamespaceRequest schema from the OpenAPI specification
type CreateHttpNamespaceRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name"`
}

// Service represents the Service schema from the OpenAPI specification
type Service struct {
	Instancecount interface{} `json:"InstanceCount,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Creatorrequestid interface{} `json:"CreatorRequestId,omitempty"`
	Healthcheckcustomconfig interface{} `json:"HealthCheckCustomConfig,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Dnsconfig interface{} `json:"DnsConfig,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Namespaceid interface{} `json:"NamespaceId,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Healthcheckconfig interface{} `json:"HealthCheckConfig,omitempty"`
}

// GetServiceResponse represents the GetServiceResponse schema from the OpenAPI specification
type GetServiceResponse struct {
	Service interface{} `json:"Service,omitempty"`
}

// NamespaceSummary represents the NamespaceSummary schema from the OpenAPI specification
type NamespaceSummary struct {
	Id interface{} `json:"Id,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
	Servicecount interface{} `json:"ServiceCount,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Createdate interface{} `json:"CreateDate,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
}

// UpdatePublicDnsNamespaceResponse represents the UpdatePublicDnsNamespaceResponse schema from the OpenAPI specification
type UpdatePublicDnsNamespaceResponse struct {
	Operationid interface{} `json:"OperationId,omitempty"`
}

// PrivateDnsPropertiesMutableChange represents the PrivateDnsPropertiesMutableChange schema from the OpenAPI specification
type PrivateDnsPropertiesMutableChange struct {
	Soa interface{} `json:"SOA"`
}

// ListServicesResponse represents the ListServicesResponse schema from the OpenAPI specification
type ListServicesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Services interface{} `json:"Services,omitempty"`
}
