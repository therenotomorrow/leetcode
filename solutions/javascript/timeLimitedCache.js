export class TimeLimitedCache {
  constructor() {
    this.cache = new Map()
  }

  set(key, value, duration) {
    const cached = this.cache.get(key)

    if (cached !== undefined) {
      clearTimeout(cached.timeId)
    }

    const timeId = setTimeout(() => this.cache.delete(key), duration)

    this.cache.set(key, {
      value,
      timeId,
    })

    return Boolean(cached)
  }

  get(key) {
    const cached = this.cache.get(key)

    if (cached !== undefined) {
      return cached.value
    }

    return -1
  }

  count() {
    return this.cache.size
  }
}
