package internal

// Depreciator calculates the quality of an item based upon the sell-by date and the previous quality.
// The quality it returns is the quality after one day.
type Depreciator interface {
	Depreciate(prevSell int, prevQuality int) (quality int)
}

// qualityBounder adjusts a quality amount to a specific range. If a depreciator implements qualityBounder
// that qualityBounder will be used over the default.
type QualityBounder interface {
	Bound(quality int) int
}

// defaultQualityBounder ensures the quality is between [0,50].
type DefaultQualityBounder struct{}

func (DefaultQualityBounder) Bound(quality int) int {
	if quality > 50 {
		return 50
	}

	if quality < 0 {
		return 0
	}

	return quality
}

// defaultDepreciator decreases an items quality by 1
type DefaultDepreciator struct{}

func (d DefaultDepreciator) Depreciate(sellIn int, prevQuality int) (quality int) {
	if prevQuality <= 0 {
		return 0
	}

	return prevQuality - 1
}

// betterWithAgeDepreciator increases in quality as sellBy decreases.
type BetterWithAgeDepreciator struct{}

func (d BetterWithAgeDepreciator) Depreciate(sellIn int, prevQuality int) (quality int) {
	return prevQuality + 1
}

type ConjuredDepreciator struct {
	Depreciator
}

func (d ConjuredDepreciator) Depreciate(sellIn int, prevQuality int) (quality int) {
	newQuality := d.Depreciator.Depreciate(sellIn, prevQuality)
	// conjured items depreciate twice as fast so calculate the difference and double it
	diff := newQuality - prevQuality
	newQuality = newQuality + diff
	return newQuality
}

// passDepreciator increases quality as the event date gets closer
type PassDepreciator struct{}

// Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
// Quality drops to 0 after the concert
func (d PassDepreciator) Depreciate(sellIn int, prevQuality int) (quality int) {
	// event has already happened so the pass is worthless
	if sellIn < 0 {
		return 0
	}

	if sellIn < 5 {
		return prevQuality + 3
	}

	if sellIn < 10 {
		return prevQuality + 2
	}

	return prevQuality + 1
}

type LegendaryDepreciator struct{}

// Depreciate always returns the same quality since legendary items never depreciate.
func (d LegendaryDepreciator) Depreciate(sellIn int, prevQuality int) (quality int) {
	return prevQuality
}

// Bound returns the items quality as legendary items have no bound
func (d LegendaryDepreciator) Bound(quality int) int {
	return quality
}
