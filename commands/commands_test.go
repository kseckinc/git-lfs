package commands

import (
	"testing"

	"github.com/git-lfs/git-lfs/v3/config"
	"github.com/stretchr/testify/assert"
)

var (
	testcfg = config.NewFrom(config.Values{
		Git: map[string][]string{
			"lfs.fetchinclude": []string{"/default/include"},
			"lfs.fetchexclude": []string{"/default/exclude"},
		},
	})
)

func TestDetermineIncludeExcludePathsReturnsCleanedPaths(t *testing.T) {
	inc := "/some/include"
	exc := "/some/exclude"
	i, e := determineIncludeExcludePaths(testcfg, &inc, &exc, true)

	assert.Equal(t, []string{"/some/include"}, i)
	assert.Equal(t, []string{"/some/exclude"}, e)
}

func TestDetermineIncludeExcludePathsReturnsEmptyPaths(t *testing.T) {
	inc := ""
	exc := ""
	i, e := determineIncludeExcludePaths(testcfg, &inc, &exc, true)

	assert.Empty(t, i)
	assert.Empty(t, e)
}

func TestDetermineIncludeExcludePathsReturnsDefaultsWhenAbsent(t *testing.T) {
	i, e := determineIncludeExcludePaths(testcfg, nil, nil, true)

	assert.Equal(t, []string{"/default/include"}, i)
	assert.Equal(t, []string{"/default/exclude"}, e)
}

func TestDetermineIncludeExcludePathsReturnsNothingWhenAbsent(t *testing.T) {
	i, e := determineIncludeExcludePaths(testcfg, nil, nil, false)

	assert.Empty(t, i)
	assert.Empty(t, e)
}
