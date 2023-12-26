package storage

import "time"

type Entry struct {
	Key        string
	Expiration time.Time
	Index      int
}

type TTLHeap []*Entry

func (h TTLHeap) Len() int           { return len(h) }
func (h TTLHeap) Less(i, j int) bool { return h[i].Expiration.Before(h[j].Expiration) }
func (h TTLHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i]; h[i].Index, h[j].Index = i, j }
func (h *TTLHeap) Push(x interface{}) {
	entry := x.(*Entry)
	entry.Index = len(*h)
	*h = append(*h, entry)
}

func (h *TTLHeap) Pop() interface{} {
	old := *h
	n := len(old)
	entry := old[n-1]
	entry.Index = -1 // for safety
	*h = old[0 : n-1]
	return entry
}
