package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	comp1 := &Competency{Title: "title", ID: "1", Levels: []*Level{{Summary: "summary"}}}
	comp1.ComputeHash()

	comp2 := &Competency{Title: "title", Levels: []*Level{{Summary: "summary"}}}
	comp2.ComputeHash()

	comp3 := &Competency{Title: "title diff", Levels: nil}
	comp3.ComputeHash()

	comp4 := &Competency{Title: "title", Levels: []*Level{{Summary: "summary", Improve: "improve"}}}
	comp4.ComputeHash()

	require.Equal(t, comp1.Hash, comp2.Hash)
	require.NotEqual(t, comp1.Hash, comp3.Hash)
	require.NotEqual(t, comp1.Hash, comp4.Hash)
}
