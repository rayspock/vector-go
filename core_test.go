package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDotProductOnPerpendicularVectors(t *testing.T) {
	a := Vector{1., 0., 0.}
	b := Vector{0., 1., 0.}
	result := a.Dot(b)
	assert.Equal(t, 0., result, "a dot b")
	result = b.Dot(a)
	assert.Equal(t, 0., result, "b dot a")
}

func TestDotProductInSameDirection(t *testing.T) {
	a := Vector{1., 0., 0.}
	b := Vector{1., 0., 0.}
	result := a.Dot(b)
	assert.Equal(t, 1., result, "a dot b in same direction")
	result = b.Dot(a)
	assert.Equal(t, 1., result, "b dot a in same direction")
}

func TestDotProductInOppositeDirection(t *testing.T) {
	a := Vector{1., 0., 0.}
	b := Vector{-1., 0., 0.}
	result := a.Dot(b)
	assert.Equal(t, -1., result, "a dot b in opposite direction")
	result = b.Dot(a)
	assert.Equal(t, -1., result, "b dot a in opposite direction")
}

func TestCrossProductOnPerpendicularVectors(t *testing.T) {
	a := Vector{1., 0., 0.}
	b := Vector{0., 1., 0.}
	result := a.Cross(b)
	assert.Equal(t, Vector{0., 0., 1.}, result, "a cross b should be coming out of the thumb")
	result = b.Cross(a)
	assert.Equal(t, Vector{0., 0., -1.}, result, "b cross a should be reversing the sign of the product vector")
}

func TestCrossProductInSameDirection(t *testing.T) {
	a := Vector{1., 0., 0.}
	b := Vector{1., 0., 0.}
	result := a.Cross(b)
	assert.Equal(t, Vector{0., 0., 0.}, result, "a cross b in same direction")
	result = b.Cross(a)
	assert.Equal(t, Vector{0., 0., 0.}, result, "b cross a in same direction")
}

func TestCrossProductInOppositeDirection(t *testing.T) {
	a := Vector{1., 0., 0.}
	b := Vector{-1., 0., 0.}
	result := a.Cross(b)
	assert.Equal(t, Vector{0., 0., 0.}, result, "a cross b in opposite direction")
	result = b.Cross(a)
	assert.Equal(t, Vector{0., 0., 0.}, result, "b cross a in opposite direction")
}
