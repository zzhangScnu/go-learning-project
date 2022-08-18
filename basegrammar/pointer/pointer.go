package pointer

import (
	"encoding/json"
	"fmt"
)

var str = "[{\"id\":\"1\",\"name\":\"首版\",\"code\":\"FIRST_VER\"},{\"id\":\"2\",\"name\":\"二版\",\"code\":\"REPEAT_VER\"},{\"id\":\"3\",\"name\":\"三版\",\"code\":\"THREE_VER\"},{\"id\":\"4\",\"name\":\"四版\",\"code\":\"FOUR_VER\"}]"

type IdNamePair struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func deserialize() []IdNamePair {
	var idNamePairList []IdNamePair
	if err := json.Unmarshal([]byte(str), &idNamePairList); err != nil {
		fmt.Printf("got exception: %v", err)
	}
	return idNamePairList
}

func ConvertToPointerMap() {
	idNamePairList := deserialize()
	idConfigPair := make(map[string]*IdNamePair)
	fmt.Printf("address of map is: %p\n", &idNamePairList)
	fmt.Printf("address of map[0] is: %p\n", &idNamePairList[0])
	fmt.Printf("address of map[1] is: %p\n", &idNamePairList[1])
	fmt.Printf("address of map[2] is: %p\n", &idNamePairList[2])
	fmt.Printf("address of map[3] is: %p\n", &idNamePairList[3])
	for _, item := range idNamePairList { // item是固定的一个变量，在这里就是值拷贝了
		fmt.Printf("address of %v is: %p\n", item, &item)
		idConfigPair[item.Id] = &item // 所以这里一直都是取的同一个变量的引用，最后map里面都是相同的值，就是最后一次遍历的对象
	}
	bytes, err := json.Marshal(idConfigPair)
	if err != nil {
		return
	}
	fmt.Printf("map=%v", string(bytes))
}

func ConvertToPointerMap2() {
	idNamePairList := deserialize()
	idConfigPair := make(map[string]*IdNamePair)
	for i := range idNamePairList {
		fmt.Printf("address of %v is: %p\n", idNamePairList[i], &idNamePairList[i])
		idConfigPair[idNamePairList[i].Id] = &idNamePairList[i]
	}
	bytes, err := json.Marshal(idConfigPair)
	if err != nil {
		return
	}
	fmt.Printf("map=%v", string(bytes))
}
