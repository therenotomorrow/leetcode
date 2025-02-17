import { isEmpty } from './isEmpty'

describe('isEmpty', () => {
  test.each([
    ['smoke 1', { obj: { x: 5, y: 42 } }, false],
    ['smoke 2', { obj: {} }, true],
    ['smoke 3', { obj: [null, false, 0] }, false],
  ])('%p', (name, { obj }, want) => {
    expect(isEmpty(obj)).toEqual(want)
  })
})
