// Code generated by goa v3.11.3, DO NOT EDIT.
//
// server gRPC client CLI support package
//
// Command:
// $ goa gen stocktrader/design

package cli

import (
	"flag"
	"fmt"
	"os"
	sopc "stocktrader/gen/grpc/sop/client"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `sop plan
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` sop plan --message '{
      "symbol": "Nesciunt delectus."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, any, error) {
	var (
		sopFlags = flag.NewFlagSet("sop", flag.ContinueOnError)

		sopPlanFlags       = flag.NewFlagSet("plan", flag.ExitOnError)
		sopPlanMessageFlag = sopPlanFlags.String("message", "", "")
	)
	sopFlags.Usage = sopUsage
	sopPlanFlags.Usage = sopPlanUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "sop":
			svcf = sopFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "sop":
			switch epn {
			case "plan":
				epf = sopPlanFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "sop":
			c := sopc.NewClient(cc, opts...)
			switch epn {
			case "plan":
				endpoint = c.Plan()
				data, err = sopc.BuildPlanPayload(*sopPlanMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// sopUsage displays the usage of the sop command and its subcommands.
func sopUsage() {
	fmt.Fprintf(os.Stderr, `The sop service provides advisors with a comprehensive view of a particular stock schedule.
Usage:
    %[1]s [globalflags] sop COMMAND [flags]

COMMAND:
    plan: Plan implements plan.

Additional help:
    %[1]s sop COMMAND --help
`, os.Args[0])
}
func sopPlanUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sop plan -message JSON

Plan implements plan.
    -message JSON: 

Example:
    %[1]s sop plan --message '{
      "symbol": "Nesciunt delectus."
   }'
`, os.Args[0])
}