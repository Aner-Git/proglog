package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLog(t *testing.T) {
	l := NewLog()
	assert.NotNil(t, l)
}

func TestAddRecord(t *testing.T) {
	l := NewLog()

	r := Record{[]byte("foo"), 0}

	offset, err := l.Append(r)
	assert.Nil(t, err)

	got, err := l.Read(offset)
	assert.Nil(t, err)

	assert.Equal(t, r, got, "Records should be equal")
}

func TestInvaidOffset(t *testing.T) {
	l := NewLog()

	var offset uint64 = 6767
	got, err := l.Read(offset)
	assert.Equal(t, Record{}, got, "Records should be empty")

	assert.Equal(t, err, ErrOffsetNotFound, "Erros should be equal")
}
