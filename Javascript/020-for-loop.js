for (let i = 0; i <= 10; i++) {
    console.log(i)
}
// here i can't be accessed outside with let


// with var, n can be accessed outside
for (var n = 0; n <= 5; n++) {
    console.log(n)
}
console.log(`current value of n = ${n}`)
