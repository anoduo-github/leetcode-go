package question

import "fmt"

//652. 寻找重复的子树
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	temp := make(map[string]*TreeNode)     //存序列
	repeat := make(map[*TreeNode]struct{}) //存重复的节点
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		str := fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right)) //格式化序列
		if n, ok := temp[str]; ok {
			//存在相同的序列，存入repeat
			//这里用n而不用node是因为n和node的结构是不同的，
			//n就表示n和node的输出序列是一致的，是重复的子树结构，
			//所以查重后用n表示node，即表示一致，如果在map里存node,
			//那么这是一个新key,不会重复的，因为他们的父节点不一样，这样最后得出的结果会重复
			repeat[n] = struct{}{}
		} else {
			temp[str] = node
		}
		return str
	}
	dfs(root)
	res := make([]*TreeNode, 0)
	for i := range repeat {
		res = append(res, i)
	}
	return res
}

//三元组解题
/* func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    type pair struct {
        node *TreeNode
        idx  int
    }
    repeat := map[*TreeNode]struct{}{}
    seen := map[[3]int]pair{}
    idx := 0
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        tri := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}
        if p, ok := seen[tri]; ok {
            repeat[p.node] = struct{}{}
            return p.idx
        }
        idx++
        seen[tri] = pair{node, idx}
        return idx
    }
    dfs(root)
    ans := make([]*TreeNode, 0, len(repeat))
    for node := range repeat {
        ans = append(ans, node)
    }
    return ans
} */
