package gildedrose

import (
	"strings"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose/internal"
)

type itemUpdater struct {
	// i is the Item to update.
	i *Item
	// depreciator will calculate the quality depreciation
	depreciator internal.Depreciator
	// sellInPolicy updates the sellIn date for the item in question.
	sellInPolicy func(int) int
}

func (u *itemUpdater) update() *Item {
	newSellIn := u.updateSellIn()
	q := u.UpdateQuality(newSellIn)
	u.i.Quality = q
	u.i.SellIn = newSellIn
	return u.i
}

func (u itemUpdater) updateSellIn() int {
	sellIn := u.i.SellIn
	if u.sellInPolicy != nil {
		return u.sellInPolicy(sellIn)
	}

	return sellIn - 1
}

func (u itemUpdater) UpdateQuality(sellIn int) int {
	quality := u.i.Quality
	newQuality := u.depreciator.Depreciate(sellIn, quality)

	diff := newQuality - quality
	// When the sell by date is < 0 depreciation multiplier is 2x
	if sellIn < 0 {
		newQuality += diff
	}

	newQuality = u.boundQuality(newQuality)

	return newQuality
}

func (u *itemUpdater) boundQuality(q int) int {
	b, ok := u.depreciator.(internal.QualityBounder)
	if !ok {
		return internal.DefaultQualityBounder{}.Bound(q)
	}

	return b.Bound(q)
}

// classifier classifies Items into an item with a corresponding depreciator
type classifier struct{}

func (c classifier) classify(i *Item) itemUpdater {
	item := itemUpdater{i: i}
	item.depreciator = c.depreciator(i)
	if c.isLegendary(i.Name) {
		item.sellInPolicy = func(i int) int {
			return i
		}
	}
	return item
}

// depreciator returns the depreciator to be used for the item
func (c classifier) depreciator(i *Item) internal.Depreciator {
	name := i.Name
	var d internal.Depreciator
	switch {
	case c.isPass(name):
		d = internal.PassDepreciator{}
	case c.isLegendary(name):
		d = internal.LegendaryDepreciator{}
	case c.isBetterWithAge(name):
		d = internal.BetterWithAgeDepreciator{}
	default:
		d = internal.DefaultDepreciator{}
	}

	if strings.HasPrefix(name, "Conjured") {
		d = internal.ConjuredDepreciator{Depreciator: d}
	}

	return d
}

func (c classifier) isLegendary(name string) bool {
	return strings.HasPrefix(name, "Sulfuras")
}
func (c classifier) isPass(name string) bool {
	return strings.HasPrefix(name, "Backstage pass")
}
func (c classifier) isBetterWithAge(name string) bool {
	return name == "Aged Brie"
}

// Do Not Change
type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	c := classifier{}
	for i := 0; i < len(items); i++ {
		item := items[i]
		updater := c.classify(item)
		updater.update()
	}
}
