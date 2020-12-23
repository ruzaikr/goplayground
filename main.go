package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preOrderTraversal(a *TreeNode, seq string) string {
	if a == nil {
		return seq
	}

	seq = fmt.Sprintf("%s,-%d", seq, a.Val)

	if a.Left != nil {
		seq = preOrderTraversal(a.Left, seq)
	}else {
		seq = fmt.Sprintf("%s,%s", seq, "lnull")
	}

	if a.Right != nil {
		seq = preOrderTraversal(a.Right, seq)
	}else {
		seq = fmt.Sprintf("%s,%s", seq, "rnull")
	}

	return seq
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	var seqS = preOrderTraversal(s, "")
	fmt.Println("seqS:",seqS)
	var seqT = preOrderTraversal(t, "")
	fmt.Println("seqT:",seqT)
	return strings.Contains(seqS, seqT)
}

func main() {

}
