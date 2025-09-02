// rest parameters  -  to make extra parameters given to get stored in array

function myFunction(x,y, ...z) {
    console.log(x)
    console.log(y)
    console.log(z)
}

myFunction(12, 24, 1,2,3,4,5)


let addAll = (...numbers) => {
    let total = 0
    for (number of numbers) {
        total += number
    }
    return total
}

let sum = addAll(10,20,30,40)
console.log(sum)