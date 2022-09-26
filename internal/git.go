package internal

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"gitlab-tag-hook/config"
	"io"
)

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func GetTags(r *git.Repository) (string, error) {
	iter, err := r.Tags()
	if err != nil {
		return "", err
	}

	var ref *plumbing.Reference
	var latestTags string
	for {
		ref, err = iter.Next()
		if errors.Is(err, io.EOF) {
			err = errors.New("no tags found")
			break
		} else {
			latestTags = ref.Name().Short()
		}
	}

	if errors.Is(err, config.NOTAG_FOUND) {
		return config.DEFAULT_VAR, nil
	}

	return latestTags, nil

}

func ConvSemVersion(major string, minor string, patch string) string {
	return major + "." + minor + "." + patch
}
