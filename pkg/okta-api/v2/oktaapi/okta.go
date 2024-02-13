package oktaapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
)

type OktaAppService interface {
	ListApplications(ctx context.Context, qp *query.Params) ([]okta.App, *okta.Response, error)
	ListApplicationGroupAssignments(ctx context.Context, appID string, qp *query.Params) ([]*okta.ApplicationGroupAssignment, *okta.Response, error)
	GetApplication(ctx context.Context, appId string, appInstance okta.App, qp *query.Params) (okta.App, *okta.Response, error)
}

type OktaGroupService interface {
	GetGroup(ctx context.Context, groupId string) (*okta.Group, *okta.Response, error)
	ListGroups(ctx context.Context, qp *query.Params) ([]*okta.Group, *okta.Response, error)
	ListGroupUsers(ctx context.Context, groupId string, qp *query.Params) ([]*okta.User, *okta.Response, error)
}

type App struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label,omitempty"`
}

type User struct {
	ID      string `json:"id"`
	Profile `json:"profile,omitempty"`
}

type Group struct {
	ID                    string `json:"id"`
	LastUpdated           string `json:"lastUpdated"`
	LastMembershipUpdated string `json:"lastMembershipUpdated"`
	Type                  string `json:"type"`
	Profile               `json:"profile,omitempty"`
}

type GroupAssignmentResp struct {
	GroupID string `json:"id"`
	Name    string `json:"name,omitempty"`
	Profile `json:"profile,omitempty"`
}

type Profile struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	SAMLRoles   []string `json:"samlRoles,omitempty"`
	Role        string   `json:"role,omitempty"`
	Email       string   `json:"email,omitempty"`
	FirstName   string   `json:"firstName,omitempty"`
	LastName    string   `json:"lastName,omitempty"`
}

type OktaClient struct {
	OktaAppService
	OktaGroupService
	Ctx context.Context
}

func NewClient(url, token string) (*OktaClient, error) {
	ctx, client, err := okta.NewClient(context.Background(), okta.WithOrgUrl(url), okta.WithToken(token))
	if err != nil {
		return nil, err
	}
	return &OktaClient{OktaAppService: client.Application, OktaGroupService: client.Group, Ctx: ctx}, nil
}

func (oc *OktaClient) ListApps(name string) ([]App, error) {
	qp := query.NewQueryParams(query.WithQ(name), query.WithFilter("status eq \"ACTIVE\""))
	_, resp, err := oc.OktaAppService.ListApplications(oc.Ctx, qp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	apps := []App{}
	if err := json.Unmarshal(b, &apps); err != nil {
		return nil, err
	}
	return apps, nil
}

func (oc *OktaClient) ListAppsGroups(appID string) (App, []GroupAssignmentResp, error) {
	app, err := oc.GetAppById(appID)
	if err != nil {
		return app, nil, err
	}
	params := query.NewQueryParams(query.WithLimit(200))
	_, resp, err := oc.OktaAppService.ListApplicationGroupAssignments(oc.Ctx, appID, params)
	if err != nil {
		return app, nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return app, nil, err
	}
	groups := []GroupAssignmentResp{}
	if err := json.Unmarshal(b, &groups); err != nil {
		return app, nil, err
	}
	for i, group := range groups {
		g, _ := oc.GetGroupById(group.GroupID)
		group.Name = g.Name
		groups[i] = group
	}
	return app, groups, nil
}

func (oc *OktaClient) ListOktaGroups(name string) ([]Group, error) {
	params := query.NewQueryParams(query.WithLimit(100), query.WithSearch(fmt.Sprintf("profile.name sw \"%s\"", name)))
	_, resp, err := oc.ListGroups(oc.Ctx, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	groups := []Group{}
	if err := json.Unmarshal(b, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

func (oc *OktaClient) ListOktaGroupUsers(groupID string) ([]User, error) {
	params := query.NewQueryParams(query.WithLimit(100))
	_, resp, err := oc.ListGroupUsers(oc.Ctx, groupID, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	users := []User{}
	if err := json.Unmarshal(b, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (oc *OktaClient) GetAppById(appID string) (App, error) {
	_, resp, err := oc.OktaAppService.GetApplication(oc.Ctx, appID, okta.NewApplication(), &query.Params{})
	app := App{}
	if err != nil {
		return app, nil
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return app, err
	}
	if err := json.Unmarshal(b, &app); err != nil {
		return app, err
	}
	return app, nil
}

func (oc *OktaClient) GetGroupById(groupID string) (Group, error) {
	group := Group{}
	_, resp, err := oc.OktaGroupService.GetGroup(oc.Ctx, groupID)
	if err != nil {
		return group, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return group, err
	}
	if err := json.Unmarshal(b, &group); err != nil {
		return group, err
	}
	return group, nil
}
