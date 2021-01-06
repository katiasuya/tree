# tree

Here is the package tree that allows you to check whether two binary trees store the same sequence.

1. Walk function walks the tree t sending all values from the tree to the channel ch.

    func Walk(t *tree.Tree, ch chan int)

2. Using Walk function, Same walks the trees to determine whether they store the same values or not.

    func Same(t1, t2 *tree.Tree) bool 