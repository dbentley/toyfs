package path

import (
	"fmt"
	"strings"
)

func Split(path string) ([]string, error) {
	if err := MustBeAbs(path); err != nil {
		return nil, err
	}

	if path == "/" {
		return []string{}, nil
	}

	path = path[1:] // lop off first / or else /foo would split into ["", "foo"]

	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	return strings.Split(path, "/"), nil
}

func Join(beginning, rest string) (string, error) {
	if err := MustBeAbs(beginning); err != nil {
		return "", err
	}

	if IsAbs(rest) {
		return "", fmt.Errorf("Join: rest must be relative; got %q", rest)
	}

	if !strings.HasSuffix(beginning, "/") {
		beginning = beginning + "/"
	}

	return beginning + rest, nil
}

func IsAbs(path string) bool {
	return strings.HasPrefix(path, "/")
}

func MustBeAbs(path string) error {
	if !IsAbs(path) {
		return fmt.Errorf("toyfs paths must start with '/'; got %q", path)
	}
	return nil
}
