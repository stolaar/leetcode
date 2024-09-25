class TNode {
  children: TNode[]
  constructor() {
    this.children = new Array(10)
  }
}

const insert = (tree: TNode, key: number) => {
  let current = tree
  const keystr = `${key}`.split('')

  for (const c of keystr) {

    if (!current.children[c.charCodeAt(0) - 48]) {
      current.children[c.charCodeAt(0) - 48] = new TNode()
    }

    current = current.children[c.charCodeAt(0) - 48]
  }

}

const search = (tree: TNode, key: number): number => {
  let current = tree
  const keystr = `${key}`.split('')
  let count = 0

  for (const c of keystr) {
    if (!current.children[c.charCodeAt(0) - 48]) {
      return count
    }
    count += 1

    current = current.children[c.charCodeAt(0) - 48]
  }
  return count
}

const longestCommonPrefix = (arr1: number[], arr2: number[]): number => {
  const tree = new TNode()

  if (arr1.length > arr2.length) {
    [arr1, arr2] = [arr2, arr1]
  }

  arr1.forEach((num) => {
    insert(tree, num)
  })

  return arr2.reduce((count, num) => {
    const depth = search(tree, num)
    if (depth > count) {
      return depth
    }
    return count
  }, 0)
}

console.log(longestCommonPrefix([1, 10, 100], [1000]))
console.log(longestCommonPrefix([1, 2, 3], [4, 4, 4]))
