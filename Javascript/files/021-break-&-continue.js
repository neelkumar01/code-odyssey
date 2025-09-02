for (let i = 0; i <= 10; i++) {
    if (i == 6) {
        console.log(`connection lost at t = ${i}s`)
        break
    }
    console.log(`connection going at t = ${i}s`)
}

console.log('')
for (let i = 0; i <= 10; i++) {
    if (i == 6) {
        continue
    }
    console.log(`connection going at t = ${i}s`)
}
