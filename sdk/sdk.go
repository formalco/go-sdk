package sdk

import (
	"net/http"

	corev1connect "buf.build/gen/go/formal/core/connectrpc/go/core/v1/corev1connect"
)

const (
	FORMAL_HOST_URL string = "https://v2api.formalcloud.net"
)

type FormalSDK struct {
	AuditServiceClient                  corev1connect.AuditServiceClient
	GroupServiceClient                  corev1connect.GroupServiceClient
	IntegrationBIServiceClient          corev1connect.IntegrationBIServiceClient
	IntegrationCloudServiceClient       corev1connect.IntegrationCloudServiceClient
	IntegrationDataCatalogServiceClient corev1connect.IntegrationDataCatalogServiceClient
	IntegrationsLogServiceClient        corev1connect.IntegrationsLogServiceClient
	IntegrationMfaServiceClient         corev1connect.IntegrationMfaServiceClient
	InventoryServiceClient              corev1connect.InventoryServiceClient
	PolicyServiceClient                 corev1connect.PolicyServiceClient
	ResourceServiceClient               corev1connect.ResourceServiceClient
	SatelliteServiceClient              corev1connect.SatelliteServiceClient
	SessionServiceClient                corev1connect.SessionServiceClient
	SidecarServiceClient                corev1connect.SidecarServiceClient
	RowLevelTrackerServiceClient        corev1connect.RowLevelTrackerServiceClient
	UserServiceClient                   corev1connect.UserServiceClient
}

func New(apiKey string) *FormalSDK {
	httpClient := &http.Client{Transport: &transport{
		apiKey:              apiKey,
		underlyingTransport: http.DefaultTransport,
	}}
	return &FormalSDK{
		AuditServiceClient:                  corev1connect.NewAuditServiceClient(httpClient, FORMAL_HOST_URL),
		GroupServiceClient:                  corev1connect.NewGroupServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationBIServiceClient:          corev1connect.NewIntegrationBIServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationCloudServiceClient:       corev1connect.NewIntegrationCloudServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationDataCatalogServiceClient: corev1connect.NewIntegrationDataCatalogServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationMfaServiceClient:         corev1connect.NewIntegrationMfaServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationsLogServiceClient:        corev1connect.NewIntegrationsLogServiceClient(httpClient, FORMAL_HOST_URL),
		InventoryServiceClient:              corev1connect.NewInventoryServiceClient(httpClient, FORMAL_HOST_URL),
		PolicyServiceClient:                 corev1connect.NewPolicyServiceClient(httpClient, FORMAL_HOST_URL),
		ResourceServiceClient:               corev1connect.NewResourceServiceClient(httpClient, FORMAL_HOST_URL),
		RowLevelTrackerServiceClient:        corev1connect.NewRowLevelTrackerServiceClient(httpClient, FORMAL_HOST_URL),
		SatelliteServiceClient:              corev1connect.NewSatelliteServiceClient(httpClient, FORMAL_HOST_URL),
		SessionServiceClient:                corev1connect.NewSessionServiceClient(httpClient, FORMAL_HOST_URL),
		SidecarServiceClient:                corev1connect.NewSidecarServiceClient(httpClient, FORMAL_HOST_URL),
		UserServiceClient:                   corev1connect.NewUserServiceClient(httpClient, FORMAL_HOST_URL),
	}
}

func NewWithUrl(apiKey string, url string) *FormalSDK {
	httpClient := &http.Client{Transport: &transport{
		apiKey:              apiKey,
		underlyingTransport: http.DefaultTransport,
	}}
	return &FormalSDK{
		AuditServiceClient:                  corev1connect.NewAuditServiceClient(httpClient, url),
		GroupServiceClient:                  corev1connect.NewGroupServiceClient(httpClient, url),
		IntegrationBIServiceClient:          corev1connect.NewIntegrationBIServiceClient(httpClient, url),
		IntegrationCloudServiceClient:       corev1connect.NewIntegrationCloudServiceClient(httpClient, url),
		IntegrationDataCatalogServiceClient: corev1connect.NewIntegrationDataCatalogServiceClient(httpClient, url),
		IntegrationMfaServiceClient:         corev1connect.NewIntegrationMfaServiceClient(httpClient, url),
		IntegrationsLogServiceClient:        corev1connect.NewIntegrationsLogServiceClient(httpClient, url),
		InventoryServiceClient:              corev1connect.NewInventoryServiceClient(httpClient, url),
		PolicyServiceClient:                 corev1connect.NewPolicyServiceClient(httpClient, url),
		ResourceServiceClient:               corev1connect.NewResourceServiceClient(httpClient, url),
		RowLevelTrackerServiceClient:        corev1connect.NewRowLevelTrackerServiceClient(httpClient, url),
		SatelliteServiceClient:              corev1connect.NewSatelliteServiceClient(httpClient, url),
		SessionServiceClient:                corev1connect.NewSessionServiceClient(httpClient, url),
		SidecarServiceClient:                corev1connect.NewSidecarServiceClient(httpClient, url),
		UserServiceClient:                   corev1connect.NewUserServiceClient(httpClient, url),
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
