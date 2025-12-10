package sdk

import (
	"net/http"

	corev1connect "buf.build/gen/go/formal/core/connectrpc/go/core/v1/corev1connect"
)

const (
	FORMAL_HOST_URL string = "https://api.joinformal.com"
)

type FormalSDK struct {
	ConnectorServiceClient              corev1connect.ConnectorServiceClient
	GraphServiceClient                  corev1connect.GraphServiceClient
	GroupServiceClient                  corev1connect.GroupServiceClient
	IntegrationBIServiceClient          corev1connect.IntegrationBIServiceClient
	IntegrationCloudServiceClient       corev1connect.IntegrationCloudServiceClient
	IntegrationDataCatalogServiceClient corev1connect.IntegrationDataCatalogServiceClient
	IntegrationsLogServiceClient        corev1connect.IntegrationsLogServiceClient
	IntegrationMDMServiceClient         corev1connect.IntegrationMDMServiceClient
	InventoryServiceClient              corev1connect.InventoryServiceClient
	LogsServiceClient                   corev1connect.LogsServiceClient
	PermissionsServiceClient            corev1connect.PermissionsServiceClient
	PoliciesServiceClient               corev1connect.PoliciesServiceClient
	PolicyDataLoaderServiceClient       corev1connect.PolicyDataLoaderServiceClient
	ResourceServiceClient               corev1connect.ResourceServiceClient
	SatelliteServiceClient              corev1connect.SatelliteServiceClient
	ScenarioMonitoringServiceClient     corev1connect.ScenarioMonitoringServiceClient
	SessionsServiceClient               corev1connect.SessionsServiceClient
	SidecarServiceClient                corev1connect.SidecarServiceClient
	SpaceServiceClient                  corev1connect.SpaceServiceClient
	UserServiceClient                   corev1connect.UserServiceClient
	WorkflowServiceClient               corev1connect.WorkflowServiceClient
}

func New(apiKey string) *FormalSDK {
	httpClient := &http.Client{Transport: &transport{
		apiKey:              apiKey,
		underlyingTransport: http.DefaultTransport,
	}}
	return &FormalSDK{
		ConnectorServiceClient:              corev1connect.NewConnectorServiceClient(httpClient, FORMAL_HOST_URL),
		GraphServiceClient:                  corev1connect.NewGraphServiceClient(httpClient, FORMAL_HOST_URL),
		GroupServiceClient:                  corev1connect.NewGroupServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationBIServiceClient:          corev1connect.NewIntegrationBIServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationCloudServiceClient:       corev1connect.NewIntegrationCloudServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationDataCatalogServiceClient: corev1connect.NewIntegrationDataCatalogServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationMDMServiceClient:         corev1connect.NewIntegrationMDMServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationsLogServiceClient:        corev1connect.NewIntegrationsLogServiceClient(httpClient, FORMAL_HOST_URL),
		InventoryServiceClient:              corev1connect.NewInventoryServiceClient(httpClient, FORMAL_HOST_URL),
		LogsServiceClient:                   corev1connect.NewLogsServiceClient(httpClient, FORMAL_HOST_URL),
		PermissionsServiceClient:            corev1connect.NewPermissionsServiceClient(httpClient, FORMAL_HOST_URL),
		PoliciesServiceClient:               corev1connect.NewPoliciesServiceClient(httpClient, FORMAL_HOST_URL),
		PolicyDataLoaderServiceClient:       corev1connect.NewPolicyDataLoaderServiceClient(httpClient, FORMAL_HOST_URL),
		ResourceServiceClient:               corev1connect.NewResourceServiceClient(httpClient, FORMAL_HOST_URL),
		SatelliteServiceClient:              corev1connect.NewSatelliteServiceClient(httpClient, FORMAL_HOST_URL),
		ScenarioMonitoringServiceClient:     corev1connect.NewScenarioMonitoringServiceClient(httpClient, FORMAL_HOST_URL),
		SessionsServiceClient:               corev1connect.NewSessionsServiceClient(httpClient, FORMAL_HOST_URL),
		SidecarServiceClient:                corev1connect.NewSidecarServiceClient(httpClient, FORMAL_HOST_URL),
		SpaceServiceClient:                  corev1connect.NewSpaceServiceClient(httpClient, FORMAL_HOST_URL),
		UserServiceClient:                   corev1connect.NewUserServiceClient(httpClient, FORMAL_HOST_URL),
		WorkflowServiceClient:               corev1connect.NewWorkflowServiceClient(httpClient, FORMAL_HOST_URL),
	}
}

func NewWithUrl(apiKey string, url string) *FormalSDK {
	httpClient := &http.Client{Transport: &transport{
		apiKey:              apiKey,
		underlyingTransport: http.DefaultTransport,
	}}
	return &FormalSDK{
		ConnectorServiceClient:              corev1connect.NewConnectorServiceClient(httpClient, url),
		GraphServiceClient:                  corev1connect.NewGraphServiceClient(httpClient, url),
		GroupServiceClient:                  corev1connect.NewGroupServiceClient(httpClient, url),
		IntegrationBIServiceClient:          corev1connect.NewIntegrationBIServiceClient(httpClient, url),
		IntegrationCloudServiceClient:       corev1connect.NewIntegrationCloudServiceClient(httpClient, url),
		IntegrationDataCatalogServiceClient: corev1connect.NewIntegrationDataCatalogServiceClient(httpClient, url),
		IntegrationMDMServiceClient:         corev1connect.NewIntegrationMDMServiceClient(httpClient, url),
		IntegrationsLogServiceClient:        corev1connect.NewIntegrationsLogServiceClient(httpClient, url),
		InventoryServiceClient:              corev1connect.NewInventoryServiceClient(httpClient, url),
		LogsServiceClient:                   corev1connect.NewLogsServiceClient(httpClient, url),
		PermissionsServiceClient:            corev1connect.NewPermissionsServiceClient(httpClient, url),
		PoliciesServiceClient:               corev1connect.NewPoliciesServiceClient(httpClient, url),
		PolicyDataLoaderServiceClient:       corev1connect.NewPolicyDataLoaderServiceClient(httpClient, url),
		ResourceServiceClient:               corev1connect.NewResourceServiceClient(httpClient, url),
		SatelliteServiceClient:              corev1connect.NewSatelliteServiceClient(httpClient, url),
		ScenarioMonitoringServiceClient:     corev1connect.NewScenarioMonitoringServiceClient(httpClient, url),
		SessionsServiceClient:               corev1connect.NewSessionsServiceClient(httpClient, url),
		SidecarServiceClient:                corev1connect.NewSidecarServiceClient(httpClient, url),
		SpaceServiceClient:                  corev1connect.NewSpaceServiceClient(httpClient, url),
		UserServiceClient:                   corev1connect.NewUserServiceClient(httpClient, url),
		WorkflowServiceClient:               corev1connect.NewWorkflowServiceClient(httpClient, url),
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
