package hashtable

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

func (b *bucket) insert(k string, v any) {
	newBucketNode := &bucketNode{key: k, value: v, next: b.head}
	b.head = newBucketNode
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
