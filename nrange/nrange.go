package nrange

/// Integer Range ///

type NRange struct{ begin, end int }

// Inclusive Range Constructor
// From begin (inclusive) to end (inclusive)
func IRange(begin, end int) *NRange {
	return &NRange{
		begin: begin,
		end:   end,
	}
}

// Exclusive Range Constructor
// From begin (inclusive) to end (exclusive)
func XRange(begin, end int) *NRange {
	return &NRange{
		begin: begin,
		end:   end - 1,
	}
}

func (this *NRange) Begin() int {
	return this.begin
}

func (this *NRange) End() int {
	return this.end
}

func (this *NRange) Length() int {
	return this.end - this.begin + 1
}

func (this *NRange) InRange(n int) bool {
	return n >= this.begin && n <= this.end
}

func (this *NRange) Enumerate() []int {
	fin := make([]int, this.Length())
	for i := this.begin; i <= this.end; i++ {
		fin[i-this.begin] = i
	}
	return fin
}

/// Iterator ///
// Implements the iter.Able interface from github.com/odysseus/iter

type NRangeIterator struct {
	nrange *NRange
	next   int
}

func Iterator(r *NRange) *NRangeIterator {
	return &NRangeIterator{
		nrange: r,
		next:   r.begin,
	}
}

func (it *NRangeIterator) Next() int {
	it.next++
	return it.next - 1
}

func (it *NRangeIterator) Peek() int {
	return it.next
}

func (it *NRangeIterator) HasNext() bool {
	if it.next <= it.nrange.end+1 {
		return true
	} else {
		return false
	}
}
