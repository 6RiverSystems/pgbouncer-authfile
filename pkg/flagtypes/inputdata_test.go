package flagtypes_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/6RiverSystems/pgbouncer-authfile/pkg/flagtypes"
)

func TestInputData(t *testing.T) {
	const (
		wantContent = "foo"
		wantValue   = "bar"
	)

	file, err := ioutil.TempFile("", "test-input-data")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	ioutil.WriteFile(file.Name(), []byte(wantContent), 0666)

	var v flagtypes.InputData

	if err = v.Set(file.Name()); err != nil {
		t.Fatalf("expected no error on setting %q but was: %v", file.Name(), err)
	}

	if err = v.Set(wantValue); err != nil {
		t.Fatalf("expected no error setting %q but was: %v", wantValue, err)
	}

	for i, want := range []string{wantContent, wantValue} {
		if v[i] != want {
			t.Fatalf("want %q but was %q", want, v[i])
		}
	}
}
