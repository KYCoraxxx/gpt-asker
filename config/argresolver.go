package config

import (
	myerror "gpt-asker/error"
)

type ArgResolver struct {
	opt map[string]string
}

var resolver *ArgResolver

func getResolver() *ArgResolver {
	if resolver == nil {
		resolver = &ArgResolver{opt: make(map[string]string)}
	}
	return resolver
}

func ResolveArgument(args []string) error {
	r := getResolver()

	for i := 1; i < len(args); i++ {
		if args[i][0] == '-' {
			if i+1 < len(args) {
				r.opt[args[i]] = args[i+1]
			} else {
				return myerror.NewArgError("Option " + args[i] + " requires input")
			}
		}
	}

	return nil
}
