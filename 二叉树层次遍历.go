import "container/list"

var (
    help = []*list.List{}
)

func levelOrder(root *TreeNode) [][]int {   
    help = []*list.List{}  // clean 
    
    helper(root, 0)
    ret := make([][]int, len(help))

    for k, lv := range help {
        for e := lv.Front(); e != nil; e = e.Next() {
            ret[k] = append(ret[k], e.Value.(int))
        }
    }
    return ret
}

func helper(root *TreeNode, level int) {
    if root == nil {
        return
    }
    if len(help) <= level {
        help = append(help, list.New())
    }
    help[level].PushBack(root.Val)
 

    helper(root.Left, level+1)
    helper(root.Right, level+1)
}
