package describe

import (
	"fmt"
	"strings"

	"github.com/slim-ai/gitadm/pkg/scm"
	"github.com/slim-ai/gitadm/pkg/util"
	"github.com/urfave/cli/v2"
)

func getUserOrgMemberships(c *cli.Context) error {
	details, err := scm.GetUserOrganizationMemberships()
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
