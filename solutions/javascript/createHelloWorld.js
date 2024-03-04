const createHelloWorld = function () {
    return function (..._) {
        return 'Hello World';
    }
};

module.exports = {createHelloWorld};
