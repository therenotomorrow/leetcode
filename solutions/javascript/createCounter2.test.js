const createCounter2 = require('./createCounter2')

test('smoke 1', () => {
  const counter = createCounter2(5)

  expect(counter.increment()).toEqual(6)
  expect(counter.reset()).toEqual(5)
  expect(counter.decrement()).toEqual(4)
})

test('smoke 2', () => {
  const counter = createCounter2(0)

  expect(counter.increment()).toEqual(1)
  expect(counter.increment()).toEqual(2)
  expect(counter.decrement()).toEqual(1)
  expect(counter.reset()).toEqual(0)
  expect(counter.reset()).toEqual(0)
})
