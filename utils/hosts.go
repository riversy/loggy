package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/totherme/unstructured"
)

// getAdoptedScope converst scope value like "foo.bar"
// into "/foo/bar" to be used by "github.com/totherme/unstructured"
func getAdoptedScope(scope string) string {
	return fmt.Sprintf(
		"/%s",
		strings.ReplaceAll(scope, ".", "/"),
	)
}

// GetHosts read config and retrieves the list
// of URI strings for SSH connections
func GetHosts(configPath string, scope string) []string {
	configYaml, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	poolData, err := unstructured.ParseYAML(string(configYaml))
	if err != nil {
		panic(err)
	}

	poolPayloadData, err := poolData.GetByPointer(getAdoptedScope(scope))
	if err != nil {
		panic(err)
	}

	if !poolPayloadData.IsList() {
		panic("scoped value has to be list")
	}

	hostsList, err := poolPayloadData.ListValue()
	if err != nil {
		panic(err)
	}

	hosts := make([]string, len(hostsList))
	for i, v := range hostsList {
		hosts[i] = v.UnsafeStringValue()
	}

	return hosts
}
