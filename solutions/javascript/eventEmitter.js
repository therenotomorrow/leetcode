export class EventEmitter {
  constructor() {
    this.events = {}
  }

  subscribe(eventName, callback) {
    this.events[eventName] = this.events[eventName] ?? []
    this.events[eventName].push(callback)

    return {
      unsubscribe: () => {
        this.events[eventName] = this.events[eventName].filter((f) => f !== callback)

        if (this.events[eventName].length === 0) {
          delete this.events[eventName]
        }
      },
    }
  }

  emit(eventName, args = []) {
    if (!(eventName in this.events)) {
      return []
    }
    return this.events[eventName].map((f) => f(...args))
  }
}
