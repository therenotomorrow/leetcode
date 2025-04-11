import { jsonStringify } from './jsonStringify'

describe('jsonStringify', () => {
  test.each([
    [
      'smoke 1',
      {
        y: 1,
        x: 2,
      },
      `{"y":1,"x":2}`,
    ],
    [
      'smoke 2',
      {
        a: 'str',
        b: -12,
        c: true,
        d: null,
      },
      `{"a":"str","b":-12,"c":true,"d":null}`,
    ],
    [
      'smoke 3',
      {
        key: {
          a: 1,
          b: [{}, null, 'Hello'],
        },
      },
      `{"key":{"a":1,"b":[{},null,"Hello"]}}`,
    ],
    ['smoke 4', true, `true`],
  ])('%p', (name, obj, want) => {
    const got = jsonStringify(obj)

    expect(got).toEqual(want)
  })
})
