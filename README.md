# Git Administration Tool

Graphical User Interfaces are awesome, but sometime you just need a dirty tool to make life easier.

## Installation

```
go install github.com/slim-ai/gitadm@latest
```

## Prerequisites
* [golang]() must be installed
* [gitlab personal token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html) - if you are using gitlab
* [github personal token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) - if you are using github

## Credential Setup

To communicate with your Gitlab or Github account, you must provide a token which will be used during API calls to the vendor. The following approaches all work:
### 1. Environment Variable
```
export GITLAB_API_TOKEN="glpat-MySuperSecretToken"
# now do something
gitadm describe user --username slimdevl
```

### 2. Configuration File
```
mkdir -p $HOME/.config/gitadm
echo "token: glpat-MySuperSecretToken" | tee $HOME/.config/gitadm/config

# now do something
gitadm describe user --username slimdevl
```

### 3. Command Option
```
gitadm --token "glpat-MySuperSecretToken" describe user --username slimdevl
```

## Available Commands
| Command                     | Description                                    |
| --------------------------- | ---------------------------------------------- |
| `get user`                  | Gets all the details on the specified user
| `get orgs`                  | Gets all orgs that the current user is a member of
| `get ssh-keys`              | Gets a list of all the ssh keys configured for the current user
| `add ssh-key`               | Adds an ssh key so SSH authentication can be used to access resouces
| `rm ssh-key`                | Deletes an SSH key


---

# Necessary Gitlab Scopes

| Scope                       | Usage Description                              |
| --------------------------- | ---------------------------------------------- |
| `api`                       | For accessing, and modification of user data. Including ssh-keys
| `read_user`                 | For getting user details
| `read_api`                  | For reading groups/organization command
| `sudo`                      | **not yet used.** This will be used for user administration later.
| `write_repository`          | **not yet used.** This will be used to create/update repositories later
| `write_registry`            | **not yet used.** This will be used to create/update repositories later

# Necessary Github Scopes

| Scope                       | Usage Description                              |
| --------------------------- | ---------------------------------------------- |
| TBD                         | no support quite yet