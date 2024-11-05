package sdk

import (
	"buf.build/gen/go/formal/core/connectrpc/go/core/v1/corev1connect"
)

const FormalHostUrl string = "https://v2api.formalcloud.net"

type FormalSDK struct {
	ConnectorServiceClient              corev1connect.ConnectorServiceClient
	GroupServiceClient                  corev1connect.GroupServiceClient
	IntegrationBIServiceClient          corev1connect.IntegrationBIServiceClient
	IntegrationCloudServiceClient       corev1connect.IntegrationCloudServiceClient
	IntegrationDataCatalogServiceClient corev1connect.IntegrationDataCatalogServiceClient
	IntegrationsLogServiceClient        corev1connect.IntegrationsLogServiceClient
	IntegrationMfaServiceClient         corev1connect.IntegrationMfaServiceClient
	IntegrationMDMServiceClient         corev1connect.IntegrationMDMServiceClient
	InventoryServiceClient              corev1connect.InventoryServiceClient
	LogsServiceClient                   corev1connect.LogsServiceClient
	MonitorsServiceClient               corev1connect.MonitorsServiceClient
	PoliciesServiceClient               corev1connect.PoliciesServiceClient
	PolicyDataLoaderServiceClient       corev1connect.PolicyDataLoaderServiceClient
	ResourceServiceClient               corev1connect.ResourceServiceClient
	SatelliteServiceClient              corev1connect.SatelliteServiceClient
	SessionServiceClient                corev1connect.SessionsServiceClient
	SidecarServiceClient                corev1connect.SidecarServiceClient
	SpaceServiceClient                  corev1connect.SpaceServiceClient
	TrackersServiceClient               corev1connect.TrackersServiceClient
	UserServiceClient                   corev1connect.UserServiceClient
}

// New creates a new FormalSDK instance
func New(apiKey string) *FormalSDK {
	return NewWithUrl(apiKey, FormalHostUrl)
}

// NewWithUrl creates a new FormalSDK instance with a custom URL
func NewWithUrl(apiKey string, url string) *FormalSDK {
	httpClient := NewClient(apiKey)

	return &FormalSDK{
		ConnectorServiceClient:              corev1connect.NewConnectorServiceClient(httpClient, url),
		GroupServiceClient:                  corev1connect.NewGroupServiceClient(httpClient, url),
		IntegrationBIServiceClient:          corev1connect.NewIntegrationBIServiceClient(httpClient, url),
		IntegrationCloudServiceClient:       corev1connect.NewIntegrationCloudServiceClient(httpClient, url),
		IntegrationDataCatalogServiceClient: corev1connect.NewIntegrationDataCatalogServiceClient(httpClient, url),
		IntegrationMfaServiceClient:         corev1connect.NewIntegrationMfaServiceClient(httpClient, url),
		IntegrationsLogServiceClient:        corev1connect.NewIntegrationsLogServiceClient(httpClient, url),
		IntegrationMDMServiceClient:         corev1connect.NewIntegrationMDMServiceClient(httpClient, url),
		InventoryServiceClient:              corev1connect.NewInventoryServiceClient(httpClient, url),
		LogsServiceClient:                   corev1connect.NewLogsServiceClient(httpClient, url),
		MonitorsServiceClient:               corev1connect.NewMonitorsServiceClient(httpClient, url),
		PoliciesServiceClient:               corev1connect.NewPoliciesServiceClient(httpClient, url),
		PolicyDataLoaderServiceClient:       corev1connect.NewPolicyDataLoaderServiceClient(httpClient, url),
		ResourceServiceClient:               corev1connect.NewResourceServiceClient(httpClient, url),
		SatelliteServiceClient:              corev1connect.NewSatelliteServiceClient(httpClient, url),
		SessionServiceClient:                corev1connect.NewSessionsServiceClient(httpClient, url),
		SidecarServiceClient:                corev1connect.NewSidecarServiceClient(httpClient, url),
		SpaceServiceClient:                  corev1connect.NewSpaceServiceClient(httpClient, url),
		TrackersServiceClient:               corev1connect.NewTrackersServiceClient(httpClient, url),
		UserServiceClient:                   corev1connect.NewUserServiceClient(httpClient, url),
	}
}
