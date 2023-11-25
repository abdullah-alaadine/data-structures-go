package hashtable

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

func (b *bucket) insert(k string) {
	newBucketNode := &bucketNode{key: k, next: b.head}
	b.head = newBucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
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
