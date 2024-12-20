class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.left = (left === undefined ? null : left)
    this.right = (right === undefined ? null : right)
  }
}

const bfs = (left: TreeNode | null, right: TreeNode | null, level: number) => {
  if (!left || !right) {
    return
  }

  if (level % 2 !== 0) {
    [left.val, right.val] = [right.val, left.val]
  }

  bfs(left.left, right.right, level + 1)
  bfs(left.right, right.left, level + 1)
}

function reverseOddLevels(root: TreeNode | null): TreeNode | null {
  if (!root?.left) {
    return root
  }

  bfs(root.left, root.right, 1)

  return root
};
