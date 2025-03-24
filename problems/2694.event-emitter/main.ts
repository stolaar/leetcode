type Callback = (...args: any[]) => any;
type Subscription = {
  unsubscribe: () => void
}

class ExtendedMap<K, V> extends Map<K, V> {
  getItem(key: any, defaultValue: V) {
    return super.get(key) ?? defaultValue
  }
}

class EventEmitter {
  private subscriptionsMap: ExtendedMap<string, Callback[]>

  constructor() {
    this.subscriptionsMap = new ExtendedMap()
  }

  subscribe(eventName: string, callback: Callback): Subscription {
    const callbacks = this.subscriptionsMap.getItem(eventName, [])
    callbacks.push(callback)

    this.subscriptionsMap.set(eventName, callbacks)

    return {
      unsubscribe: () => {
        const updatedCallbacks = this.subscriptionsMap.getItem(eventName, []).filter((cb) => cb !== callback)
        this.subscriptionsMap.set(eventName, updatedCallbacks)
      }
    };
  }

  emit(eventName: string, args: any[] = []): any[] {
    return this.subscriptionsMap.getItem(eventName, []).map((cb) => cb(...args))
  }
}

/**
 * const emitter = new EventEmitter();
 *
 * // Subscribe to the onClick event with onClickCallback
 * function onClickCallback() { return 99 }
 * const sub = emitter.subscribe('onClick', onClickCallback);
 *
 * emitter.emit('onClick'); // [99]
 * sub.unsubscribe(); // undefined
 * emitter.emit('onClick'); // []
 */
