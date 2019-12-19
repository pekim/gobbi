package generate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersions(t *testing.T) {
	testVersions := func() *versions {
		vv := &versions{}
		vv.add(versionNew("1.3"))
		vv.add(versionNew(" "))
		vv.add(versionNew("1.3.2"))
		vv.add(versionNew("1.1.0"))
		vv.add(versionNew("1.1"))
		vv.add(versionNew("1.3.2"))
		vv.add(versionNew("1.3.1"))
		vv.add(versionNew(" "))
		return vv
	}

	t.Run("add", func(t *testing.T) {
		vv := testVersions()

		assert.Equal(t, 8, vv.len())
	})

	t.Run("dedupe", func(t *testing.T) {
		vv := testVersions()
		vv.dedupe()

		assert.Equal(t, 5, vv.len())
	})

	t.Run("sort", func(t *testing.T) {
		vv := testVersions()
		vv.sort()

		assert.Equal(t, "0.0.0 1.1.0 1.3.0 1.3.1 1.3.2", vv.String())
	})
}
