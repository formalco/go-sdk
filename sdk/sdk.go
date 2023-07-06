package sdk

import (
	adminv1connect "buf.build/gen/go/formal/admin/bufbuild/connect-go/admin/v1/adminv1connect"
	"net/http"
)

const (
	FORMAL_HOST_URL string = "https://adminv2.api.formalcloud.net"
)

type FormalSDK struct {
	DevServiceClient                   adminv1connect.DevServiceClient
	AuditLogsServiceClient             adminv1connect.AuditLogsServiceClient
	CordServiceClient                  adminv1connect.CordServiceClient
	DataStoreServiceClient             adminv1connect.DataStoreServiceClient
	FieldEncryptionPolicyServiceClient adminv1connect.FieldEncryptionPolicyServiceClient
	FieldEncryptionServiceClient       adminv1connect.FieldEncryptionServiceClient
	UserServiceClient                  adminv1connect.UserServiceClient
	GroupServiceClient                 adminv1connect.GroupServiceClient
	SlackServiceClient                 adminv1connect.SlackServiceClient
	CloudServiceClient                 adminv1connect.CloudServiceClient
	KmsServiceClient                   adminv1connect.KmsServiceClient
	GithubServiceClient                adminv1connect.GithubServiceClient
	CodeRepositoryServiceClient        adminv1connect.CodeRepositoryServiceClient
	IncidentServiceClient              adminv1connect.IncidentServiceClient
	SsoServiceClient                   adminv1connect.SsoServiceClient
	AppServiceClient                   adminv1connect.AppServiceClient
	LogsServiceClient                  adminv1connect.LogsServiceClient
	ExternalApiServiceClient           adminv1connect.ExternalApiServiceClient
	InventoryServiceClient             adminv1connect.InventoryServiceClient
	MetricsServiceClient               adminv1connect.MetricsServiceClient
	NativeUserServiceClient            adminv1connect.NativeUserServiceClient
	OutputsServiceClient               adminv1connect.OutputsServiceClient
	PolicyServiceClient                adminv1connect.PolicyServiceClient
	RegistryServiceClient              adminv1connect.RegistryServiceClient
	SidecarServiceClient               adminv1connect.SidecarServiceClient
	DSyncClient                        adminv1connect.DSyncClient
}

func NewFormalSDK(apiKey string) *FormalSDK {
	httpClient := &http.Client{Transport: &transport{
		apiKey:              apiKey,
		underlyingTransport: http.DefaultTransport,
	}}
	return &FormalSDK{
		AppServiceClient:                   adminv1connect.NewAppServiceClient(httpClient, FORMAL_HOST_URL),
		AuditLogsServiceClient:             adminv1connect.NewAuditLogsServiceClient(httpClient, FORMAL_HOST_URL),
		CloudServiceClient:                 adminv1connect.NewCloudServiceClient(httpClient, FORMAL_HOST_URL),
		CodeRepositoryServiceClient:        adminv1connect.NewCodeRepositoryServiceClient(httpClient, FORMAL_HOST_URL),
		CordServiceClient:                  adminv1connect.NewCordServiceClient(httpClient, FORMAL_HOST_URL),
		DSyncClient:                        adminv1connect.NewDSyncClient(httpClient, FORMAL_HOST_URL),
		DataStoreServiceClient:             adminv1connect.NewDataStoreServiceClient(httpClient, FORMAL_HOST_URL),
		DevServiceClient:                   adminv1connect.NewDevServiceClient(httpClient, FORMAL_HOST_URL),
		ExternalApiServiceClient:           adminv1connect.NewExternalApiServiceClient(httpClient, FORMAL_HOST_URL),
		FieldEncryptionPolicyServiceClient: adminv1connect.NewFieldEncryptionPolicyServiceClient(httpClient, FORMAL_HOST_URL),
		FieldEncryptionServiceClient:       adminv1connect.NewFieldEncryptionServiceClient(httpClient, FORMAL_HOST_URL),
		GithubServiceClient:                adminv1connect.NewGithubServiceClient(httpClient, FORMAL_HOST_URL),
		GroupServiceClient:                 adminv1connect.NewGroupServiceClient(httpClient, FORMAL_HOST_URL),
		IncidentServiceClient:              adminv1connect.NewIncidentServiceClient(httpClient, FORMAL_HOST_URL),
		InventoryServiceClient:             adminv1connect.NewInventoryServiceClient(httpClient, FORMAL_HOST_URL),
		KmsServiceClient:                   adminv1connect.NewKmsServiceClient(httpClient, FORMAL_HOST_URL),
		LogsServiceClient:                  adminv1connect.NewLogsServiceClient(httpClient, FORMAL_HOST_URL),
		MetricsServiceClient:               adminv1connect.NewMetricsServiceClient(httpClient, FORMAL_HOST_URL),
		NativeUserServiceClient:            adminv1connect.NewNativeUserServiceClient(httpClient, FORMAL_HOST_URL),
		OutputsServiceClient:               adminv1connect.NewOutputsServiceClient(httpClient, FORMAL_HOST_URL),
		PolicyServiceClient:                adminv1connect.NewPolicyServiceClient(httpClient, FORMAL_HOST_URL),
		RegistryServiceClient:              adminv1connect.NewRegistryServiceClient(httpClient, FORMAL_HOST_URL),
		SidecarServiceClient:               adminv1connect.NewSidecarServiceClient(httpClient, FORMAL_HOST_URL),
		SlackServiceClient:                 adminv1connect.NewSlackServiceClient(httpClient, FORMAL_HOST_URL),
		SsoServiceClient:                   adminv1connect.NewSsoServiceClient(httpClient, FORMAL_HOST_URL),
		UserServiceClient:                  adminv1connect.NewUserServiceClient(httpClient, FORMAL_HOST_URL),
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
