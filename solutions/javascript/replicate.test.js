import { replicate } from './replicate'

describe('replicate', () => {
  test.each([
    ['smoke 1', { str: 'hello', times: 2 }, 'hellohello'],
    ['smoke 2', { str: 'code', times: 3 }, 'codecodecode'],
    ['smoke 3', { str: 'js', times: 1 }, 'js'],
  ])('%p', (name, { str, times }, want) => {
    expect(replicate(str, times)).toEqual(want)
  })
})
