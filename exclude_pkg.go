package main

import (
	"os"
	"strings"
)

func parseExcludeConfig(filename string) error {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	lines := strings.Split(string(bs), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			excludePkg = append(excludePkg, line)
		}
	}
	return nil
}

func IsExcludePkg(pkg string) bool {
	for _, expkg := range excludePkg {
		if strings.HasPrefix(expkg, "prefix:") {
			ex := strings.TrimPrefix(expkg, "prefix:")
			if strings.HasPrefix(pkg, ex) {
				return true
			}
		} else {
			if expkg == pkg {
				return true
			}
		}
	}
	return false
}
