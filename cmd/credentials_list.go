package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mobingi/mobingi/pkg/cli"
	"github.com/mobingilabs/mobingi-sdk-go/mobingi/credentials"
	"github.com/mobingilabs/mobingi-sdk-go/pkg/cmdline"
	"github.com/mobingilabs/mobingi-sdk-go/pkg/pretty"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func CredentialsListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list your credentials",
		Long: `List your credentials.

Supported formats: raw, json

Examples:

  $ ` + cmdline.Args0() + ` creds list
  $ ` + cmdline.Args0() + ` creds list --fmt=json
  $ ` + cmdline.Args0() + ` creds list --fmt=raw`,
		Run: credsList,
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().String("vendor", "aws", "credentials vendor")
	return cmd
}

func credsList(cmd *cobra.Command, args []string) {
	creds, body, err := getCredsList(cmd)
	cli.ErrorExit(err, 1)

	pfmt := cli.GetCliStringFlag(cmd, "fmt")
	switch pfmt {
	case "raw":
		fmt.Println(string(body))
	case "json":
		indent := cli.GetCliIntFlag(cmd, "indent")
		mi, err := json.MarshalIndent(creds, "", pretty.Indent(indent))
		cli.ErrorExit(err, 1)

		fmt.Println(string(mi))
	default:
		w := tabwriter.NewWriter(os.Stdout, 0, 10, 5, ' ', 0)
		fmt.Fprintf(w, "VENDOR\tID\tACCOUNT\tLAST MODIFIED\n")
		for _, cred := range creds {
			timestr := cred.LastModified
			t, err := time.Parse(time.RFC3339, cred.LastModified)
			if err == nil {
				timestr = t.Format(time.RFC1123)
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
				cli.GetCliStringFlag(cmd, "vendor"),
				cred.Id,
				cred.Account,
				timestr)
		}

		w.Flush()
	}
}

func getCredsList(cmd *cobra.Command) ([]credentials.VendorCredentials, []byte, error) {
	vendor := cli.GetCliStringFlag(cmd, "vendor")
	sess, err := clisession()
	if err != nil {
		return nil, nil, errors.Wrap(err, "clisession failed")
	}

	svc := credentials.New(sess)
	resp, body, err := svc.List(&credentials.CredentialsListInput{
		Vendor: vendor,
	})

	if err != nil {
		return nil, body, errors.Wrap(err, "svc list failed")
	}

	exitOn401(resp)
	var creds []credentials.VendorCredentials
	err = json.Unmarshal(body, &creds)
	if err != nil {
		return nil, body, errors.Wrap(err, "unmarshal failed")
	}

	return creds, body, nil
}
