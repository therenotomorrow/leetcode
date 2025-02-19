import { Calculator, divisionByZero } from './calculator'

test('smoke 1', () => {
  const calc = new Calculator(10)

  const got = calc.add(5).subtract(7).getResult()

  expect(got).toEqual(8)
})

test('smoke 2', () => {
  const calc = new Calculator(2)

  const got = calc.multiply(5).power(2).getResult()

  expect(got).toEqual(100)
})

test('smoke 3', () => {
  const calc = new Calculator(20)

  try {
    calc.divide(0).getResult()
  } catch (err) {
    expect(err).toEqual(divisionByZero)
  }
})
