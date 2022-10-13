package card

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValidSuit(t *testing.T) {
	assert.True(t, isValidSuit(SuitDiamonds))
	assert.True(t, isValidSuit(SuitSpades))
	assert.True(t, isValidSuit(SuitClubs))
	assert.True(t, isValidSuit(SuitHearts))

	assert.False(t, isValidSuit("Invalid"))
}

func Test_isValidFace(t *testing.T) {
	assert.True(t, isValidFace(Face2))
	assert.True(t, isValidFace(Face3))
	assert.True(t, isValidFace(Face4))
	assert.True(t, isValidFace(Face5))
	assert.True(t, isValidFace(Face6))
	assert.True(t, isValidFace(Face7))
	assert.True(t, isValidFace(Face8))
	assert.True(t, isValidFace(Face9))
	assert.True(t, isValidFace(Face10))
	assert.True(t, isValidFace(FaceJack))
	assert.True(t, isValidFace(FaceQueen))
	assert.True(t, isValidFace(FaceKing))
	assert.True(t, isValidFace(FaceAce))

	assert.False(t, isValidFace("Invalid"))
}
