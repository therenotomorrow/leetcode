import { memoize } from './memoize'

test('smoke 1', () => {
  let calls = 0

  const sum = (a, b) => {
    calls++

    return a + b
  }
  const memoized = memoize(sum)

  expect(memoized(2, 2)).toEqual(4)
  expect(memoized(2, 2)).toEqual(4)
  expect(memoized(1, 2)).toEqual(3)

  expect(calls).toEqual(2)
})

test('smoke 2', () => {
  let calls = 0

  const factorial = (n) => {
    calls++

    return n <= 1 ? 1 : n * factorial(n - 1)
  }
  const memoized = memoize(factorial)

  expect(memoized(2)).toEqual(2)
  expect(memoized(3)).toEqual(6)
  expect(memoized(2)).toEqual(2)

  expect(calls).toEqual(5) // 5 because of recursive

  expect(memoized(3)).toEqual(6)

  expect(calls).toEqual(5) // 5 because of recursive
})

test('smoke 3', () => {
  let calls = 0

  const fib = function (n) {
    calls++

    return n <= 1 ? 1 : fib(n - 1) + fib(n - 2)
  }
  const memoized = memoize(fib)

  expect(memoized(5)).toEqual(8)

  expect(calls).toEqual(15) // 15 because of recursive

  expect(memoized(5)).toEqual(8)

  expect(calls).toEqual(15) // 15 because of recursive
})
