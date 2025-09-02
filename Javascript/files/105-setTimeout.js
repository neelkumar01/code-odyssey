console.log('start')

const delay = 0
const id = setTimeout(() => {
    console.log('inside setTimeout')
}, delay) 

let num = 0
for (let i = 0; i<=500; i++) {
    num += i
}
console.log(num)

console.log('end')

// use of id

clearTimeout(id)  
// this removes call back from queue