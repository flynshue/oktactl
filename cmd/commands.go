/*
Copyright © 2024 Felicia Lyn-Shue

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// GitCommit is updated with the Git tag by the Goreleaser build
	GitCommit = "unknown"
	// BuildDate is updated with the current ISO timestamp by the Goreleaser build
	BuildDate = "unknown"
	// Version is updated with the latest tag by the Goreleaser build
	Version = "unreleased"
)

var listAppsCmd = &cobra.Command{
	Use:   "apps [name]",
	Short: "list apps by name",
	Long:  "Searches the name or label property of applications using startsWith that matches what the string starts with to the query",
	Example: `  # List apps that contain test in name or label
  oktactl list apps test

  Okta App ID            Name
  0oa1gjh63g214q0Hq0g4   Test Custom Saml 2.0 App
  0oabkvBLDEKCNXBGYUAS   Test Sample Plugin App
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must supply app name")
		}
		return listApps(newClient(), args[0])
	},
}

var listAppGroupAssignment = &cobra.Command{
	Use:   "groups [app ID]",
	Short: "List groups assigned to application",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must supply app id")
		}
		return listAppsGroups(newClient(), args[0])
	},
}

var listGroupsCmd = &cobra.Command{
	Use:   "groups [group name]",
	Short: "Searches the name property of groups using startsWith that matches what the string starts with to the query",
	Example: ` # List groups that start with fake
  oktactl list groups fake
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must supply group name")
		}
		keywords := strings.Join(args, " ")
		return listOktaGroups(newClient(), keywords)
	},
}

var listGroupUsersCmd = &cobra.Command{
	Use:   "users [group ID]",
	Short: "List users in group",
	Example: ` # List users in group
  oktactl list users 00g1hqieohhlPBv581d8
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("must supply group ID")
		}
		return listOktaGroupUsers(newClient(), args[0])
	},
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [command]",
	Short: "list resources",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version for oktactl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:\t", Version)
		fmt.Println("Git commit:\t", GitCommit)
		fmt.Println("Date:\t\t", BuildDate)
	},
}

func init() {
	rootCmd.AddCommand(listCmd, versionCmd)
	listCmd.AddCommand(listAppsCmd, listGroupsCmd, listGroupUsersCmd)
	listAppsCmd.AddCommand(listAppGroupAssignment)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
