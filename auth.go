package zabbix

import (
	"log"
)

type AuthPrototype struct {
	AuthenticationType      int    `json:"authentication_type,omitempty"`
	HttpAuthEnabled         int    `json:"http_auth_enabled,omitempty"`
	HttpLoginForm           int    `json:"http_login_form,omitempty"`
	HttpStripDomains        string `json:"http_strip_domains,omitempty"`
	HttpCaseSensitive       int    `json:"http_case_sensitive,omitempty"`
	LdapConfigured          int    `json:"ldap_configured,omitempty"`
	LdapCaseSensitive       int    `json:"ldap_case_sensitive,omitempty"`
	LdapUserdirectoryid     int    `json:"ldap_userdirectoryid,omitempty"`
	SamlAuthEnabled         int    `json:"saml_auth_enabled,omitempty"`
	SamlIdpEntityid         string `json:"saml_idp_entityid,omitempty"`
	SamlSsoUrl              string `json:"saml_sso_url,omitempty"`
	SamlSloUrl              string `json:"saml_slo_url,omitempty"`
	SamlUsernameAttribute   string `json:"saml_username_attribute,omitempty"`
	SamlSpEntityid          string `json:"saml_sp_entityid,omitempty"`
	SamlNameidFormat        string `json:"saml_nameid_format,omitempty"`
	SamlSignMessages        int    `json:"saml_sign_messages,omitempty"`
	SamlSignAssertions      int    `json:"saml_sign_assertions,omitempty"`
	SamlSignAuthnRequests   int    `json:"saml_sign_authn_requests,omitempty"`
	SamlSignLogoutRequests  int    `json:"saml_sign_logout_requests,omitempty"`
	SamlSignLogoutResponses int    `json:"saml_sign_logout_responses,omitempty"`
	SamlEncryptNameid       int    `json:"saml_encrypt_nameid,omitempty"`
	SamlEncryptAssertions   int    `json:"saml_encrypt_assertions,omitempty"`
	SamlCaseSensitive       int    `json:"saml_case_sensitive,omitempty"`
	PasswdMinLength         int    `json:"passwd_min_length,omitempty"`
	PasswdCheckRules        string `json:"passwd_check_rules,omitempty"`
}

func (api *API) AuthGet(params Params) (res Response, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	res, err = api.CallWithError("authentication.get", params)

	if err != nil {
		return
	}

	return
}

func (api *API) AuthSet(params AuthPrototype) (res Response, err error) {

	res, err = api.CallWithError("authentication.update", params)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	return
}
