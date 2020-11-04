package medium

/**
剑指 Offer 59 - II. 队列的最大值

https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof/

请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

示例 1：

输入:
["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
[[],[1],[2],[],[],[]]
输出: [null,null,null,2,1,2]
示例 2：

输入:
["MaxQueue","pop_front","max_value"]
[[],[],[]]
输出: [null,-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type MaxQueue struct {
	Max []int
	First *QueueNode
	Last *QueueNode
}

type QueueNode struct {
	Val int
	Next *QueueNode
}

func Constructor() MaxQueue {
	return MaxQueue{Max: []int{},First: nil, Last: nil}
}


func (this *MaxQueue) Max_value() int {
	if len(this.Max)==0{
		return -1
	}
	return this.Max[0]
}


func (this *MaxQueue) Push_back(value int)  {
	if this.First == nil{
		temp := &QueueNode{Val: value}
		this.Last = temp
		this.First = temp
		this.Max = []int{value}
	}else{
		this.Last.Next = &QueueNode{Val: value}
		this.Last = this.Last.Next
		i := len(this.Max)
		for i>0{
			if this.Max[i-1]>=value{
				this.Max = append(this.Max[0:i], value)
				break
			}
			i--
		}
		if i == 0{
			this.Max = []int{value}
		}
	}
}


func (this *MaxQueue) Pop_front() int {
	temp := this.First
	if temp == nil{
		return -1
	}
	this.First = this.First.Next
	if this.First == nil{
		this.Last = this.First
		this.Max = []int{}
	}else if temp.Val == this.Max[0]{
		this.Max = this.Max[1:]
	}
	return temp.Val
}