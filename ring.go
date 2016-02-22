package main //引入包
import "container/ring"

func main() {
	//创建闭环，这里创建10个元素的闭环
	r := ring.New(10)
	//给闭环中的元素附值
	for i := 1; i <= r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	//循环打印闭环元素的值，这里的操作方法很像javascript
	r.Do(func(p interface{}) {
		println(p.(int))
	})
	//当前元素就是
	// r
	//当前元素的值就是
	// r.Value       //这里是 1
	//取得当前元素之后的第5个元素
	r5 := r.Move(5) //这里是 6
	println("------")
	r5.Do(func(p interface{}) {
		println(p.(int))
	})
	//链接 当前元素r与r5，相当于删除了r跟r5之间的元素，使 r.Next() == r5
	rl := r.Link(r5) //这样 r环里有 1 6 7 8 9 10
	// rl 环路里有  2 3 4 5
	println("------")
	rl.Do(func(p interface{}) {
		println(p.(int))
	})
	//把rl环加回到 r环原来的位置
	//要确保r的当前位置是 1，rl的当前位置为2
	rf := r.Link(rl)
	println("------")
	rf.Do(func(p interface{}) {
		println(p.(int))
	})
	//这样 r就变回了 1 2 3 4 5 6 7 8 9 10
	//rf则是 从 rl.Next()开始环，内容则与 r的一样。
	//6 7 8 9 10 1 2 3 4 5
}
