package sdk

import (
	"net/http"

	"buf.build/gen/go/formal/core/connectrpc/go/core/v1/corev1connect"
)

const FormalHostUrl string = "https://v2api.formalcloud.net"

type FormalSDK struct {
	GroupServiceClient                  corev1connect.GroupServiceClient
	IntegrationBIServiceClient          corev1connect.IntegrationBIServiceClient
	IntegrationCloudServiceClient       corev1connect.IntegrationCloudServiceClient
	IntegrationDataCatalogServiceClient corev1connect.IntegrationDataCatalogServiceClient
	IntegrationsLogServiceClient        corev1connect.IntegrationsLogServiceClient
	IntegrationMfaServiceClient         corev1connect.IntegrationMfaServiceClient
	IntegrationMDMServiceClient         corev1connect.IntegrationMDMServiceClient
	InventoryServiceClient              corev1connect.InventoryServiceClient
	ListenersServiceClient              corev1connect.ListenerServiceClient
	LogsServiceClient                   corev1connect.LogsServiceClient
	PoliciesServiceClient               corev1connect.PoliciesServiceClient
	ResourceServiceClient               corev1connect.ResourceServiceClient
	SatelliteServiceClient              corev1connect.SatelliteServiceClient
	SessionServiceClient                corev1connect.SessionsServiceClient
	SidecarServiceClient                corev1connect.SidecarServiceClient
	TrackersServiceClient               corev1connect.TrackersServiceClient
	UserServiceClient                   corev1connect.UserServiceClient
}

// New creates a new FormalSDK instance
func New(apiKey string) *FormalSDK {
	httpClient := &http.Client{
		Transport: &transport{
			apiKey:              apiKey,
			underlyingTransport: http.DefaultTransport,
		},
	}

	return &FormalSDK{
		GroupServiceClient:                  corev1connect.NewGroupServiceClient(httpClient, FormalHostUrl),
		IntegrationBIServiceClient:          corev1connect.NewIntegrationBIServiceClient(httpClient, FormalHostUrl),
		IntegrationCloudServiceClient:       corev1connect.NewIntegrationCloudServiceClient(httpClient, FormalHostUrl),
		IntegrationDataCatalogServiceClient: corev1connect.NewIntegrationDataCatalogServiceClient(httpClient, FormalHostUrl),
		IntegrationMfaServiceClient:         corev1connect.NewIntegrationMfaServiceClient(httpClient, FormalHostUrl),
		IntegrationsLogServiceClient:        corev1connect.NewIntegrationsLogServiceClient(httpClient, FormalHostUrl),
		IntegrationMDMServiceClient:         corev1connect.NewIntegrationMDMServiceClient(httpClient, FormalHostUrl),
		InventoryServiceClient:              corev1connect.NewInventoryServiceClient(httpClient, FormalHostUrl),
		ListenersServiceClient:              corev1connect.NewListenerServiceClient(httpClient, FormalHostUrl),
		LogsServiceClient:                   corev1connect.NewLogsServiceClient(httpClient, FormalHostUrl),
		PoliciesServiceClient:               corev1connect.NewPoliciesServiceClient(httpClient, FormalHostUrl),
		ResourceServiceClient:               corev1connect.NewResourceServiceClient(httpClient, FormalHostUrl),
		SatelliteServiceClient:              corev1connect.NewSatelliteServiceClient(httpClient, FormalHostUrl),
		SessionServiceClient:                corev1connect.NewSessionsServiceClient(httpClient, FormalHostUrl),
		SidecarServiceClient:                corev1connect.NewSidecarServiceClient(httpClient, FormalHostUrl),
		TrackersServiceClient:               corev1connect.NewTrackersServiceClient(httpClient, FormalHostUrl),
		UserServiceClient:                   corev1connect.NewUserServiceClient(httpClient, FormalHostUrl),
	}
}

// NewWithUrl creates a new FormalSDK instance with a custom URL
func NewWithUrl(apiKey string, url string) *FormalSDK {
	httpClient := &http.Client{
		Transport: &transport{
			apiKey:              apiKey,
			underlyingTransport: http.DefaultTransport,
		},
	}

	return &FormalSDK{
		GroupServiceClient:                  corev1connect.NewGroupServiceClient(httpClient, url),
		IntegrationBIServiceClient:          corev1connect.NewIntegrationBIServiceClient(httpClient, url),
		IntegrationCloudServiceClient:       corev1connect.NewIntegrationCloudServiceClient(httpClient, url),
		IntegrationDataCatalogServiceClient: corev1connect.NewIntegrationDataCatalogServiceClient(httpClient, url),
		IntegrationMfaServiceClient:         corev1connect.NewIntegrationMfaServiceClient(httpClient, url),
		IntegrationsLogServiceClient:        corev1connect.NewIntegrationsLogServiceClient(httpClient, url),
		IntegrationMDMServiceClient:         corev1connect.NewIntegrationMDMServiceClient(httpClient, url),
		InventoryServiceClient:              corev1connect.NewInventoryServiceClient(httpClient, url),
		ListenersServiceClient:              corev1connect.NewListenerServiceClient(httpClient, url),
		LogsServiceClient:                   corev1connect.NewLogsServiceClient(httpClient, url),
		PoliciesServiceClient:               corev1connect.NewPoliciesServiceClient(httpClient, url),
		ResourceServiceClient:               corev1connect.NewResourceServiceClient(httpClient, url),
		SatelliteServiceClient:              corev1connect.NewSatelliteServiceClient(httpClient, url),
		SessionServiceClient:                corev1connect.NewSessionsServiceClient(httpClient, url),
		SidecarServiceClient:                corev1connect.NewSidecarServiceClient(httpClient, url),
		TrackersServiceClient:               corev1connect.NewTrackersServiceClient(httpClient, url),
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
