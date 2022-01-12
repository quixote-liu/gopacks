package test

import (
	"fmt"
	"testing"
)

func TestFindWildcard(t *testing.T) {
	path := "/go/proxy/:service/*os_api"

	wildcard, index, valid := FindWildcard(path)
	fmt.Printf("wildcard: %v\n", wildcard)
	fmt.Printf("index: %v\n", index)
	fmt.Printf("valid: %v\n", valid)

	fmt.Printf("path[:index]: %v\n", path[:index])
	fmt.Printf("path[index:]: %v\n", path[index:])
}
