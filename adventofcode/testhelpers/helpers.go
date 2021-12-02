// package helpers provides useful helper functions to read inputs for advent of code
package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"testing"
)

const (
	inputDirectory       = "../../../.inputs"
	PermissionsReadWrite = 0o755
)

func ReadInputs(t *testing.T, adventYear, adventDay int) []int {
	t.Helper()

	if err := os.Mkdir(inputDirectory, PermissionsReadWrite); !errors.Is(err, os.ErrExist) {
		t.Fatalf("unable to create input dir: %v\n", err)
	}
	filename := path.Join(inputDirectory, fmt.Sprintf("%d-%d.txt", adventYear, adventDay))
	// filename, err := filepath.Abs(fn)
	// if err != nil {
	// 	t.Fatalf("failed to resolve test input file: %s", err)
	// }
	f, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	var inputs []int

	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Fatalf("failed converting text line to int: %v\n", err)
		}
		inputs = append(inputs, v)
	}
	t.Logf("filename: %s inputcount: %d\n", filepath.Base(filename), len(inputs))
	return inputs
}
