package main

type Entry struct {
	key   string
	value int
	next  *Entry
}

type HashMap struct {
	buckets []*Entry
	size    int
}

func NewHashMap() *HashMap {
	return &HashMap{
		buckets: make([]*Entry, 10),
		size:    10,
	}
}

func (h *HashMap) Get(key string) (int, bool) {
	index := h.hash(key)
	entry := h.buckets[index]

	if entry == nil {
		return 0, false
	}

	for entry != nil {
		if entry.key == key {
			return entry.value, true
		}

		entry = entry.next
	}

	return 0, false
}

func (h *HashMap) Put(key string, value int) {
	index := h.hash(key)
	entry := h.buckets[index]

	if entry == nil {
		h.buckets[index] = &Entry{key: key, value: value}
		return
	}

	for entry.next != nil {
		if entry.key == key {
			entry.value = value
			return
		}

		entry = entry.next
	}

	if entry.key == key {
		entry.value = value
	} else {
		entry.next = &Entry{key: key, value: value}
	}
}

func (h *HashMap) hash(key string) int {
	hash := 0

	for i := 0; i < len(key); i++ {
		hash = (hash*31 + int(key[i])) % h.size
	}

	return hash
}
