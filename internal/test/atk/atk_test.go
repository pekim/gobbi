package atk

import (
	"github.com/pekim/gobbi/lib/atk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanBuild(t *testing.T) {
	assert.Equal(t, atk.HyperlinkStateFlags(1), atk.ATK_HYPERLINK_IS_INLINE)
}
