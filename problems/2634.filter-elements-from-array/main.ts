type Fn = (n: number, i: number) => any

function filter(arr: number[], fn: Fn): number[] {
  const result: number[] = []

  arr.forEach((num, i) => {
    if (fn(num, i)) {
      result.push(num)
    }
  })

  return result
};
