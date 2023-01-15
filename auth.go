package zabbix

type AuthenticationSettings struct {
	AuthenticationType      int    `json:"authentication_type,string"`
	HttpAuthEnabled         int    `json:"http_auth_enabled,string"`
	HttpLoginForm           int    `json:"http_login_form,string"`
	HttpStripDomains        string `json:"http_strip_domains"`
	HttpCaseSensitive       int    `json:"http_case_sensitive,string"`
	LdapConfigured          int    `json:"ldap_configured,string"`
	LdapCaseSensitive       int    `json:"ldap_case_sensitive,string"`
	LdapUserdirectoryid     int    `json:"ldap_userdirectoryid,string,omitempty"`
	SamlAuthEnabled         int    `json:"saml_auth_enabled,string"`
	SamlIdpEntityid         string `json:"saml_idp_entityid"`
	SamlSsoUrl              string `json:"saml_sso_url"`
	SamlSloUrl              string `json:"saml_slo_url"`
	SamlUsernameAttribute   string `json:"saml_username_attribute"`
	SamlSpEntityid          string `json:"saml_sp_entityid"`
	SamlNameidFormat        string `json:"saml_nameid_format"`
	SamlSignMessages        int    `json:"saml_sign_messages,string"`
	SamlSignAssertions      int    `json:"saml_sign_assertions,string"`
	SamlSignAuthnRequests   int    `json:"saml_sign_authn_requests,string"`
	SamlSignLogoutRequests  int    `json:"saml_sign_logout_requests,string"`
	SamlSignLogoutResponses int    `json:"saml_sign_logout_responses,string"`
	SamlEncryptNameid       int    `json:"saml_encrypt_nameid,string"`
	SamlEncryptAssertions   int    `json:"saml_encrypt_assertions,string"`
	SamlCaseSensitive       int    `json:"saml_case_sensitive,string"`
	PasswdMinLength         int    `json:"passwd_min_length,string,omitempty"`
	PasswdCheckRules        int    `json:"passwd_check_rules,string"`
}

func (api *API) AuthGet() (authenticationSettings *AuthenticationSettings, err error) {
	authenticationSettings = &AuthenticationSettings{}
	err = api.CallWithErrorParse("authentication.get", Params{"output": "extend"}, authenticationSettings)
	return
}

func (api *API) AuthSet(authenticationSettings *AuthenticationSettings) (err error) {
	_, err = api.CallWithError("authentication.update", authenticationSettings)
	return
}
