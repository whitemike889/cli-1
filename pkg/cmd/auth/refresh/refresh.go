package refresh

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/cli/cli/internal/config"
	"github.com/cli/cli/pkg/cmdutil"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/spf13/cobra"
)

type RefreshOptions struct {
	IO     *iostreams.IOStreams
	Config func() (config.Config, error)

	Hostname string
	Scopes   []string
}

func NewCmdRefresh(f *cmdutil.Factory, runF func(*RefreshOptions) error) *cobra.Command {
	opts := &RefreshOptions{
		IO:     f.IOStreams,
		Config: f.Config,
	}

	cmd := &cobra.Command{
		Use:   "refresh",
		Args:  cobra.ExactArgs(0),
		Short: "Request new scopes for a token",
		Long: heredoc.Doc(`Expand the permission scopes for a given host's token.

			This command allows you to add additional scopes to an existing authentication token via a web
			browser. This enables gh to access more of the GitHub API, which may be required as gh adds
			features or as you use the gh api command. 

			Unfortunately at this time there is no way to add scopes without a web browser's involvement
			due to how GitHub authentication works.

			The --hostname flag allows you to operate on a GitHub host other than github.com.

			The --scopes flag accepts a comma separated list of scopes you want to add to a token. If
			absent, this command ensures that a host's token has the default set of scopes required by gh.

			Note that if GITHUB_TOKEN is in the current environment, this command will not work.
		`),
		Example: heredoc.Doc(`
			$ gh auth refresh --scopes write:org,read:public_key
			# => open a browser to add write:org and read:public_key scopes for use with gh api

			$ gh auth refresh
			# => ensure that the required minimum scopes are enabled for a token and open a browser to add if not
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if runF != nil {
				return runF(opts)
			}

			return refreshRun(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Hostname, "hostname", "h", "", "The GitHub host to use for authentication")
	cmd.Flags().StringSliceVarP(&opts.Scopes, "scopes", "s", []string{}, "Additional scopes to add to a token")

	return cmd
}

func refreshRun(opts *RefreshOptions) error {
	// TODO check for GITHUB_TOKEN and error if found, mentioning token management URL

	// TODO ensure a token exists for host and validate it, recommending gh auth login if it fails

	// TODO need a function for checking if scopes are already in place

	// TODO need a form of authFlow that can take a list of scopes
	return nil
}
