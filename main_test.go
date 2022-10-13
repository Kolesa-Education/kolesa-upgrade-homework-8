package main

import (
	"github.com/Kolesa-Education/kolesa-upgrade-homework-8/card"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_cardsToRepresentation(t *testing.T) {
	t.Run("valid cards", func(t *testing.T) {
		card1, err := card.New(card.SuitDiamonds, card.FaceAce)
		require.NoError(t, err)
		card2, err := card.New(card.SuitSpades, card.FaceJack)
		require.NoError(t, err)
		card3, err := card.New(card.SuitSpades, card.Face10)
		require.NoError(t, err)

		cards := []card.Card{*card1, *card2, *card3}
		representations := cardsToRepresentations(cards)
		assert.Equal(t, []string{"♦A", "♠J", "♠10"}, representations)
	})

	t.Run("invalid cards produce empty representations", func(t *testing.T) {
		card1 := card.Card{
			Face: "invalid",
			Suit: card.SuitSpades,
		}

		card2 := card.Card{
			Face: card.FaceAce,
			Suit: "invalid",
		}

		card3 := card.Card{
			Face: card.FaceKing,
			Suit: "invalid",
		}

		cards := []card.Card{card1, card2, card3}
		representations := cardsToRepresentations(cards)
		assert.Equal(t, []string{"", "", ""}, representations)
	})
}
