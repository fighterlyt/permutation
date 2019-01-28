package permutation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	testintdata   = []int{1, 2, 3, 4}
	testindex     = 2
	expectedperms = [][]int{
		{1, 2, 3, 4},
		{1, 2, 4, 3},
		{1, 3, 2, 4},
		{1, 3, 4, 2},
		{1, 4, 2, 3},
		{1, 4, 3, 2},
		{2, 1, 3, 4},
		{2, 1, 4, 3},
		{2, 3, 1, 4},
		{2, 3, 4, 1},
		{2, 4, 1, 3},
		{2, 4, 3, 1},
		{3, 1, 2, 4},
		{3, 1, 4, 2},
		{3, 2, 1, 4},
		{3, 2, 4, 1},
		{3, 4, 1, 2},
		{3, 4, 2, 1},
		{4, 1, 2, 3},
		{4, 1, 3, 2},
		{4, 2, 1, 3},
		{4, 2, 3, 1},
		{4, 3, 1, 2},
		{4, 3, 2, 1}}
)

func Test_Length(t *testing.T) {
	p, err := NewPerm(testintdata, nil)
	require.NoError(t, err)
	require.Equal(t, len(testintdata), p.Length())
}

func Test_Reset(t *testing.T) {
	p, err := NewPerm(testintdata, nil)
	require.NoError(t, err)

	newindex, err := p.MoveIndex(testindex)
	require.NoError(t, err)
	require.Equal(t, testindex, newindex)

	p.Reset()
	require.Equal(t, 0, p.Index())
}

func Test_MoveIndex(t *testing.T) {
	p, err := NewPerm(testintdata, nil)
	require.NoError(t, err)

	newindex, err := p.MoveIndex(testindex)
	require.NoError(t, err)
	require.Equal(t, testindex, newindex)
	require.Equal(t, newindex, p.Index()+1)

	// Test that an error occurs if we go beyond the end of the index
	newindex, err = p.MoveIndex(len(testintdata) + 1)
	require.Error(t, err)

	// Test that an error occurs if we specify a negative index
	newindex, err = p.MoveIndex(-1)
	require.Error(t, err)
}

func Test_NextN(t *testing.T) {
	p, err := NewPerm(testintdata, nil)
	require.NoError(t, err)
	perms := p.NextN(24)
	require.Equal(t, expectedperms, perms)
}

func Test_Index(t *testing.T) {
	p, err := NewPerm(testintdata, nil)
	require.NoError(t, err)
	require.Equal(t, 0, p.Index())

	newindex, err := p.MoveIndex(1)
	require.NoError(t, err)
	require.Equal(t, 1, newindex)

	newindex, err = p.MoveIndex(2)
	require.NoError(t, err)
	require.Equal(t, 2, newindex)

	newindex, err = p.MoveIndex(3)
	require.NoError(t, err)
	require.Equal(t, 3, newindex)

	newindex, err = p.MoveIndex(0)
	require.NoError(t, err)
	require.Equal(t, 0, newindex)
}
func Test_Next(t *testing.T) {
	p, err := NewPerm(testintdata, nil)
	require.NoError(t, err)
	require.Equal(t, 0, p.Index())

	perm, err := p.Next()
	require.NoError(t, err)
	require.Equal(t, expectedperms[0], perm)
	require.Equal(t, 1, p.Index())

	perm, err = p.Next()
	require.NoError(t, err)
	require.Equal(t, expectedperms[1], perm)
	require.Equal(t, 2, p.Index())

	perm, err = p.Next()
	require.NoError(t, err)
	require.Equal(t, expectedperms[2], perm)
	require.Equal(t, 3, p.Index())
}

func Test_Left(t *testing.T) {
	p, err := NewPerm(testintdata, nil)
	require.NoError(t, err)
	left := p.Left()
	require.Equal(t, 24, left)
}
