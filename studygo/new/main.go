package main

import "fmt"
func PrintInterface(x interface{}){
	if x==nil{
		fmt.Println("empty interface")
		return
	}

}

type Node struct {
	next *Node
	val int
}
func getval(root *Node,a int)(*Node){
	if root==nil{
		return nil
	}else{
		for root!=nil{
			if root.val==a{
				return root
			}
			root=root.next
		}
	}
	return nil
}
 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
func pathSum(node *TreeNode, targetSum int)(ans [][]int) {
	var path=[]int{}
    var dfs func(*TreeNode,int)
	dfs=func(root *TreeNode,left int){
		if root==nil{
			return
		}
		left-=root.Val
		path=append(path,root.Val)
		defer func(){
			path=path[:len(path)-1]
		}()
		if root.Left==nil&&root.Right==nil&&left==0 {
			ans=append(ans,append([]int(nil),path...))
			return
		}
		dfs(root.Left,left)
		dfs(root.Right,left)
	}
	dfs(node,targetSum)
	return
}
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}
func main() {
	//root:=Node{
	//	&Node{val: 3,next: nil},
	//	6,
	//}
	//tmp:=getval(&root,3)
	//fmt.Println(tmp.val)
	c:=make(chan int)
	go func() {
		defer fmt.Println("gorouting 结束")
		fmt.Println("gorouting 运行")
		c<-666
	}()
	num:=<-c
	fmt.Printf("num:%d\n",num)
	fmt.Println("main gorouting 结束")
}
