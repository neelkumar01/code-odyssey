// reduce only takes array

const numbers = [1,2,3,4,5]

const sum1 = numbers.reduce((previousValue, currentValue) => {
    return previousValue + currentValue
})

/*
previousValue | currentValue | return
    1               2           3
    3               3           6
    6               4           10
    10              5           15
*/
console.log(sum1)



// example 2

const numberList = [10,20,30,40]

const sum2 = numberList.reduce((accumulator, currentValue) => {
    return accumulator + currentValue
}, 100)

/*
accumulator is previousValue only!

here we set initial accumulator value to 100
so,
accumulator  |  currentValue  |   return
    100             10              110
    110             20              130
    130             30              160
    160             40              200
*/
console.log(sum2)



// example 3

const userCart = [
    {productId: 1, productName: 'Laptop', price:80000},
    {productId: 2, productName: 'Mobile', price:40000},
    {productId: 3, productName: 'PC', price:100000}
]

const totalAmount = userCart.reduce((totalCost, currentCost) => {
    return totalCost + currentCost.price
}, 0)

/*
totalCost  |  currentCost  |  return
    0           80,000          80,000
    80,000      40,000          1,20,000
    1,20,000    1,00,000        2,20,000
*/

console.log(totalAmount)

