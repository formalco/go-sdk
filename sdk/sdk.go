package sdk

import (
	adminv1connect "buf.build/gen/go/formal/admin/bufbuild/connect-go/admin/v1/adminv1connect"
	"net/http"
)

const (
	FORMAL_HOST_URL string = "https://adminv2.api.formalcloud.net"
)

type FormalSDK struct {
	DevServiceClient                    adminv1connect.DevServiceClient
	AuditLogsServiceClient              adminv1connect.AuditLogsServiceClient
	EncryptionServiceClient             adminv1connect.EncryptionServiceClient
	RoleServiceClient                   adminv1connect.RoleServiceClient
	GroupServiceClient                  adminv1connect.GroupServiceClient
	SlackServiceClient                  adminv1connect.SlackServiceClient
	CloudServiceClient                  adminv1connect.CloudServiceClient
	KMSServiceClient                    adminv1connect.KMSServiceClient
	GithubServiceClient                 adminv1connect.GithubServiceClient
	CodeRepositoryServiceClient         adminv1connect.CodeRepositoryServiceClient
	IncidentServiceClient               adminv1connect.IncidentServiceClient
	SsoServiceClient                    adminv1connect.SsoServiceClient
	AppServiceClient                    adminv1connect.AppServiceClient
	LogsServiceClient                   adminv1connect.LogsServiceClient
	IntegrationExternalAPIServiceClient adminv1connect.IntegrationExternalAPIServiceClient
	InventoryServiceClient              adminv1connect.InventoryServiceClient
	MetricsServiceClient                adminv1connect.MetricsServiceClient
	PolicyServiceClient                 adminv1connect.PolicyServiceClient
	RegistryServiceClient               adminv1connect.RegistryServiceClient
	SidecarServiceClient                adminv1connect.SidecarServiceClient
	DSyncClient                         adminv1connect.DSyncClient
}

func NewFormalSDK(apiKey string) *FormalSDK {
	httpClient := &http.Client{Transport: &transport{
		apiKey:              apiKey,
		underlyingTransport: http.DefaultTransport,
	}}
	sdk := &FormalSDK{
		AppServiceClient:                    adminv1connect.NewAppServiceClient(httpClient, FORMAL_HOST_URL),
		AuditLogsServiceClient:              adminv1connect.NewAuditLogsServiceClient(httpClient, FORMAL_HOST_URL),
		CloudServiceClient:                  adminv1connect.NewCloudServiceClient(httpClient, FORMAL_HOST_URL),
		CodeRepositoryServiceClient:         adminv1connect.NewCodeRepositoryServiceClient(httpClient, FORMAL_HOST_URL),
		DSyncClient:                         adminv1connect.NewDSyncClient(httpClient, FORMAL_HOST_URL),
		DevServiceClient:                    adminv1connect.NewDevServiceClient(httpClient, FORMAL_HOST_URL),
		EncryptionServiceClient:             adminv1connect.NewEncryptionServiceClient(httpClient, FORMAL_HOST_URL),
		GithubServiceClient:                 adminv1connect.NewGithubServiceClient(httpClient, FORMAL_HOST_URL),
		GroupServiceClient:                  adminv1connect.NewGroupServiceClient(httpClient, FORMAL_HOST_URL),
		IncidentServiceClient:               adminv1connect.NewIncidentServiceClient(httpClient, FORMAL_HOST_URL),
		IntegrationExternalAPIServiceClient: adminv1connect.NewIntegrationExternalAPIServiceClient(httpClient, FORMAL_HOST_URL),
		InventoryServiceClient:              adminv1connect.NewInventoryServiceClient(httpClient, FORMAL_HOST_URL),
		KMSServiceClient:                    adminv1connect.NewKMSServiceClient(httpClient, FORMAL_HOST_URL),
		LogsServiceClient:                   adminv1connect.NewLogsServiceClient(httpClient, FORMAL_HOST_URL),
		MetricsServiceClient:                adminv1connect.NewMetricsServiceClient(httpClient, FORMAL_HOST_URL),
		PolicyServiceClient:                 adminv1connect.NewPolicyServiceClient(httpClient, FORMAL_HOST_URL),
		RegistryServiceClient:               adminv1connect.NewRegistryServiceClient(httpClient, FORMAL_HOST_URL),
		RoleServiceClient:                   adminv1connect.NewRoleServiceClient(httpClient, FORMAL_HOST_URL),
		SidecarServiceClient:                adminv1connect.NewSidecarServiceClient(httpClient, FORMAL_HOST_URL),
		SlackServiceClient:                  adminv1connect.NewSlackServiceClient(httpClient, FORMAL_HOST_URL),
		SsoServiceClient:                    adminv1connect.NewSsoServiceClient(httpClient, FORMAL_HOST_URL),
	}
	return sdk
}

type transport struct {
	underlyingTransport http.RoundTripper
	apiKey              string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Api-Key", t.apiKey)
	return t.underlyingTransport.RoundTrip(req)
}
