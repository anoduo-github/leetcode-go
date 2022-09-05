package question

//146. LRU 缓存
type LRUCache struct {
	data     map[int]int //数据
	keys     []int       //循环队列
	size     int         //大小
	capacity int         //容量
}

func Constructor146(capacity int) LRUCache {
	return LRUCache{
		data:     make(map[int]int, 0),
		keys:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.data[key]; ok {
		//更新键值队列
		index := 0
		//找到下标
		for i, v := range this.keys {
			if v == key {
				index = i
			}
		}
		//往前移一位
		for i := index; i < this.size-1; i++ {
			this.keys[i] = this.keys[i+1]
		}
		//末尾赋值
		this.keys[this.size-1] = key
		return value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.data[key]; ok {
		this.data[key] = value
		//更新键值队列
		index := 0
		//找到下标
		for i, v := range this.keys {
			if v == key {
				index = i
			}
		}
		//往前移一位
		for i := index; i < this.size-1; i++ {
			this.keys[i] = this.keys[i+1]
		}
		//末尾赋值
		this.keys[this.size-1] = key
		return
	}
	if this.size < this.capacity {
		//存值
		this.data[key] = value
		//存键队列
		this.keys[this.size] = key
		//大小加一
		this.size++
	} else {
		tempKey := this.keys[0]
		//剔除最少使用的，即键队列头
		for i := 0; i < this.size-1; i++ {
			this.keys[i] = this.keys[i+1]
		}
		//删除key
		delete(this.data, tempKey)
		//添加新的键值
		this.data[key] = value
		//存键值队列
		this.keys[this.size-1] = key
	}
}
