// when any arg not passed while calling function

let sum = (a,b = 4) => {
    return a+b
}

console.log(sum(10, 20))

console.log(sum(10))

