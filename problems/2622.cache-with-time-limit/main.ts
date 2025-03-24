class TimeLimitedCache {
  #map: Map<number, { value: number, duration: number }>

  constructor() {
    this.#map = new Map()
  }

  set(key: number, value: number, duration: number): boolean {
    const current = this.#map.get(key)
    const now = Date.now()

    this.#map.set(key, { value, duration: now + duration })
    if (current && current.duration > now) {
      return true
    }

    return false
  }

  get(key: number): number {
    const current = this.#map.get(key)

    if (!current) {
      return -1
    }

    if (current.duration < Date.now()) {
      this.#map.delete(key)
      return -1
    }

    return current.value
  }

  count(): number {
    let counter = 0
    const now = Date.now()

    this.#map.forEach((value, key) => {
      if (value.duration > now) {
        counter++
        return
      }

      this.#map.delete(key)
    })

    return counter
  }
}

/**
 * const timeLimitedCache = new TimeLimitedCache()
 * timeLimitedCache.set(1, 42, 1000); // false
 * timeLimitedCache.get(1) // 42
 * timeLimitedCache.count() // 1
 */
