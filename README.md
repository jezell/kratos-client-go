# Go API client for client

Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs
are exposed on different ports. Public APIs can face the public internet without any protection
while administrative APIs should never be exposed without prior authorization. To protect
the administative API port you should use something like Nginx, Ory Oathkeeper, or any other
technology capable of authorizing incoming requests.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.7.1-alpha.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import client "github.com/ory/kratos-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MetadataApi* | [**GetVersion**](docs/MetadataApi.md#getversion) | **Get** /version | Return Running Software Version.
*MetadataApi* | [**IsAlive**](docs/MetadataApi.md#isalive) | **Get** /health/alive | Check HTTP Server Status
*MetadataApi* | [**IsReady**](docs/MetadataApi.md#isready) | **Get** /health/ready | Check HTTP Server and Database Status
*MetadataApi* | [**Prometheus**](docs/MetadataApi.md#prometheus) | **Get** /metrics/prometheus | Get snapshot metrics from the service. If you&#39;re using k8s, you can then add annotations to your deployment like so:
*V0alpha1Api* | [**AdminCreateIdentity**](docs/V0alpha1Api.md#admincreateidentity) | **Post** /identities | Create an Identity
*V0alpha1Api* | [**AdminCreateSelfServiceRecoveryLink**](docs/V0alpha1Api.md#admincreateselfservicerecoverylink) | **Post** /recovery/link | Create a Recovery Link
*V0alpha1Api* | [**AdminDeleteIdentity**](docs/V0alpha1Api.md#admindeleteidentity) | **Delete** /identities/{id} | Delete an Identity
*V0alpha1Api* | [**AdminGetIdentity**](docs/V0alpha1Api.md#admingetidentity) | **Get** /identities/{id} | Get an Identity
*V0alpha1Api* | [**AdminListIdentities**](docs/V0alpha1Api.md#adminlistidentities) | **Get** /identities | List Identities
*V0alpha1Api* | [**AdminUpdateIdentity**](docs/V0alpha1Api.md#adminupdateidentity) | **Put** /identities/{id} | Update an Identity
*V0alpha1Api* | [**CreateSelfServiceLogoutFlowUrlForBrowsers**](docs/V0alpha1Api.md#createselfservicelogoutflowurlforbrowsers) | **Get** /self-service/logout/browser | Create a Logout URL for Browsers
*V0alpha1Api* | [**GetJsonSchema**](docs/V0alpha1Api.md#getjsonschema) | **Get** /schemas/{id} | 
*V0alpha1Api* | [**GetSelfServiceError**](docs/V0alpha1Api.md#getselfserviceerror) | **Get** /self-service/errors | Get Self-Service Errors
*V0alpha1Api* | [**GetSelfServiceLoginFlow**](docs/V0alpha1Api.md#getselfserviceloginflow) | **Get** /self-service/login/flows | Get Login Flow
*V0alpha1Api* | [**GetSelfServiceRecoveryFlow**](docs/V0alpha1Api.md#getselfservicerecoveryflow) | **Get** /self-service/recovery/flows | Get Recovery Flow
*V0alpha1Api* | [**GetSelfServiceRegistrationFlow**](docs/V0alpha1Api.md#getselfserviceregistrationflow) | **Get** /self-service/registration/flows | Get Registration Flow
*V0alpha1Api* | [**GetSelfServiceSettingsFlow**](docs/V0alpha1Api.md#getselfservicesettingsflow) | **Get** /self-service/settings/flows | Get Settings Flow
*V0alpha1Api* | [**GetSelfServiceVerificationFlow**](docs/V0alpha1Api.md#getselfserviceverificationflow) | **Get** /self-service/verification/flows | Get Verification Flow
*V0alpha1Api* | [**InitializeSelfServiceLoginFlowForBrowsers**](docs/V0alpha1Api.md#initializeselfserviceloginflowforbrowsers) | **Get** /self-service/login/browser | Initialize Login Flow for Browsers
*V0alpha1Api* | [**InitializeSelfServiceLoginFlowWithoutBrowser**](docs/V0alpha1Api.md#initializeselfserviceloginflowwithoutbrowser) | **Get** /self-service/login/api | Initialize Login Flow for APIs, Services, Apps, ...
*V0alpha1Api* | [**InitializeSelfServiceRecoveryFlowForBrowsers**](docs/V0alpha1Api.md#initializeselfservicerecoveryflowforbrowsers) | **Get** /self-service/recovery/browser | Initialize Recovery Flow for Browsers
*V0alpha1Api* | [**InitializeSelfServiceRecoveryFlowWithoutBrowser**](docs/V0alpha1Api.md#initializeselfservicerecoveryflowwithoutbrowser) | **Get** /self-service/recovery/api | Initialize Recovery Flow for APIs, Services, Apps, ...
*V0alpha1Api* | [**InitializeSelfServiceRegistrationFlowForBrowsers**](docs/V0alpha1Api.md#initializeselfserviceregistrationflowforbrowsers) | **Get** /self-service/registration/browser | Initialize Registration Flow for Browsers
*V0alpha1Api* | [**InitializeSelfServiceRegistrationFlowWithoutBrowser**](docs/V0alpha1Api.md#initializeselfserviceregistrationflowwithoutbrowser) | **Get** /self-service/registration/api | Initialize Registration Flow for APIs, Services, Apps, ...
*V0alpha1Api* | [**InitializeSelfServiceSettingsFlowForBrowsers**](docs/V0alpha1Api.md#initializeselfservicesettingsflowforbrowsers) | **Get** /self-service/settings/browser | Initialize Settings Flow for Browsers
*V0alpha1Api* | [**InitializeSelfServiceSettingsFlowWithoutBrowser**](docs/V0alpha1Api.md#initializeselfservicesettingsflowwithoutbrowser) | **Get** /self-service/settings/api | Initialize Settings Flow for APIs, Services, Apps, ...
*V0alpha1Api* | [**InitializeSelfServiceVerificationFlowForBrowsers**](docs/V0alpha1Api.md#initializeselfserviceverificationflowforbrowsers) | **Get** /self-service/verification/browser | Initialize Verification Flow for Browser Clients
*V0alpha1Api* | [**InitializeSelfServiceVerificationFlowWithoutBrowser**](docs/V0alpha1Api.md#initializeselfserviceverificationflowwithoutbrowser) | **Get** /self-service/verification/api | Initialize Verification Flow for APIs, Services, Apps, ...
*V0alpha1Api* | [**SubmitSelfServiceLoginFlow**](docs/V0alpha1Api.md#submitselfserviceloginflow) | **Post** /self-service/login | Submit a Login Flow
*V0alpha1Api* | [**SubmitSelfServiceLogoutFlow**](docs/V0alpha1Api.md#submitselfservicelogoutflow) | **Get** /self-service/logout | Complete Self-Service Logout
*V0alpha1Api* | [**SubmitSelfServiceLogoutFlowWithoutBrowser**](docs/V0alpha1Api.md#submitselfservicelogoutflowwithoutbrowser) | **Delete** /self-service/logout/api | Perform Logout for APIs, Services, Apps, ...
*V0alpha1Api* | [**SubmitSelfServiceRecoveryFlow**](docs/V0alpha1Api.md#submitselfservicerecoveryflow) | **Post** /self-service/recovery | Complete Recovery Flow
*V0alpha1Api* | [**SubmitSelfServiceRegistrationFlow**](docs/V0alpha1Api.md#submitselfserviceregistrationflow) | **Post** /self-service/registration | Submit a Registration Flow
*V0alpha1Api* | [**SubmitSelfServiceSettingsFlow**](docs/V0alpha1Api.md#submitselfservicesettingsflow) | **Post** /self-service/settings | Complete Settings Flow
*V0alpha1Api* | [**SubmitSelfServiceVerificationFlow**](docs/V0alpha1Api.md#submitselfserviceverificationflow) | **Post** /self-service/verification | Complete Verification Flow
*V0alpha1Api* | [**ToSession**](docs/V0alpha1Api.md#tosession) | **Get** /sessions/whoami | Check Who the Current HTTP Session Belongs To


## Documentation For Models

 - [AdminCreateIdentityBody](docs/AdminCreateIdentityBody.md)
 - [AdminCreateSelfServiceRecoveryLinkBody](docs/AdminCreateSelfServiceRecoveryLinkBody.md)
 - [AdminUpdateIdentityBody](docs/AdminUpdateIdentityBody.md)
 - [AuthenticateOKBody](docs/AuthenticateOKBody.md)
 - [ContainerChangeResponseItem](docs/ContainerChangeResponseItem.md)
 - [ContainerCreateCreatedBody](docs/ContainerCreateCreatedBody.md)
 - [ContainerTopOKBody](docs/ContainerTopOKBody.md)
 - [ContainerUpdateOKBody](docs/ContainerUpdateOKBody.md)
 - [ContainerWaitOKBody](docs/ContainerWaitOKBody.md)
 - [ContainerWaitOKBodyError](docs/ContainerWaitOKBodyError.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [GenericError](docs/GenericError.md)
 - [GraphDriverData](docs/GraphDriverData.md)
 - [HealthNotReadyStatus](docs/HealthNotReadyStatus.md)
 - [HealthStatus](docs/HealthStatus.md)
 - [IdResponse](docs/IdResponse.md)
 - [Identity](docs/Identity.md)
 - [IdentityCredentials](docs/IdentityCredentials.md)
 - [IdentityState](docs/IdentityState.md)
 - [ImageDeleteResponseItem](docs/ImageDeleteResponseItem.md)
 - [ImageSummary](docs/ImageSummary.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse503](docs/InlineResponse503.md)
 - [JsonError](docs/JsonError.md)
 - [Meta](docs/Meta.md)
 - [Plugin](docs/Plugin.md)
 - [PluginConfig](docs/PluginConfig.md)
 - [PluginConfigArgs](docs/PluginConfigArgs.md)
 - [PluginConfigInterface](docs/PluginConfigInterface.md)
 - [PluginConfigLinux](docs/PluginConfigLinux.md)
 - [PluginConfigNetwork](docs/PluginConfigNetwork.md)
 - [PluginConfigRootfs](docs/PluginConfigRootfs.md)
 - [PluginConfigUser](docs/PluginConfigUser.md)
 - [PluginDevice](docs/PluginDevice.md)
 - [PluginEnv](docs/PluginEnv.md)
 - [PluginInterfaceType](docs/PluginInterfaceType.md)
 - [PluginMount](docs/PluginMount.md)
 - [PluginSettings](docs/PluginSettings.md)
 - [Port](docs/Port.md)
 - [RecoveryAddress](docs/RecoveryAddress.md)
 - [SelfServiceError](docs/SelfServiceError.md)
 - [SelfServiceLoginFlow](docs/SelfServiceLoginFlow.md)
 - [SelfServiceLogoutUrl](docs/SelfServiceLogoutUrl.md)
 - [SelfServiceRecoveryFlow](docs/SelfServiceRecoveryFlow.md)
 - [SelfServiceRecoveryFlowState](docs/SelfServiceRecoveryFlowState.md)
 - [SelfServiceRecoveryLink](docs/SelfServiceRecoveryLink.md)
 - [SelfServiceRegistrationFlow](docs/SelfServiceRegistrationFlow.md)
 - [SelfServiceSettingsFlow](docs/SelfServiceSettingsFlow.md)
 - [SelfServiceSettingsFlowState](docs/SelfServiceSettingsFlowState.md)
 - [SelfServiceVerificationFlow](docs/SelfServiceVerificationFlow.md)
 - [SelfServiceVerificationFlowState](docs/SelfServiceVerificationFlowState.md)
 - [ServiceUpdateResponse](docs/ServiceUpdateResponse.md)
 - [Session](docs/Session.md)
 - [SettingsProfileFormConfig](docs/SettingsProfileFormConfig.md)
 - [SubmitSelfServiceLoginFlowBody](docs/SubmitSelfServiceLoginFlowBody.md)
 - [SubmitSelfServiceLoginFlowWithOidcMethodBody](docs/SubmitSelfServiceLoginFlowWithOidcMethodBody.md)
 - [SubmitSelfServiceLoginFlowWithPasswordMethodBody](docs/SubmitSelfServiceLoginFlowWithPasswordMethodBody.md)
 - [SubmitSelfServiceLogoutFlowWithoutBrowserBody](docs/SubmitSelfServiceLogoutFlowWithoutBrowserBody.md)
 - [SubmitSelfServiceRecoveryFlowBody](docs/SubmitSelfServiceRecoveryFlowBody.md)
 - [SubmitSelfServiceRecoveryFlowWithLinkMethodBody](docs/SubmitSelfServiceRecoveryFlowWithLinkMethodBody.md)
 - [SubmitSelfServiceRegistrationFlowBody](docs/SubmitSelfServiceRegistrationFlowBody.md)
 - [SubmitSelfServiceRegistrationFlowWithOidcMethodBody](docs/SubmitSelfServiceRegistrationFlowWithOidcMethodBody.md)
 - [SubmitSelfServiceRegistrationFlowWithPasswordMethodBody](docs/SubmitSelfServiceRegistrationFlowWithPasswordMethodBody.md)
 - [SubmitSelfServiceSettingsFlowBody](docs/SubmitSelfServiceSettingsFlowBody.md)
 - [SubmitSelfServiceSettingsFlowWithOidcMethodBody](docs/SubmitSelfServiceSettingsFlowWithOidcMethodBody.md)
 - [SubmitSelfServiceSettingsFlowWithPasswordMethodBody](docs/SubmitSelfServiceSettingsFlowWithPasswordMethodBody.md)
 - [SubmitSelfServiceSettingsFlowWithProfileMethodBody](docs/SubmitSelfServiceSettingsFlowWithProfileMethodBody.md)
 - [SubmitSelfServiceVerificationFlowBody](docs/SubmitSelfServiceVerificationFlowBody.md)
 - [SubmitSelfServiceVerificationFlowWithLinkMethodBody](docs/SubmitSelfServiceVerificationFlowWithLinkMethodBody.md)
 - [SuccessfulSelfServiceLoginWithoutBrowser](docs/SuccessfulSelfServiceLoginWithoutBrowser.md)
 - [SuccessfulSelfServiceRegistrationWithoutBrowser](docs/SuccessfulSelfServiceRegistrationWithoutBrowser.md)
 - [SuccessfulSelfServiceSettingsWithoutBrowser](docs/SuccessfulSelfServiceSettingsWithoutBrowser.md)
 - [UiContainer](docs/UiContainer.md)
 - [UiNode](docs/UiNode.md)
 - [UiNodeAnchorAttributes](docs/UiNodeAnchorAttributes.md)
 - [UiNodeAttributes](docs/UiNodeAttributes.md)
 - [UiNodeImageAttributes](docs/UiNodeImageAttributes.md)
 - [UiNodeInputAttributes](docs/UiNodeInputAttributes.md)
 - [UiNodeTextAttributes](docs/UiNodeTextAttributes.md)
 - [UiText](docs/UiText.md)
 - [VerifiableIdentityAddress](docs/VerifiableIdentityAddress.md)
 - [Version](docs/Version.md)
 - [Volume](docs/Volume.md)
 - [VolumeUsageData](docs/VolumeUsageData.md)


## Documentation For Authorization



### oryAccessToken

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

hi@ory.sh

