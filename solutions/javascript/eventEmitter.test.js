import { EventEmitter } from './eventEmitter'

test('smoke 1', () => {
  let got

  const emitter = new EventEmitter()

  got = emitter.emit('firstEvent')
  expect(got).toEqual([])

  emitter.subscribe('firstEvent', () => 5)
  emitter.subscribe('firstEvent', () => 6)

  got = emitter.emit('firstEvent')
  expect(got).toEqual([5, 6])
})

test('smoke 2', () => {
  let got

  const emitter = new EventEmitter()

  emitter.subscribe('firstEvent', (...args) => args.join(','))

  got = emitter.emit('firstEvent', [1, 2, 3])
  expect(got).toEqual(['1,2,3'])

  got = emitter.emit('firstEvent', [3, 4, 6])
  expect(got).toEqual(['3,4,6'])
})

test('smoke 3', () => {
  let got

  const emitter = new EventEmitter()
  const sub = emitter.subscribe('firstEvent', (...args) => args.join(','))

  got = emitter.emit('firstEvent', [1, 2, 3])
  expect(got).toEqual(['1,2,3'])

  sub.unsubscribe()

  got = emitter.emit('firstEvent', [4, 5, 6])
  expect(got).toEqual([])
})

test('smoke 4', () => {
  let got

  const emitter = new EventEmitter()

  const sub1 = emitter.subscribe('firstEvent', (x) => x + 1)

  emitter.subscribe('firstEvent', (x) => x + 2)
  sub1.unsubscribe()

  got = emitter.emit('firstEvent', [5])
  expect(got).toEqual([7])
})
