function createCounter(n: number): () => number {
  let counter = n - 1

  return function() {
    counter++
    return counter
  }
}


const counter = createCounter(10)

console.log(counter(), counter(), counter())
