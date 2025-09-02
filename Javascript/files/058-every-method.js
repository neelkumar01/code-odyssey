const numbers = [2,4,6,8,10]

// every method --> returns true/false

const ans = numbers.every((number) => {
    return (number % 2 === 0)
})

console.log(ans)

