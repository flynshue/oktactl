# oktactl
oktactl is CLI based tool to help okta admins quickly find okta apps by name, list app group assignments, okta groups, and users in group.

## Configuration
Before you can run any of the commands, you'll need to create a config file `$HOME/.oktactl.yaml`

```bash
vim ~/.oktactl.yaml
```

```yaml
org: "https://yourOrg.okta.com"
token: "fakeToken"
```

You'll need an okta api token for your org that has at least read permissions for Applications, Users and Groups (Application Reader role and User and Group Reader)