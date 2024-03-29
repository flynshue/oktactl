package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/flynshue/oktactl/pkg/okta-api/v2/oktaapi"
	"github.com/spf13/viper"
)

var (
	client *oktaapi.OktaClient
	err    error
)

type OktaService interface {
	ListApps(name string) ([]oktaapi.App, error)
	GetAppById(appID string) (oktaapi.App, error)
	ListAppsGroups(appID string) (oktaapi.App, []oktaapi.GroupAssignmentResp, error)
	ListOktaGroups(name string) ([]oktaapi.Group, error)
	ListOktaGroupUsers(groupID string) ([]oktaapi.User, error)
}

func listApps(os OktaService, name string) error {
	apps, err := os.ListApps(name)
	if err != nil {
		return err
	}
	if len(apps) == 0 {
		fmt.Printf("no apps found using keyword %s\n", name)
	}
	w := newTabWriter()
	fmt.Fprintln(w, "Okta App ID\t Name\t")
	for _, app := range apps {
		fmt.Fprintf(w, "%s\t %s\t\n", app.ID, app.Label)
	}
	w.Flush()
	return nil
}

func getAppById(os OktaService, appID string) error {
	app, err := os.GetAppById(appID)
	if err != nil {
		return err
	}
	w := newTabWriter()
	fmt.Fprintln(w, "Okta App ID\t Name\t")
	fmt.Fprintf(w, "%s\t %s\t\n", app.ID, app.Label)
	w.Flush()
	return nil
}

func listAppsGroups(os OktaService, appID string) error {
	app, groups, err := os.ListAppsGroups(appID)
	if err != nil {
		return err
	}
	fmt.Printf("Group assignment for %s %s\n", app.ID, app.Label)
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
	return nil
}

func listOktaGroups(os OktaService, keyword string) error {
	groups, err := os.ListOktaGroups(keyword)
	if err != nil {
		return err
	}
	w := newTabWriter()
	fmt.Fprintln(w, "Okta Group ID\t Name\t")
	for _, group := range groups {
		fmt.Fprintf(w, "%s\t %s\t\n", group.ID, group.Name)
	}
	w.Flush()
	return nil
}

func listOktaGroupUsers(os OktaService, groupID string) error {
	users, err := os.ListOktaGroupUsers(groupID)
	if err != nil {
		return err
	}
	w := newTabWriter()
	fmt.Fprintln(w, "Okta User ID\t First Name\t Last Name\t Email\t")
	for _, user := range users {
		fmt.Fprintf(w, "%s\t %s\t %s\t %s\n", user.ID, user.FirstName, user.LastName, user.Email)
	}
	w.Flush()
	return nil
}

func newClient() *oktaapi.OktaClient {
	if client != nil {
		return client
	}
	client, err = oktaapi.NewClient(viper.GetString("org"), viper.GetString("token"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func newTabWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)
}
