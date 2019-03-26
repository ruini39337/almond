package ringTest

import (
	"container/ring"
	"fmt"
)

func RingTest(){
	const rLen = 3

	r := ring.New(rLen)

	for i := 0 ;i < rLen; i++ {
		r.Value = i
		r = r.Next()
	}

	printRing := func(v interface{}) {
		fmt.Println(v," ")
	}



	//将r之后的第二个元素的值乘2
	r.Move(2).Value = r.Move(2).Value.(int) * 2
	r.Do(printRing)
	fmt.Println( )

	//删除R与R+2之间的元素，既删除R+1
	//返回删除的元素组成的ring的指针
	result := r.Link(r.Move(2))

	r.Do(printRing)
	fmt.Println( )

	result.Do(printRing)
	fmt.Println( )

	another := ring.New(rLen)
	another.Value = 7
	another.Next().Value = 8 // 给 another + 1 表示的元素赋值，即第二个元素
	another.Prev().Value = 9 // 给 another - 1 表示的元素赋值，即第三个元素
	another.Do(printRing) // 7 8 9
	fmt.Println()

	// 插入another到r后面，返回插入前r的下一个元素
	result = r.Link(another)
	r.Do(printRing) // 0 7 8 9 4
	fmt.Println()

	result.Do(printRing) // 4 0 7 8 9
	fmt.Println()

	// 删除r之后的三个元素，返回被删除元素组成的ring的指针
	result = r.Unlink(3)

	r.Do(printRing) // 0 4
	fmt.Println()

	result.Do(printRing) //7 8 9
	fmt.Println()

}
