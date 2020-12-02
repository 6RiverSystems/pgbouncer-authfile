package flagtypes

import (
	"io/ioutil"
	"os"
	"strings"
)

// InputData contains strings read from files or string value if file does not exist.
type InputData []string

func (v InputData) String() string {
	return strings.Join(v, ",")
}

// Set tries to read file with given string value. If such a file
// does not exist the value itself is used.
func (v *InputData) Set(str string) error {
	if b, err := ioutil.ReadFile(str); err == nil {
		str = strings.TrimSuffix(string(b), "\n")
	} else if !os.IsNotExist(err) {
		return err
	}
	*v = append(*v, str)
	return nil
}
