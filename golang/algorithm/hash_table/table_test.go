package hashtable

import (
	"fmt"
	"testing"
)

// hash函数的测试
func TestHashFunc(t *testing.T) {
	keys := []string{"hi", "my", "friend", "I", "love", "you", "my", "apple"}
	for _, key := range keys {
		t.Logf("xxhash('%s')=%d\n", key, hashAlgorithm([]byte(key)))
	}
}

func TestHashMap(t *testing.T) {
	hashMap := NewHashMap(16)

	// 放35个值
	for i := 0; i < 35; i++ {
		hashMap.Put(fmt.Sprintf("%d", i), fmt.Sprintf("v%d", i))
	}
	t.Log("cap: ", hashMap.capaticy, "len: ", hashMap.length)
	// 打印
	hashMap.Range()
	// 查询
	key := "4"
	value, ok := hashMap.Get(key)
	if ok {
		t.Logf("4=%v", value)
	} else {
		t.Errorf("key %s not found", key)
	}
	//删除
	hashMap.Delete(key)
	t.Log("cap: ", hashMap.capaticy, "len: ", hashMap.length)
	value, ok = hashMap.Get(key)
	if ok {
		t.Logf("4=%v", value)
	} else {
		t.Errorf("key %s not found", key)
	}
}
