const numbers = [1,3,5,7,10]

// if any of them is even it returns true
// example,

let ans = numbers.some((number) => {
    return number % 2 === 0
})

console.log(ans)