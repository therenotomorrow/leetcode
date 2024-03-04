const assert = require('assert');
const {createHelloWorld} = require('./createHelloWorld');

const testCases = [
    {name: 'smoke 1', args: [], want: 'Hello World'},
    {name: 'smoke 2', args: [[{}, null, 42]], want: 'Hello World'},
];

const tests = () => {
    assert.strictEqual(typeof createHelloWorld(), 'function')

    testCases.forEach(testCase => {
        const got = createHelloWorld()(...testCase.args);

        assert.strictEqual(got, testCase.want, `createHelloWorld = ${got}, want = ${testCase.want}`);
    });
}

tests();
