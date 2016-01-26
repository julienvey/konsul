package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jawher/mow.cli"
	"github.com/racker/perigee"
)

func listServices(cmd *cli.Cmd) {
	cmd.Action = func() {
		var services ServiceList

		err := perigee.Get(fmt.Sprintf("%s/v1/agent/services", *consulURL), perigee.Options{
			Results: &services,
			OkCodes: []int{200},
		})
		if err != nil {
			fmt.Printf("Error %v", err)
		} else {
			w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
			fmt.Fprint(w, "ID\tSERVICE\tADDRESS\tPORT\n")
			for _, service := range services {
				fmt.Fprintf(w, "%s\t%s\t%s\t%d\n",
					service.ID,
					service.Service,
					service.Address,
					service.Port)
			}
			w.Flush()
		}
	}
}
