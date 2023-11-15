package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/require"
)

func Test_Foo(t *testing.T) {
	items := []gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	t.Run("one day", func(t *testing.T) {
		items := toPointer(items)
		gildedrose.UpdateQuality(items)
		require.EqualValues(
			t,
			[]*gildedrose.Item{
				{"+5 Dexterity Vest", 9, 19},
				{"Aged Brie", 1, 1},
				{"Elixir of the Mongoose", 4, 6},
				{"Sulfuras, Hand of Ragnaros", 0, 80},
				{"Sulfuras, Hand of Ragnaros", -1, 80},
				{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
				{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
				{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
				{"Conjured Mana Cake", 2, 4},
			}, items)
	})

	t.Run("5 days", func(t *testing.T) {
		items := toPointer(items)
		for range make([]any, 5) {
			gildedrose.UpdateQuality(items)
		}
		require.EqualValues(
			t,
			[]*gildedrose.Item{
				{"+5 Dexterity Vest", 5, 15},
				{"Aged Brie", -3, 8},
				{"Elixir of the Mongoose", 0, 2},
				{"Sulfuras, Hand of Ragnaros", 0, 80},
				{"Sulfuras, Hand of Ragnaros", -1, 80},
				{"Backstage passes to a TAFKAL80ETC concert", 10, 25},
				{"Backstage passes to a TAFKAL80ETC concert", 5, 50},
				{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
				{"Conjured Mana Cake", -2, 0},
			}, items)
	})

	t.Run("10 days", func(t *testing.T) {
		items := toPointer(items)
		for range make([]any, 10) {
			gildedrose.UpdateQuality(items)
		}
		require.EqualValues(
			t,
			[]*gildedrose.Item{
				{"+5 Dexterity Vest", 0, 10},
				{"Aged Brie", -8, 18},
				{"Elixir of the Mongoose", -5, 0},
				{"Sulfuras, Hand of Ragnaros", 0, 80},
				{"Sulfuras, Hand of Ragnaros", -1, 80},
				{"Backstage passes to a TAFKAL80ETC concert", 5, 35},
				{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
				{"Backstage passes to a TAFKAL80ETC concert", -5, 0},
				{"Conjured Mana Cake", -7, 0},
			}, items)
	})
}

func toPointer(item []gildedrose.Item) []*gildedrose.Item {
	cpy := make([]gildedrose.Item, len(item))
	copy(cpy, item)
	ret := make([]*gildedrose.Item, 0)
	for _, i := range cpy {
		i := i
		ret = append(ret, &i)
	}

	return ret
}
