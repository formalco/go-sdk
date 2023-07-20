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
	AppServiceClient                   adminv1connect.AppServiceClient
	CloudServiceClient                 adminv1connect.CloudServiceClient
	CodeRepositoryServiceClient        adminv1connect.CodeRepositoryServiceClient
	ExternalApiServiceClient           adminv1connect.ExternalApiServiceClient
	GithubServiceClient                adminv1connect.GithubServiceClient
	IncidentServiceClient              adminv1connect.IncidentServiceClient
	KmsServiceClient                   adminv1connect.KmsServiceClient
	LogsServiceClient                  adminv1connect.LogsServiceClient
	SlackServiceClient                 adminv1connect.SlackServiceClient
	SsoServiceClient                   adminv1connect.SsoServiceClient
	InventoryServiceClient             adminv1connect.InventoryServiceClient
	MetricsServiceClient               adminv1connect.MetricsServiceClient
	NativeUserServiceClient            adminv1connect.NativeUserServiceClient
	OutputsServiceClient               adminv1connect.OutputsServiceClient
	PolicyServiceClient                adminv1connect.PolicyServiceClient
	RegistryServiceClient              adminv1connect.RegistryServiceClient
	SearchServiceClient                adminv1connect.SearchServiceClient
	SidecarServiceClient               adminv1connect.SidecarServiceClient
	DSyncServiceClient                 adminv1connect.DSyncServiceClient
}

func New(apiKey string) *FormalSDK {
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
		DSyncServiceClient:                 adminv1connect.NewDSyncServiceClient(httpClient, FORMAL_HOST_URL),
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
		SearchServiceClient:                adminv1connect.NewSearchServiceClient(httpClient, FORMAL_HOST_URL),
		SidecarServiceClient:               adminv1connect.NewSidecarServiceClient(httpClient, FORMAL_HOST_URL),
		SlackServiceClient:                 adminv1connect.NewSlackServiceClient(httpClient, FORMAL_HOST_URL),
		SsoServiceClient:                   adminv1connect.NewSsoServiceClient(httpClient, FORMAL_HOST_URL),
		UserServiceClient:                  adminv1connect.NewUserServiceClient(httpClient, FORMAL_HOST_URL),
	}
}

func NewWithUrl(apiKey string, url string) *FormalSDK {
	httpClient := &http.Client{Transport: &transport{
		apiKey:              apiKey,
		underlyingTransport: http.DefaultTransport,
	}}
	return &FormalSDK{
		AppServiceClient:                   adminv1connect.NewAppServiceClient(httpClient, url),
		AuditLogsServiceClient:             adminv1connect.NewAuditLogsServiceClient(httpClient, url),
		CloudServiceClient:                 adminv1connect.NewCloudServiceClient(httpClient, url),
		CodeRepositoryServiceClient:        adminv1connect.NewCodeRepositoryServiceClient(httpClient, url),
		CordServiceClient:                  adminv1connect.NewCordServiceClient(httpClient, url),
		DSyncServiceClient:                 adminv1connect.NewDSyncServiceClient(httpClient, url),
		DataStoreServiceClient:             adminv1connect.NewDataStoreServiceClient(httpClient, url),
		DevServiceClient:                   adminv1connect.NewDevServiceClient(httpClient, url),
		ExternalApiServiceClient:           adminv1connect.NewExternalApiServiceClient(httpClient, url),
		FieldEncryptionPolicyServiceClient: adminv1connect.NewFieldEncryptionPolicyServiceClient(httpClient, url),
		FieldEncryptionServiceClient:       adminv1connect.NewFieldEncryptionServiceClient(httpClient, url),
		GithubServiceClient:                adminv1connect.NewGithubServiceClient(httpClient, url),
		GroupServiceClient:                 adminv1connect.NewGroupServiceClient(httpClient, url),
		IncidentServiceClient:              adminv1connect.NewIncidentServiceClient(httpClient, url),
		InventoryServiceClient:             adminv1connect.NewInventoryServiceClient(httpClient, url),
		KmsServiceClient:                   adminv1connect.NewKmsServiceClient(httpClient, url),
		LogsServiceClient:                  adminv1connect.NewLogsServiceClient(httpClient, url),
		MetricsServiceClient:               adminv1connect.NewMetricsServiceClient(httpClient, url),
		NativeUserServiceClient:            adminv1connect.NewNativeUserServiceClient(httpClient, url),
		OutputsServiceClient:               adminv1connect.NewOutputsServiceClient(httpClient, url),
		PolicyServiceClient:                adminv1connect.NewPolicyServiceClient(httpClient, url),
		RegistryServiceClient:              adminv1connect.NewRegistryServiceClient(httpClient, url),
		SearchServiceClient:                adminv1connect.NewSearchServiceClient(httpClient, url),
		SidecarServiceClient:               adminv1connect.NewSidecarServiceClient(httpClient, url),
		SlackServiceClient:                 adminv1connect.NewSlackServiceClient(httpClient, url),
		SsoServiceClient:                   adminv1connect.NewSsoServiceClient(httpClient, url),
		UserServiceClient:                  adminv1connect.NewUserServiceClient(httpClient, url),
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
