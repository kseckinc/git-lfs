//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

package tools

import (
	"io"

	"github.com/git-lfs/git-lfs/v3/errors"
)

func CheckCloneFileSupported(dir string) (supported bool, err error) {
	return false, errors.New("unsupported platform")
}

func CloneFile(writer io.Writer, reader io.Reader) (bool, error) {
	return false, nil
}

func CloneFileByPath(_, _ string) (bool, error) {
	return false, nil
}
