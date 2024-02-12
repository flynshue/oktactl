package oktaapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

type MockOktaAppService struct{}

var mockAS OktaAppService = &MockOktaAppService{}

func (m *MockOktaAppService) ListApplications(ctx context.Context, qp *query.Params) ([]okta.App, *okta.Response, error) {
	body := `[
		{
		  "id": "0oa1gjh63g214q0Hq0g4",
		  "name": "testorgone_customsaml20app_1",
		  "label": "Custom Saml 2.0 App",
		  "status": "ACTIVE",
		  "lastUpdated": "2016-08-09T20:12:19.000Z",
		  "created": "2016-08-09T20:12:19.000Z",
		  "accessibility": {
			"selfService": false,
			"errorRedirectUrl": null,
			"loginRedirectUrl": null
		  },
		  "visibility": {
			"autoSubmitToolbar": false,
			"hide": {
			  "iOS": false,
			  "web": false
			},
			"appLinks": {
			  "testorgone_customsaml20app_1_link": true
			}
		  },
		  "features": [],"signOnMode": "SAML_2_0",
		  "credentials": {
			"userNameTemplate": {
			  "template": "${fn:substringBefore(source.login, \"@\")}",
			  "type": "BUILT_IN"
			},
			"signing": {}
		  },
		  "settings": {
			"app": {},
			"notifications": {
			  "vpn": {
				"network": {
				  "connection": "DISABLED"
				},
				"message": null,
				"helpUrl": null
			  }
			},
			"signOn": {
			  "defaultRelayState": "",
			  "ssoAcsUrl": "https://{yourOktaDomain}",
			  "idpIssuer": "http://www.okta.com/${org.externalKey}",
			  "audience": "https://example.com/tenant/123",
			  "recipient": "http://recipient.okta.com",
			  "destination": "http://destination.okta.com",
			  "subjectNameIdTemplate": "${user.userName}",
			  "subjectNameIdFormat": "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress",
			  "responseSigned": true,
			  "assertionSigned": true,
			  "signatureAlgorithm": "RSA_SHA256",
			  "digestAlgorithm": "SHA256",
			  "honorForceAuthn": true,
			  "authnContextClassRef": "urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport",
			  "slo": {
				"enabled": true,
				"spIssuer": "https://testorgone.okta.com",
				"logoutUrl": "https://testorgone.okta.com/logout"
			  },
			  "participateSlo": {
				"enabled": true,
				"logoutRequestUrl": "https://testorgone.okta.com/logout/participate",
				"sessionIndexRequired": true,
				"bindingType": "REDIRECT"
			  },
			  "spCertificate": {
				"x5c": [
			  "MIIFnDCCA4QCCQDBSLbiON2T1zANBgkqhkiG9w0BAQsFADCBjzELMAkGA1UEBhMCVVMxDjAMBgNV\r\nBAgMBU1haW5lMRAwDgYDVQQHDAdDYXJpYm91MRcwFQYDVQQKDA5Tbm93bWFrZXJzIEluYzEUMBIG\r\nA1UECwwLRW5naW5lZXJpbmcxDTALBgNVBAMMBFNub3cxIDAeBgkqhkiG9w0BCQEWEWVtYWlsQGV4\r\nYW1wbGUuY29tMB4XDTIwMTIwMzIyNDY0M1oXDTMwMTIwMTIyNDY0M1owgY8xCzAJBgNVBAYTAlVT\r\nMQ4wDAYDVQQIDAVNYWluZTEQMA4GA1UEBwwHQ2FyaWJvdTEXMBUGA1UECgwOU25vd21ha2VycyBJ\r\nbmMxFDASBgNVBAsMC0VuZ2luZWVyaW5nMQ0wCwYDVQQDDARTbm93MSAwHgYJKoZIhvcNAQkBFhFl\r\nbWFpbEBleGFtcGxlLmNvbTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBANMmWDjXPdoa\r\nPyzIENqeY9njLan2FqCbQPSestWUUcb6NhDsJVGSQ7XR+ozQA5TaJzbP7cAJUj8vCcbqMZsgOQAu\r\nO/pzYyQEKptLmrGvPn7xkJ1A1xLkp2NY18cpDTeUPueJUoidZ9EJwEuyUZIktzxNNU1pA1lGijiu\r\n2XNxs9d9JR/hm3tCu9Im8qLVB4JtX80YUa6QtlRjWR/H8a373AYCOASdoB3c57fIPD8ATDNy2w/c\r\nfCVGiyKDMFB+GA/WTsZpOP3iohRp8ltAncSuzypcztb2iE+jijtTsiC9kUA2abAJqqpoCJubNShi\r\nVff4822czpziS44MV2guC9wANi8u3Uyl5MKsU95j01jzadKRP5S+2f0K+n8n4UoV9fnqZFyuGAKd\r\nCJi9K6NlSAP+TgPe/JP9FOSuxQOHWJfmdLHdJD+evoKi9E55sr5lRFK0xU1Fj5Ld7zjC0pXPhtJf\r\nsgjEZzD433AsHnRzvRT1KSNCPkLYomznZo5n9rWYgCQ8HcytlQDTesmKE+s05E/VSWNtH84XdDrt\r\nieXwfwhHfaABSu+WjZYxi9CXdFCSvXhsgufUcK4FbYAHl/ga/cJxZc52yFC7Pcq0u9O2BSCjYPdQ\r\nDAHs9dhT1RhwVLM8RmoAzgxyyzau0gxnAlgSBD9FMW6dXqIHIp8yAAg9cRXhYRTNAgMBAAEwDQYJ\r\nKoZIhvcNAQELBQADggIBADofEC1SvG8qa7pmKCjB/E9Sxhk3mvUO9Gq43xzwVb721Ng3VYf4vGU3\r\nwLUwJeLt0wggnj26NJweN5T3q9T8UMxZhHSWvttEU3+S1nArRB0beti716HSlOCDx4wTmBu/D1MG\r\nt/kZYFJw+zuzvAcbYct2pK69AQhD8xAIbQvqADJI7cCK3yRry+aWtppc58P81KYabUlCfFXfhJ9E\r\nP72ffN4jVHpX3lxxYh7FKAdiKbY2FYzjsc7RdgKI1R3iAAZUCGBTvezNzaetGzTUjjl/g1tcVYij\r\nltH9ZOQBPlUMI88lxUxqgRTerpPmAJH00CACx4JFiZrweLM1trZyy06wNDQgLrqHr3EOagBF/O2h\r\nhfTehNdVr6iq3YhKWBo4/+RL0RCzHMh4u86VbDDnDn4Y6HzLuyIAtBFoikoKM6UHTOa0Pqv2bBr5\r\nwbkRkVUxl9yJJw/HmTCdfnsM9dTOJUKzEglnGF2184Gg+qJDZB6fSf0EAO1F6sTqiSswl+uHQZiy\r\nDaZzyU7Gg5seKOZ20zTRaX3Ihj9Zij/ORnrARE7eM/usKMECp+7syUwAUKxDCZkGiUdskmOhhBGL\r\nJtbyK3F2UvoJoLsm3pIcvMak9KwMjSTGJB47ABUP1+w+zGcNk0D5Co3IJ6QekiLfWJyQ+kKsWLKt\r\nzOYQQatrnBagM7MI2/T4\r\n"
				]
			  },
		  "requestCompressed": false,
			  "allowMultipleAcsEndpoints": false,
			  "acsEndpoints": [],
			  "attributeStatements": [],
			  "inlineHooks": [
				{
				  "id": "${inlineHookId}",
				  "_links": {
					"self": {
					  "href": "https://{yourOktaDomain}/api/v1/inlineHooks/${inlineHookId}",
					  "hints": {
						"allow": [
						  "GET",
						  "PUT",
						  "DELETE"
						]
					  }
					}
				  }
				}
			  ]
			}
		  },
		  "_links": {
			"logo": [
			  {
				"name": "medium",
				"href": "http://testorgone.okta.com/assets/img/logos/default.6770228fb0dab49a1695ef440a5279bb.png",
				"type": "image/png"
			  }
			],
			"appLinks": [
			  {
				"name": "testorgone_customsaml20app_1_link",
				"href": "http://testorgone.okta.com/home/testorgone_customsaml20app_1/0oa1gjh63g214q0Hq0g4/aln1gofChJaerOVfY0g4",
				"type": "text/html"
			  }
			],
			"help": {
			  "href": "http://testorgone-admin.okta.com/app/testorgone_customsaml20app_1/0oa1gjh63g214q0Hq0g4/setup/help/SAML_2_0/instructions",
			  "type": "text/html"
			},
			"users": {
			  "href": "http://testorgone.okta.com/api/v1/apps/0oa1gjh63g214q0Hq0g4/users"
			},
			"deactivate": {
			  "href": "http://testorgone.okta.com/api/v1/apps/0oa1gjh63g214q0Hq0g4/lifecycle/deactivate"
			},
			"groups": {
			  "href": "http://testorgone.okta.com/api/v1/apps/0oa1gjh63g214q0Hq0g4/groups"
			},
			"metadata": {
			  "href": "http://testorgone.okta.com:/api/v1/apps/0oa1gjh63g214q0Hq0g4/sso/saml/metadata",
			  "type": "application/xml"
			}
		  }
		},
		{
		  "id": "0oabkvBLDEKCNXBGYUAS",
		  "name": "template_swa",
		  "label": "Sample Plugin App",
		  "status": "ACTIVE",
		  "lastUpdated": "2013-09-11T17:58:54.000Z",
		  "created": "2013-09-11T17:46:08.000Z",
		  "accessibility": {
			"selfService": false,
			"errorRedirectUrl": null
		  },
		  "visibility": {
			"autoSubmitToolbar": false,
			"hide": {
			  "iOS": false,
			  "web": false
			},
			"appLinks": {
			  "login": true
			}
		  },
		  "features": [],
		  "signOnMode": "BROWSER_PLUGIN",
		  "credentials": {
			"scheme": "EDIT_USERNAME_AND_PASSWORD",
			"userNameTemplate": {
			  "template": "${source.login}",
			  "type": "BUILT_IN"
			}
		  },
		  "settings": {
			"app": {
			  "buttonField": "btn-login",
			  "passwordField": "txtbox-password",
			  "usernameField": "txtbox-username",
			  "url": "https://example.com/login.html"
			}
		  },
		  "_links": {
			"logo": [
			  {
				"href": "https:/example.okta.com/img/logos/logo_1.png",
				"name": "medium",
				"type": "image/png"
			  }
			],
			"users": {
			  "href": "https://{yourOktaDomain}/api/v1/apps/0oabkvBLDEKCNXBGYUAS/users"
			},
			"groups": {
			  "href": "https://{yourOktaDomain}/api/v1/apps/0oabkvBLDEKCNXBGYUAS/groups"
			},
			"self": {
			  "href": "https://{yourOktaDomain}/api/v1/apps/0oabkvBLDEKCNXBGYUAS"
			},
			"deactivate": {
			  "href": "https://{yourOktaDomain}/api/v1/apps/0oabkvBLDEKCNXBGYUAS/lifecycle/deactivate"
			}
		  }
		}
	  ]
	  `
	buf := bytes.NewBufferString(body)
	resp := &http.Response{Body: io.NopCloser(buf), Status: "200 Ok", StatusCode: 200}
	return []okta.App{}, &okta.Response{Response: resp}, nil
}

func (m *MockOktaAppService) GetApplication(ctx context.Context, appId string, appInstance okta.App, qp *query.Params) (okta.App, *okta.Response, error) {
	body := `{
	"id": "0oa1gjh63g214q0Hq0g4",
	"name": "testorgone_customsaml20app_1",
	"label": "Custom Saml 2.0 App",
	"status": "ACTIVE",
	"lastUpdated": "2016-08-09T20:12:19.000Z",
	"created": "2016-08-09T20:12:19.000Z",
	"accessibility": {
		"selfService": false,
		"errorRedirectUrl": null,
		"loginRedirectUrl": null
	}
}`
	buf := bytes.NewBufferString(body)
	resp := &http.Response{Body: io.NopCloser(buf), Status: "200 Ok", StatusCode: 200}
	return okta.NewApplication(), &okta.Response{Response: resp}, nil
}

func (m *MockOktaAppService) ListApplicationGroupAssignments(ctx context.Context, appID string, qp *query.Params) ([]*okta.ApplicationGroupAssignment, *okta.Response, error) {
	body := `[
		{
		  "id": "00gbkkGFFWZDLCNTAGQR",
		  "lastUpdated": "2013-10-02T07:38:20.000Z",
		  "priority": 0,
		  "profile": {
			"samlRoles": ["samlRoles01","samlRoles02"],
			"role": "ReadRole"
		  }
		},
		{
		  "id": "00gg0xVALADWBPXOFZAS",
		  "lastUpdated": "2013-10-02T14:40:29.000Z",
		  "priority": 1,
		  "profile": {
			"samlRoles": ["samlRoles01","samlRoles02"],
			"role": "ReadRole"
		  }
		},
		{
			"id": "00gg0xVALADWBPXOFZAK",
			"lastUpdated": "2013-10-02T14:40:29.000Z",
			"priority": 1
		}
	  ]
	  `
	buf := bytes.NewBufferString(body)
	resp := &http.Response{Body: io.NopCloser(buf), Status: "200 Ok", StatusCode: 200}
	return nil, &okta.Response{Response: resp}, nil
}

type MockOktaGroupService struct{}

var mockGS OktaGroupService = &MockOktaGroupService{}

func (m *MockOktaGroupService) GetGroup(ctx context.Context, groupId string) (*okta.Group, *okta.Response, error) {
	body := `{
		"id": "00g1emaKYZTWRYYRRTSK",
		"created": "2015-02-06T10:11:28.000Z",
		"lastUpdated": "2015-10-05T19:16:43.000Z",
		"lastMembershipUpdated": "2015-11-28T19:15:32.000Z",
		"objectClass": [
		  "okta:user_group"
		],
		"type": "OKTA_GROUP",
		"profile": {
		  "name": "West Coast Users",
		  "description": "All Users West of The Rockies"
		},
		"_links": {
		  "logo": [
			{
			  "name": "medium",
			  "href": "https://{yourOktaDomain}/img/logos/groups/okta-medium.png",
			  "type": "image/png"
			},
			{
			  "name": "large",
			  "href": "https://{yourOktaDomain}/img/logos/groups/okta-large.png",
			  "type": "image/png"
			}
		  ],
		  "users": {
			"href": "https://{yourOktaDomain}/api/v1/groups/00g1emaKYZTWRYYRRTSK/users"
		  },
		  "apps": {
			"href": "https://{yourOktaDomain}/api/v1/groups/00g1emaKYZTWRYYRRTSK/apps"
		  }
		}
	  }
	  `
	buf := bytes.NewBufferString(body)
	resp := &http.Response{Body: io.NopCloser(buf), Status: "200 Ok", StatusCode: 200}
	return nil, &okta.Response{Response: resp}, nil
}

func (m *MockOktaGroupService) ListGroups(ctx context.Context, qp *query.Params) ([]*okta.Group, *okta.Response, error) {
	body := `[
		{
		  "id": "00g1emaKYZTWRYYRRTSK",
		  "created": "2015-02-06T10:11:28.000Z",
		  "lastUpdated": "2015-10-05T19:16:43.000Z",
		  "lastMembershipUpdated": "2015-11-28T19:15:32.000Z",
		  "objectClass": [
			"okta:user_group"
		  ],
		  "type": "OKTA_GROUP",
		  "profile": {
			"name": "West Coast Users",
			"description": "All Users West of The Rockies"
		  },
		  "_links": {
			"logo": [
			  {
				"name": "medium",
				"href": "https://{yourOktaDomain}/img/logos/groups/okta-medium.png",
				"type": "image/png"
			  },
			  {
				"name": "large",
				"href": "https://{yourOktaDomain}/img/logos/groups/okta-large.png",
				"type": "image/png"
			  }
			],
			"users": {
			  "href": "https://{yourOktaDomain}/api/v1/groups/00g1emaKYZTWRYYRRTSK/users"
			},
			"apps": {
			  "href": "https://{yourOktaDomain}/api/v1/groups/00g1emaKYZTWRYYRRTSK/apps"
			}
		  }
		},
		{
		  "id": "00gak46y5hydV6NdM0g4",
		  "created": "2015-07-22T08:45:03.000Z",
		  "lastUpdated": "2015-07-22T08:45:03.000Z",
		  "lastMembershipUpdated": "2015-10-22T08:45:03.000Z",
		  "objectClass": [
			"okta:user_group"
		  ],
		  "type": "OKTA_GROUP",
		  "profile": {
			"name": "Squabble of Users",
			"description": "Keep Calm and Single Sign-On"
		  },
		  "_links": {
			"logo": [
			  {
				"name": "medium",
				"href": "https://{yourOktaDomain}/img/logos/groups/okta-medium.png",
				"type": "image/png"
			  },
			  {
				"name": "large",
				"href": "https://{yourOktaDomain}/img/logos/groups/okta-large.png",
				"type": "image/png"
			  }
			],
			"users": {
			  "href": "https://{yourOktaDomain}/api/v1/groups/00gak46y5hydV6NdM0g4/users"
			},
			"apps": {
			  "href": "https://{yourOktaDomain}/api/v1/groups/00gak46y5hydV6NdM0g4/apps"
			}
		  }
		}
	  ]`

	buf := bytes.NewBufferString(body)
	return nil, &okta.Response{Response: &http.Response{Body: io.NopCloser(buf), Status: "200 OK", StatusCode: 200}}, nil
}

func TestOktaClient_ListApps(t *testing.T) {
	client := &OktaClient{OktaAppService: mockAS, OktaGroupService: mockGS, Ctx: context.Background()}
	apps, err := client.ListApps("datadog")
	if err != nil {
		t.Error(err)
	}
	for _, app := range apps {
		fmt.Printf("%s  %s\n", app.ID, app.Label)
	}
}

func TestOktaClient_GetAppById(t *testing.T) {
	client := &OktaClient{OktaAppService: mockAS, OktaGroupService: mockGS, Ctx: context.Background()}
	app, err := client.GetAppById("0oa1gjh63g214q0Hq0g4")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s  %s\n", app.ID, app.Label)
}

func TestOktaClient_ListAppsGroups(t *testing.T) {
	client := &OktaClient{OktaAppService: mockAS, OktaGroupService: mockGS, Ctx: context.Background()}
	app, groups, err := client.ListAppsGroups("0oa1gjh63g214q0Hq0g4")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("App %s %s group assignment\n", app.ID, app.Label)
	fmt.Printf("groups %d\n", len(groups))
	for _, group := range groups {
		fmt.Printf("%s  %s\n", group.GroupID, group.Name)
		for _, roles := range group.SAMLRoles {
			fmt.Println(roles)
		}
		if group.Role != "" {
			fmt.Println(group.Role)
		}
	}
}

func TestOktaClient_ListGroups(t *testing.T) {
	client := &OktaClient{OktaAppService: mockAS, OktaGroupService: mockGS, Ctx: context.Background()}
	groups, err := client.ListOktaGroups("test")
	if err != nil {
		t.Error(err)
	}
	for _, g := range groups {
		fmt.Printf("%s  %s\n", g.ID, g.Profile.Name)
	}
}
