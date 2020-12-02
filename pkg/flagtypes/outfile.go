package flagtypes

import "os"

const (
	stderr = "stderr"
	stdout = "stdout"
)

// OutFile holds output file
type OutFile struct{ *os.File }

func (v OutFile) String() string {
	if v.File == nil {
		return ""
	}
	return v.Name()
}

// Set sets output file
func (v *OutFile) Set(value string) (err error) {
	switch value {
	case stderr:
		v.File = os.Stderr
	case stdout:
		v.File = os.Stdout
	default:
		v.File, err = os.Create(value)
	}
	return
}
