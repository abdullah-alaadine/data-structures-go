package hashtable

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

func (ht *HashTable) Insert(k string, v any) {
	idx := hash(k) % ArraySize
	if v := ht.array[idx].get(k); v != nil {
		ht.array[idx].put(k, v)
		return
	}
	ht.array[idx].insert(k, v)
}

type bucket struct {
	head *bucketNode
}

func (b *bucket) insert(k string, v any) {
	newBucketNode := &bucketNode{key: k, value: v, next: b.head}
	b.head = newBucketNode
}

func (b *bucket) put(k string, v any) {
	if b.head == nil {
		return
	}
	curr := b.head
	for curr != nil {
		if curr.key == k {
			curr.value = v
			return
		}
		curr = curr.next
	}
}

func (b *bucket) get(k string) any {
	currentBucket := b.head
	for currentBucket != nil {
		if currentBucket.key == k {
			return currentBucket.value
		}
		currentBucket = currentBucket.next
	}
	return nil
}

func (b *bucket) delete(k string) {
	if b.head == nil {
		return
	}
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	currentBucket := b.head
	for currentBucket.next != nil {
		if currentBucket.next.key == k {
			currentBucket.next = currentBucket.next.next
			return
		}
		currentBucket = currentBucket.next
	}
}

type bucketNode struct {
	key   string
	value any
	next  *bucketNode
}

func New() *HashTable {
	ht := &HashTable{}
	for i := range ht.array {
		ht.array[i] = &bucket{}
	}
	return ht
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum
}
