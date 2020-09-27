package octopusdeploy

const (
	constAccessKey                                   string = "access_key"
	constAccount                                     string = "account"
	constAccountType                                 string = "account_type"
	constAccountTypeAzure                            string = "Azure"
	constAccountTypeAWS                              string = "AWS"
	constActions                                     string = "actions"
	constActionType                                  string = "action_type"
	constActiveDirectoryEndpointBaseURI              string = "active_directory_endpoint_base_uri"
	constAcquisitionLocation                         string = "acquisition_location"
	constAddress                                     string = "address"
	constAllowDynamicInfrastructure                  string = "allow_dynamic_infrastructure"
	constAPIKey                                      string = "apikey"
	constApplyTerraformAction                        string = "apply_terraform_action"
	constArguments                                   string = "arguments"
	constAuthentication                              string = "authentication"
	constAutomaticDeploymentTargets                  string = "automatic_deployment_targets"
	constAzureEnvironment                            string = "azure_environment"
	constCertificate                                 string = "certificate"
	constCertificateData                             string = "certificate_data"
	constChannel                                     string = "channel"
	constChannels                                    string = "channels"
	constClientID                                    string = "client_id"
	constColor                                       string = "color"
	constCondition                                   string = "condition"
	constConditionExpression                         string = "condition_expression"
	constCustomAccountName                           string = "custom_account_name"
	constCustomAccountPassword                       string = "custom_account_password"
	constDefaultFailureMode                          string = "default_failure_mode"
	constDependencies                                string = "dependencies"
	constDeployKubernetesSecretAction                string = "deploy_kubernetes_secret_action"
	constDeployPackageAction                         string = "deploy_package_action"
	constDeployWindowsServiceAction                  string = "deploy_windows_service_action"
	constDescription                                 string = "description"
	constDisplayName                                 string = "display_name"
	constDownloadAttempts                            string = "download_attempts"
	constDownloadRetryBackoffSeconds                 string = "download_retry_backoff_seconds"
	constEmptyString                                 string = ""
	constEncryptedValue                              string = "encrypted_value"
	constEndpointCommunicationStyle                  string = "endpoint_communicationstyle"
	constEndpoint                                    string = "endpoint"
	constEndpointID                                  string = "endpoint_id"
	constEndpointThumbprint                          string = "endpoint_thumbprint"
	constEndpointURI                                 string = "endpoint_uri"
	constEnhancedMode                                string = "enhanced_mode"
	constEnvironment                                 string = "environment"
	constEnvironmentIDs                              string = "environment_ids"
	constEnvironments                                string = "environments"
	constExcludedEnvironments                        string = "excluded_environments"
	constExecutablePath                              string = "executable_path"
	constExtractDuringDeployment                     string = "extract_during_deployment"
	constEventCategories                             string = "event_categories"
	constEventGroups                                 string = "event_groups"
	constFeed                                        string = "feed"
	constFeedID                                      string = "feed_id"
	constFeedType                                    string = "feed_type"
	constFeedURI                                     string = "feed_uri"
	constInstructions                                string = "instructions"
	constIsDefault                                   string = "is_default"
	constIsOptionalPhase                             string = "is_optional_phase"
	constIsSensitive                                 string = "is_sensitive"
	constKey                                         string = "key"
	constKeyFingerprint                              string = "key_fingerprint"
	constLabel                                       string = "label"
	constLibraryVariableSet                          string = "library_variable_set"
	constLifecycle                                   string = "lifecycle"
	constLifecycleID                                 string = "lifecycle_id"
	constMachine                                     string = "machine"
	constMachines                                    string = "machines"
	constMachinePolicy                               string = "machine_policy"
	constManualInterventionAction                    string = "manual_intervention_action"
	constMinimumEnvironmentsBeforePromotion          string = "minimum_environments_before_promotion"
	constName                                        string = "name"
	constNamespace                                   string = "namespace"
	constNotes                                       string = "notes"
	constOctopusDeployAccount                        string = "octopusdeploy_account"
	constOctopusDeployAWSAccount                     string = "octopusdeploy_aws_account"
	constOctopusDeployAzureServicePrincipal          string = "octopusdeploy_azure_service_principal"
	constOctopusDeployCertificate                    string = "octopusdeploy_certificate"
	constOctopusDeployChannel                        string = "octopusdeploy_channel"
	constOctopusDeployDeploymentProcess              string = "octopusdeploy_deployment_process"
	constOctopusDeployEnvironment                    string = "octopusdeploy_environment"
	constOctopusDeployFeed                           string = "octopusdeploy_feed"
	constOctopusDeployLibraryVariableSet             string = "octopusdeploy_library_variable_set"
	constOctopusDeployLifecycle                      string = "octopusdeploy_lifecycle"
	constOctopusDeployMachine                        string = "octopusdeploy_machine"
	constOctopusDeployMachinePolicy                  string = "octopusdeploy_machinepolicy"
	constOctopusDeployNuGetFeed                      string = "octopusdeploy_nuget_feed"
	constOctopusDeployProject                        string = "octopusdeploy_project"
	constOctopusDeployProjectDeploymentTargetTrigger string = "octopusdeploy_project_deployment_target_trigger"
	constOctopusDeployProjectGroup                   string = "octopusdeploy_project_group"
	constOctopusDeploySSHKeyAccount                  string = "octopusdeploy_sshkey_account"
	constOctopusDeployTagSet                         string = "octopusdeploy_tag_set"
	constOctopusDeployUsernamePasswordAccount        string = "octopusdeploy_usernamepassword_account"
	constOctopusDeployVariable                       string = "octopusdeploy_variable"
	constOptionalDeploymentTargets                   string = "optional_deployment_targets"
	constPackage                                     string = "package"
	constPackageID                                   string = "package_id"
	constPackageRequirement                          string = "package_requirement"
	constPassword                                    string = "password"
	constPhase                                       string = "phase"
	constPrimaryPackage                              string = "primary_package"
	constProject                                     string = "project"
	constProjectGroupID                              string = "project_group_id"
	constProjectID                                   string = "project_id"
	constPrompt                                      string = "prompt"
	constProperty                                    string = "property"
	constQuantityToKeep                              string = "quantity_to_keep"
	constRequired                                    string = "required"
	constResourceManagementEndpointBaseURI           string = "resource_management_endpoint_base_uri"
	constResponsibleTeams                            string = "responsible_teams"
	constReleaseRetentionPolicy                      string = "release_retention_policy"
	constRoles                                       string = "roles"
	constRule                                        string = "rule"
	constRunKubectlScriptAction                      string = "run_kubectl_script_action"
	constRunOnServer                                 string = "run_on_server"
	constScope                                       string = "scope"
	constScriptParameters                            string = "script_parameters"
	constScriptFileName                              string = "script_file_name"
	constSecretKey                                   string = "secret_key"
	constSecretName                                  string = "secret_name"
	constSecretValues                                string = "secret_values"
	constSensitiveValue                              string = "sensitive_value"
	constServiceAccount                              string = "service_account"
	constServiceName                                 string = "service_name"
	constShouldRedeploy                              string = "should_redeploy"
	constSkipMachineBehavior                         string = "skip_machine_behavior"
	constSpaceID                                     string = "space_id"
	constStartMode                                   string = "start_mode"
	constStartTrigger                                string = "start_trigger"
	constStatus                                      string = "status"
	constSubscriptionNumber                          string = "subscription_number"
	constTag                                         string = "tag"
	constTagSet                                      string = "tag_set"
	constTargetRoles                                 string = "target_roles"
	constTemplates                                   string = "templates"
	constTenantID                                    string = "tenant_id"
	constTenantIDs                                   string = "tenant_ids"
	constTenants                                     string = "tenants"
	constTenantTags                                  string = "tenant_tags"
	constTenantedDeploymentParticipation             string = "tenanted_deployment_participation"
	constTentacleRetentionPolicy                     string = "tentacle_retention_policy"
	constThumbprint                                  string = "thumbprint"
	constTrue                                        string = "true"
	constType                                        string = "type"
	constUnit                                        string = "unit"
	constURI                                         string = "uri"
	constUseGuidedFailure                            string = "use_guided_failure"
	constUsername                                    string = "username"
	constValue                                       string = "value"
	constVariableSetID                               string = "variable_set_id"
	constVariableSubstitutionInFiles                 string = "variable_substitution_in_files"
	constVersionRange                                string = "version_range"
	constWindowsService                              string = "windows_service"
	constWindowSize                                  string = "window_size"
	constWorkerPoolID                                string = "worker_pool_id"

	CertificateIDs                                 string = "Certificate_ids"
	EndpointCommunicationStyle                     string = "endpoint_communicationstyle"
	EndpointProxyID                                string = "endpoint_proxyid"
	EndpointTentacleVersionDetailsUpgradeLocked    string = "endpoint_tentacleversiondetails_upgradelocked"
	EndpointTentacleVersionDetailsUpgradeRequired  string = "endpoint_tentacleversiondetails_upgraderequired"
	EndpointTentacleVersionDetailsUpgradeSuggested string = "endpoint_tentacleversiondetails_upgradesuggested"
	EndpointTentacleVersionDetailsVersion          string = "endpoint_tentacleversiondetails_version"
	EndpointThumbprint                             string = "endpoint_thumbprint"
	Extract                                        string = "Extract"
	HasLatestCalamari                              string = "haslatestcalamari"
	IsDefault2                                     string = "isdefault"
	IsInProcess                                    string = "isinprocess"
	MachinePolicy                                  string = "machinepolicy"
	PGPKey                                         string = "pgp_key"
	RunScriptAction                                string = "run_script_action"
	StatusSummary                                  string = "statussummary"
)