import { curry } from './curry'

test('smoke 1', async () => {
  const fn = (a, b, c) => a + b + c
  const curried = curry(fn)

  const got = curried(1)(2)(3)
  const want = fn(1, 2, 3)

  expect(got).toEqual(want)
})

test('smoke 2', async () => {
  const fn = (a, b, c) => a + b + c
  const curried = curry(fn)

  const got = curried(1, 2)(3)
  const want = fn(1, 2, 3)

  expect(got).toEqual(want)
})

test('smoke 3', async () => {
  const fn = (a, b, c) => a + b + c
  const curried = curry(fn)

  const got = curried()()(1, 2, 3)
  const want = fn(1, 2, 3)

  expect(got).toEqual(want)
})

test('smoke 4', async () => {
  const life = () => 42
  const curriedLife = curry(life)

  expect(curriedLife()).toEqual(42)
})
