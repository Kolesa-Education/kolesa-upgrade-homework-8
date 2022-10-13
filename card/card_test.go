package card

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestNew(t *testing.T) {
	t.Run("valid creation of cards", func(t *testing.T) {
		suits := []string{SuitDiamonds, SuitSpades, SuitClubs, SuitHearts}
		faces := []string{Face2, Face3, Face4, Face5, Face6, Face7, Face8, Face9, Face10, FaceJack, FaceQueen, FaceKing, FaceAce}

		for _, suit := range suits {
			for _, face := range faces {
				c, err := New(suit, face)
				require.NoError(t, err)
				assert.Equal(t, c.Face, face)
				assert.Equal(t, c.Suit, suit)
			}
		}
	})

	t.Run("invalid face results in error", func(t *testing.T) {
		c, err := New(SuitHearts, "invalid")
		require.Error(t, err)
		assert.Nil(t, c)
	})

	t.Run("invalid suit results in error", func(t *testing.T) {
		c, err := New("invalid", FaceAce)
		require.Error(t, err)
		assert.Nil(t, c)
	})
}

func TestCard_SuitUnicode(t *testing.T) {
	t.Run("spades", func(t *testing.T) {
		c := Card{
			Suit: SuitSpades,
			Face: FaceAce,
		}
		unicode, err := c.SuitUnicode()
		require.NoError(t, err)
		assert.Equal(t, unicode, SuitSpadesUnicode)
	})

	t.Run("diamonds", func(t *testing.T) {
		c := Card{
			Suit: SuitDiamonds,
			Face: FaceAce,
		}
		unicode, err := c.SuitUnicode()
		require.NoError(t, err)
		assert.Equal(t, unicode, SuitDiamondsUnicode)
	})

	t.Run("hearts", func(t *testing.T) {
		c := Card{
			Suit: SuitHearts,
			Face: FaceAce,
		}
		unicode, err := c.SuitUnicode()
		require.NoError(t, err)
		assert.Equal(t, unicode, SuitHeartsUnicode)
	})

	t.Run("clubs", func(t *testing.T) {
		c := Card{
			Suit: SuitClubs,
			Face: FaceAce,
		}
		unicode, err := c.SuitUnicode()
		require.NoError(t, err)
		assert.Equal(t, unicode, SuitClubsUnicode)
	})

	t.Run("invalid suit", func(t *testing.T) {
		c := Card{
			Suit: "invalid",
			Face: FaceAce,
		}
		unicode, err := c.SuitUnicode()
		assert.Equal(t, "", unicode)
		require.Error(t, err)
	})
}
