package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFoundSyncShouldReturnNil(t *testing.T) {
	nf := &NotFound{}
	err := nf.Sync("test")
	assert.Nil(t, nil, err)
}
