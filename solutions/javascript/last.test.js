import { last } from './last'

describe('last', () => {
  test.each([
    ['smoke 1', [null, {}, 3], 3],
    ['smoke 2', [], -1],
    ['test 25: wrong answer', [null], null],
  ])('%p', (name, nums, want) => {
    nums.last = last()

    expect(nums.last()).toEqual(want)
  })
})
