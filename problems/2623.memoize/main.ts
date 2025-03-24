type Fn = (...params: number[]) => number

function memoize(fn: Fn): Fn {
  const map = new Map<Fn, Map<string, number>>()

  const hashInput = (...args: unknown[]): string => args.join("-")

  return function(...args) {
    const hash = hashInput(...args)

    if (!map.get(fn)?.has(hashInput(...args))) {
      const result = fn(...args)
      if (!map.has(fn)) {
        const fnResultsMap = new Map([[hash, result]])
        map.set(fn, fnResultsMap)
        return result
      }

      map.get(fn)?.set(hash, result)
      return result
    }

    return map.get(fn)?.get(hash) as number
  }
}


let callCount = 0;
const memoizedFn = memoize(function(a, b) {
  callCount += 1;
  return a + b;
})
memoizedFn(2, 3) // 5
memoizedFn(2, 3) // 5
console.log(callCount) // 1
