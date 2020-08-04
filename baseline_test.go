package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	n := NewLinkedList(10)
	require.NotNil(t, n)
}

func TestAppend(t *testing.T) {
	require.Equal(t, 1, 1)
	fmt.Println("!")
}
