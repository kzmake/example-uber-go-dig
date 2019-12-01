package cli

import (
	"github.com/spf13/cobra"
)

func update(ms ...map[string]*cobra.Command) map[string]*cobra.Command {
	res := map[string]*cobra.Command{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}

	return res
}

func list(ms ...map[string]*cobra.Command) map[string][]*cobra.Command {
	res := map[string][]*cobra.Command{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = append(res[k], v)
		}
	}

	return res
}

func merge(ms ...map[string][]*cobra.Command) map[string][]*cobra.Command {
	res := map[string][]*cobra.Command{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = append(res[k], v...)
		}
	}

	return res
}
