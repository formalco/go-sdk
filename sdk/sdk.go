package sdk

import (
	"net/http"

	coreconnect "buf.build/gen/go/formal/core/connectrpc/go/core/v1/corev1connect"
)

const (
	FORMAL_HOST_URL string = "https://adminv2.api.formalcloud.net"
)

type FormalSDK struct {
	AuditService                  coreconnect.AuditServiceClient
	ResourceService               coreconnect.ResourceServiceClient
	UserService                   coreconnect.UserServiceClient
	GroupService                  coreconnect.GroupServiceClient
	IntegrationCloudService       coreconnect.IntegrationCloudServiceClient
	IntegrationDataCatalogService coreconnect.IntegrationDataCatalogServiceClient
	IntegrationMfaService         coreconnect.IntegrationMfaServiceClient
	IntegrationBIService          coreconnect.IntegrationBIServiceClient
	IntegrationLogService         coreconnect.IntegrationsLogServiceClient
	InventoryService              coreconnect.InventoryServiceClient
	PolicyService                 coreconnect.PolicyServiceClient
	RowLevelTrackerService        coreconnect.RowLevelTrackerServiceClient
	SatelliteService              coreconnect.SatelliteServiceClient
	SidecarService                coreconnect.SidecarServiceClient
	SessionService                coreconnect.SessionServiceClient
}

func New(apiKey string) *FormalSDK {
	httpClient := &http.Client{Transport: &transport{
		apiKey:              apiKey,
		underlyingTransport: http.DefaultTransport,
	}}
	return &FormalSDK{
		AuditService:                  coreconnect.NewAuditServiceClient(httpClient, FORMAL_HOST_URL),
		ResourceService:               coreconnect.NewResourceServiceClient(httpClient, FORMAL_HOST_URL),
		GroupService:                  coreconnect.NewGroupServiceClient(httpClient, FORMAL_HOST_URL),
		UserService:                   coreconnect.NewUserServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationCloudService:       coreconnect.NewIntegrationCloudServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationMfaService:         coreconnect.NewIntegrationMfaServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationDataCatalogService: coreconnect.NewIntegrationDataCatalogServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationBIService:          coreconnect.NewIntegrationBIServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationLogService:         coreconnect.NewIntegrationsLogServiceClient(httpClient, FORMAL_HOST_URL),
		InventoryService:              coreconnect.NewInventoryServiceClient(httpClient, FORMAL_HOST_URL),
		PolicyService:                 coreconnect.NewPolicyServiceClient(httpClient, FORMAL_HOST_URL),
		RowLevelTrackerService:        coreconnect.NewRowLevelTrackerServiceClient(httpClient, FORMAL_HOST_URL),
		SatelliteService:              coreconnect.NewSatelliteServiceClient(httpClient, FORMAL_HOST_URL),
		SidecarService:                coreconnect.NewSidecarServiceClient(httpClient, FORMAL_HOST_URL),
	}
}

func NewWithUrl(apiKey string, url string) *FormalSDK {
	httpClient := &http.Client{Transport: &transport{
		apiKey:              apiKey,
		underlyingTransport: http.DefaultTransport,
	}}
	return &FormalSDK{
		AuditService:                  coreconnect.NewAuditServiceClient(httpClient, FORMAL_HOST_URL),
		ResourceService:               coreconnect.NewResourceServiceClient(httpClient, FORMAL_HOST_URL),
		GroupService:                  coreconnect.NewGroupServiceClient(httpClient, FORMAL_HOST_URL),
		UserService:                   coreconnect.NewUserServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationCloudService:       coreconnect.NewIntegrationCloudServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationMfaService:         coreconnect.NewIntegrationMfaServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationDataCatalogService: coreconnect.NewIntegrationDataCatalogServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationBIService:          coreconnect.NewIntegrationBIServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationLogService:         coreconnect.NewIntegrationsLogServiceClient(httpClient, FORMAL_HOST_URL),
		InventoryService:              coreconnect.NewInventoryServiceClient(httpClient, FORMAL_HOST_URL),
		PolicyService:                 coreconnect.NewPolicyServiceClient(httpClient, FORMAL_HOST_URL),
		RowLevelTrackerService:        coreconnect.NewRowLevelTrackerServiceClient(httpClient, FORMAL_HOST_URL),
		SatelliteService:              coreconnect.NewSatelliteServiceClient(httpClient, FORMAL_HOST_URL),
		SidecarService:                coreconnect.NewSidecarServiceClient(httpClient, FORMAL_HOST_URL),
	}
}

type transport struct {
	underlyingTransport http.RoundTripper
	apiKey              string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Api-Key", t.apiKey)
	return t.underlyingTransport.RoundTrip(req)
}
