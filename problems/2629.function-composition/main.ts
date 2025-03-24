type Func = (x: number) => number;

function compose(functions: Func[]): Func {

  return function(x) {
    if (functions.length === 0) {
      return x
    }

    return functions.reduceRight((acc, curr) => {
      return curr(acc)
    }, x)
  }
};


const fn = compose([x => x + 1, x => x * x, x => 2 * x])
console.log(fn(4))
