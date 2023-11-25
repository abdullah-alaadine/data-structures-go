package hashtable

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
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
