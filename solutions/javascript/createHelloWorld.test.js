const createHelloWorld = require('./createHelloWorld')
const expect = require('expect')

describe('createHelloWorld', () => {
    test.each([
        ['smoke 1', [], 'Hello World'],
        ['smoke 2', [{}, null, 42], 'Hello World']
    ])('%p', (name, args, want) => {
        expect.expect(createHelloWorld()(...args)).toEqual(want)
    })
})
