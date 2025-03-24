type F = (...args: unknown[]) => void

function debounce(fn: F, t: number): F {
  let timeoutId: any

  return function(...args) {
    if (timeoutId) {
      clearTimeout(timeoutId)
    }

    timeoutId = setTimeout(() => fn(...args), t)
  }
};


const log = debounce(console.log, 100);
log('Hello'); // cancelled
log('Hello'); // cancelled
log('Hello here'); // Logged at t=100ms
