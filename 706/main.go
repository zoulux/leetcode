package main

func main() {

}

type MyHashMap struct {
	bucketKeys   [10000][]int
	bucketValues [10000][]int
}

func Constructor() MyHashMap {
	bk := [10000][]int{}
	bv := [10000][]int{}

	for i := 0; i < len(bk); i++ {
		bk[i] = make([]int, 0)
		bv[i] = make([]int, 0)
	}
	return MyHashMap{
		bucketKeys:   bk,
		bucketValues: bv,
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % 1000
}

func (this *MyHashMap) Put(key int, value int) {
	b := this.hash(key)

	target := -1
	for i, v := range this.bucketKeys[b] {
		if v == key {
			target = i
		}
	}
	if target != -1 {
		this.bucketValues[b][target] = value
	} else {
		this.bucketKeys[b] = append(this.bucketKeys[b], key)
		this.bucketValues[b] = append(this.bucketValues[b], value)
	}

}

func (this *MyHashMap) Get(key int) int {
	b := this.hash(key)
	target := -1
	for i, v := range this.bucketKeys[b] {
		if v == key {
			target = i
		}
	}
	if target != -1 {
		return this.bucketValues[b][target]
	}
	return target
}

func (this *MyHashMap) Remove(key int) {
	b := this.hash(key)

	target := -1
	for i, v := range this.bucketKeys[b] {
		if v == key {
			target = i
		}
	}
	if target != -1 {
		k1 := this.bucketKeys[b][:target]
		k2 := this.bucketKeys[b][target+1:]
		this.bucketKeys[b] = append(k1, k2...)

		v1 := this.bucketValues[b][:target]
		v2 := this.bucketValues[b][target+1:]
		this.bucketValues[b] = append(v1, v2...)
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
