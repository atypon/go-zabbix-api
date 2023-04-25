package zabbix

type AuthenticationSettings struct {
	AuthenticationType int    `json:"authentication_type,string,omitempty"`
	HttpAuthEnabled    int    `json:"http_auth_enabled,string,omitempty"`
	HttpLoginForm      int    `json:"http_login_form,string,omitempty"`
	HttpStripDomains   string `json:"http_strip_domains,omitempty"`
	HttpCaseSensitive  int    `json:"http_case_sensitive,string,omitempty"`
	LdapConfigured     int    `json:"ldap_configured,string,omitempty"`
	LdapCaseSensitive  int    `json:"ldap_case_sensitive,string,omitempty"`
	//LdapUserdirectoryid     string `json:"ldap_userdirectoryid,string,omitempty"`
	SamlAuthEnabled         int    `json:"saml_auth_enabled,string,omitempty"`
	SamlIdpEntityid         string `json:"saml_idp_entityid,omitempty"`
	SamlSsoUrl              string `json:"saml_sso_url,omitempty"`
	SamlSloUrl              string `json:"saml_slo_url,omitempty"`
	SamlUsernameAttribute   string `json:"saml_username_attribute,omitempty"`
	SamlSpEntityid          string `json:"saml_sp_entityid,omitempty"`
	SamlNameidFormat        string `json:"saml_nameid_format,omitempty"`
	SamlSignMessages        int    `json:"saml_sign_messages,string,omitempty"`
	SamlSignAssertions      int    `json:"saml_sign_assertions,string,omitempty"`
	SamlSignAuthnRequests   int    `json:"saml_sign_authn_requests,string,omitempty"`
	SamlSignLogoutRequests  int    `json:"saml_sign_logout_requests,string,omitempty"`
	SamlSignLogoutResponses int    `json:"saml_sign_logout_responses,string,omitempty"`
	SamlEncryptNameid       int    `json:"saml_encrypt_nameid,string,omitempty"`
	SamlEncryptAssertions   int    `json:"saml_encrypt_assertions,string,omitempty"`
	SamlCaseSensitive       int    `json:"saml_case_sensitive,string,omitempty"`
	//PasswdMinLength         int    `json:"passwd_min_length,string,omitempty"`
	//PasswdCheckRules int `json:"passwd_check_rules,string,omitempty"`
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
