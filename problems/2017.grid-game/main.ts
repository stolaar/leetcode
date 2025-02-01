function gridGame(grid: number[][]): number {
  const sums = grid[0].reduce((acc, curr, idx) => {
    if (idx === 0) {
      acc.right += curr
      return acc
    }

    acc.down += grid[1][idx]
    return acc
  }, { right: 0, down: 0 })

  const firstRobotMaxPath = grid[0].reduce((acc, curr, idx) => {

    return acc
  }, { currDown: sums.down, currRight: sums.right, maxFirst: -Infinity })


  return 4
};

console.log(gridGame([[]]))
