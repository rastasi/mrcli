package commands

import (
	"flag"
	"fmt"
)

func Help() {
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("%s: %s\n", f.Name, f.Usage)
	})
}
