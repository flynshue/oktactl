package cmd

import (
	"testing"

	"github.com/flynshue/oktactl/pkg/okta-api/v2/oktaapi"
)

type MockOktaClient struct{}

func (m *MockOktaClient) ListApps(name string) ([]oktaapi.App, error) {
	return []oktaapi.App{
		{
			ID:    "0oa1gjh63g214q0Hq0g4",
			Name:  "testorgone_customsaml20app_1",
			Label: "Test Custom Saml 2.0 App",
		},
		{
			ID:    "0oabkvBLDEKCNXBGYUAS",
			Name:  "template_swa",
			Label: "Test Sample Plugin App",
		},
	}, nil
}

func (m *MockOktaClient) GetAppById(appID string) (oktaapi.App, error) {
	return oktaapi.App{
		ID:    "0oa1gjh63g214q0Hq0g4",
		Name:  "testorgone_customsaml20app_1",
		Label: "Test Custom Saml 2.0 App",
	}, nil
}

func (m *MockOktaClient) ListAppsGroups(appID string) (oktaapi.App, []oktaapi.GroupAssignmentResp, error) {
	profile := oktaapi.Profile{
		SAMLRoles: []string{"samlRoles01", "samlRoles02"},
		Role:      "ReadRole",
	}
	return oktaapi.App{
			ID:    "0oa1gjh63g214q0Hq0g4",
			Name:  "testorgone_customsaml20app_1",
			Label: "Test Custom Saml 2.0 App",
		}, []oktaapi.GroupAssignmentResp{
			{GroupID: "00gbkkGFFWZDLCNTAGQR", Name: "Fake Group 01", Profile: profile},
			{GroupID: "00gg0xVALADWBPXOFZAS", Name: "Fake Group 02", Profile: profile},
			{GroupID: "00gg0xVALADWBPXOFZAK", Name: "Fake Group 03", Profile: profile},
		}, nil
}

func (m *MockOktaClient) ListOktaGroups(name string) ([]oktaapi.Group, error) {
	return []oktaapi.Group{
		{ID: "00g1emaKYZTWRYYRRTSK", Profile: oktaapi.Profile{Name: "Fake Group 01"}},
		{ID: "00gg0xVALADWBPXOFZAS", Profile: oktaapi.Profile{Name: "Fake Group 02"}},
		{ID: "00gg0xVALADWBPXOFZAK", Profile: oktaapi.Profile{Name: "Fake Group 03"}},
	}, nil
}

func (m *MockOktaClient) ListOktaGroupUsers(groupID string) ([]oktaapi.User, error) {
	return []oktaapi.User{
		{ID: "00g1emaKYZTWRYYRRTSK", Profile: oktaapi.Profile{FirstName: "Test", LastName: "User-0", Email: "user0@example.com"}},
		{ID: "00gg0xVALADWBPXOFZAS", Profile: oktaapi.Profile{FirstName: "Test", LastName: "User_1", Email: "user1@example.com"}},
		{ID: "00gg0xVALADWBPXOFZAK", Profile: oktaapi.Profile{FirstName: "Test", LastName: "User'2", Email: "user2@example.com"}},
	}, nil
}

func TestListApps(t *testing.T) {
	if err := listApps(&MockOktaClient{}, "test"); err != nil {
		t.Error(err)
	}
}

func TestGetAppByID(t *testing.T) {
	if err := getAppById(&MockOktaClient{}, "0oa1gjh63g214q0Hq0g4"); err != nil {
		t.Error(err)
	}
}

func TestListAppsGroups(t *testing.T) {
	if err := listAppsGroups(&MockOktaClient{}, "0oa1gjh63g214q0Hq0g4"); err != nil {
		t.Error(err)
	}
}

func TestListOktaGroups(t *testing.T) {
	if err := listOktaGroups(&MockOktaClient{}, "test"); err != nil {
		t.Error(err)
	}
}

func TestListOktaGroupUsers(t *testing.T) {
	if err := listOktaGroupUsers(&MockOktaClient{}, "0oa1gjh63g214q0Hq0g4"); err != nil {
		t.Error(err)
	}
}
