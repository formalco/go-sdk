package sdk

import (
	"net/http"

	"buf.build/gen/go/formal/core/connectrpc/go/core/v1/corev1connect"
)

const (
	FORMAL_HOST_URL string = "https://v2api.formalcloud.net"
)

type FormalSDK struct {
	LogsServiceClient                   corev1connect.LogsServiceClient
	GroupServiceClient                  corev1connect.GroupServiceClient
	IntegrationBIServiceClient          corev1connect.IntegrationBIServiceClient
	IntegrationCloudServiceClient       corev1connect.IntegrationCloudServiceClient
	IntegrationDataCatalogServiceClient corev1connect.IntegrationDataCatalogServiceClient
	IntegrationsLogServiceClient        corev1connect.IntegrationsLogServiceClient
	IntegrationMfaServiceClient         corev1connect.IntegrationMfaServiceClient
	IntegrationMDMServiceClient         corev1connect.IntegrationMDMServiceClient
	InventoryServiceClient              corev1connect.InventoryServiceClient
	PoliciesServiceClient               corev1connect.PoliciesServiceClient
	ResourceServiceClient               corev1connect.ResourceServiceClient
	SatelliteServiceClient              corev1connect.SatelliteServiceClient
	SessionServiceClient                corev1connect.SessionsServiceClient
	SidecarServiceClient                corev1connect.SidecarServiceClient
	TrackersServiceClient               corev1connect.TrackersServiceClient
	UserServiceClient                   corev1connect.UserServiceClient
}

func New(apiKey string) *FormalSDK {
	httpClient := &http.Client{Transport: &transport{
		apiKey:              apiKey,
		underlyingTransport: http.DefaultTransport,
	}}
	return &FormalSDK{
		LogsServiceClient:                   corev1connect.NewLogsServiceClient(httpClient, FORMAL_HOST_URL),
		GroupServiceClient:                  corev1connect.NewGroupServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationBIServiceClient:          corev1connect.NewIntegrationBIServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationCloudServiceClient:       corev1connect.NewIntegrationCloudServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationDataCatalogServiceClient: corev1connect.NewIntegrationDataCatalogServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationMfaServiceClient:         corev1connect.NewIntegrationMfaServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationsLogServiceClient:        corev1connect.NewIntegrationsLogServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationMDMServiceClient:         corev1connect.NewIntegrationMDMServiceClient(httpClient, FORMAL_HOST_URL),
		InventoryServiceClient:              corev1connect.NewInventoryServiceClient(httpClient, FORMAL_HOST_URL),
		PoliciesServiceClient:               corev1connect.NewPoliciesServiceClient(httpClient, FORMAL_HOST_URL),
		ResourceServiceClient:               corev1connect.NewResourceServiceClient(httpClient, FORMAL_HOST_URL),
		TrackersServiceClient:               corev1connect.NewTrackersServiceClient(httpClient, FORMAL_HOST_URL),
		SatelliteServiceClient:              corev1connect.NewSatelliteServiceClient(httpClient, FORMAL_HOST_URL),
		SessionServiceClient:                corev1connect.NewSessionsServiceClient(httpClient, FORMAL_HOST_URL),
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
		LogsServiceClient:                   corev1connect.NewLogsServiceClient(httpClient, url),
		GroupServiceClient:                  corev1connect.NewGroupServiceClient(httpClient, url),
		IntegrationBIServiceClient:          corev1connect.NewIntegrationBIServiceClient(httpClient, url),
		IntegrationCloudServiceClient:       corev1connect.NewIntegrationCloudServiceClient(httpClient, url),
		IntegrationDataCatalogServiceClient: corev1connect.NewIntegrationDataCatalogServiceClient(httpClient, url),
		IntegrationMfaServiceClient:         corev1connect.NewIntegrationMfaServiceClient(httpClient, url),
		IntegrationsLogServiceClient:        corev1connect.NewIntegrationsLogServiceClient(httpClient, url),
		IntegrationMDMServiceClient:         corev1connect.NewIntegrationMDMServiceClient(httpClient, url),
		InventoryServiceClient:              corev1connect.NewInventoryServiceClient(httpClient, url),
		PoliciesServiceClient:               corev1connect.NewPoliciesServiceClient(httpClient, url),
		ResourceServiceClient:               corev1connect.NewResourceServiceClient(httpClient, url),
		TrackersServiceClient:               corev1connect.NewTrackersServiceClient(httpClient, url),
		SatelliteServiceClient:              corev1connect.NewSatelliteServiceClient(httpClient, url),
		SessionServiceClient:                corev1connect.NewSessionsServiceClient(httpClient, url),
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
