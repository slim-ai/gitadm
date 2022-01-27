package describe

import (
	"fmt"
	"strings"

	_ "github.com/google/go-github/v42/github"
	"github.com/slim-ai/gitadm/pkg/scm"
	"github.com/slim-ai/gitadm/pkg/util"
	"github.com/urfave/cli/v2"
)

func getUserOrgMembershipsForGitlab(c *cli.Context) error {
	details, err := scm.GetUserOrganizationMembershipsForGitlab()
	if err != nil {
		return err
	}
	if !c.Bool("short") {
		util.PrettyPrint(details)
		return nil
	}
	paths := make([]string, len(details))
	for i, org := range details {
		paths[i] = fmt.Sprintf("gitlab.com/%s", org.Path)
	}
	util.Printf("%s\n", strings.Join(paths, ","))
	return nil
}

func getUserOrgMembershipsForGitHub(c *cli.Context) error {

	details, err := scm.GetUserOrganizationMembershipsForGitHub(c.String("username"))
	if err != nil {
		return err
	}
	if !c.Bool("short") {
		util.PrettyPrint(details)
		return nil
	}

	paths := make([]string, len(details))
	for i, org := range details {
		paths[i] = fmt.Sprintf(*org.HTMLURL)
	}
	util.Printf("%s\n", strings.Join(paths, ","))
	return nil
}
