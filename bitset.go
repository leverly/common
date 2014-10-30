package common

// Word size of a bit set
const wordSize = uint(64)
const log2WordSize = uint(6)

// The zero value of a BitSet is an empty set of length 0.
type BitSet struct {
	length uint
	set    []uint64
}

func NewBitSet(length uint) *BitSet {
	return &BitSet{length, make([]uint64, wordsNeeded(length))}
}

func (this *BitSet) Len() uint {
	return this.length
}

// Clone this BitSet
func (this *BitSet) Clone() *BitSet {
	c := NewBitSet(this.length)
	if this.set != nil {
		copy(c.set, this.set)
	}
	return c
}

// Set bit i to 1
func (this *BitSet) Set(i uint) *BitSet {
	this.extendSetMaybe(i)
	this.set[i>>log2WordSize] |= 1 << (i & (wordSize - 1))
	return this
}

// Unset bit i to 0
func (this *BitSet) UnSet(i uint) *BitSet {
	if i >= this.length {
		return this
	}
	this.set[i>>log2WordSize] &^= 1 << (i & (wordSize - 1))
	return this
}

// Flip bit at i
func (this *BitSet) Flip(i uint) *BitSet {
	if i >= this.length {
		return this.Set(i)
	}
	this.set[i>>log2WordSize] ^= 1 << (i & (wordSize - 1))
	return this
}

// Clear entire BitSet
func (this *BitSet) Clear() *BitSet {
	if this != nil && this.set != nil {
		for i := range this.set {
			this.set[i] = 0
		}
	}
	return this
}

// Test the equvalence of two BitSets.
// False if they are of different sizes, otherwise true
// only if all the same bits are set
func (this *BitSet) Equal(c *BitSet) bool {
	if c == nil {
		return false
	}
	if this.length != c.length {
		return false
	}
	if this.length == 0 { // if they have both length == 0, then could have nil set
		return true
	}
	// testing for equality shoud not transform the bitset (no call to safeSet)
	for p, v := range this.set {
		if c.set[p] != v {
			return false
		}
	}
	return true
}

// Difference of base set and other set
// This is the BitSet equivalent of &^ (and not)
func (this *BitSet) Difference(compare *BitSet) (result *BitSet) {
	result = this.Clone() // clone b (in case b is bigger than compare)
	l := int(wordsNeeded(compare.length))
	if l > int(wordsNeeded(this.length)) {
		l = int(wordsNeeded(this.length))
	}
	for i := 0; i < l; i++ {
		result.set[i] = this.set[i] &^ compare.set[i]
	}
	return
}

// Intersection of base set and other set
// This is the BitSet equivalent of & (and)
func (this *BitSet) Intersection(compare *BitSet) (result *BitSet) {
	this, compare = sortByLength(this, compare)
	result = NewBitSet(this.length)
	for i, word := range this.set {
		result.set[i] = word & compare.set[i]
	}
	return
}

// Union of base set and other set
// This is the BitSet equivalent of | (or)
func (this *BitSet) Union(compare *BitSet) (result *BitSet) {
	this, compare = sortByLength(this, compare)
	result = compare.Clone()
	for i, word := range this.set {
		result.set[i] = word | compare.set[i]
	}
	return
}

func wordsNeeded(i uint) int {
	if i > ((^uint(0)) - wordSize + 1) {
		return int((^uint(0)) >> log2WordSize)
	}
	return int((i + (wordSize - 1)) >> log2WordSize)
}

func (this *BitSet) extendSetMaybe(i uint) {
	if i >= this.length {
		nsize := wordsNeeded(i + 1)
		if this.set == nil {
			this.set = make([]uint64, nsize)
		} else if len(this.set) < nsize {
			newset := make([]uint64, nsize)
			copy(newset, this.set)
			this.set = newset
		}
		this.length = i + 1
	}
}

func sortByLength(a *BitSet, b *BitSet) (ap *BitSet, bp *BitSet) {
	if a.length <= b.length {
		ap, bp = a, b
	} else {
		ap, bp = b, a
	}
	return
}
