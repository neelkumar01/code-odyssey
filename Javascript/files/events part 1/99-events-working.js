const buttons = document.querySelectorAll('.functional-buttons button')

console.log(buttons)

// buttons.forEach((button) => {
//     button.addEventListener('click', (event)=>{
//         console.log(event.currentTarget)
//         console.log(event.currentTarget.textContent)
//     })
// })


buttons.forEach((button) => {
    button.addEventListener('click', (event) => {

        // following for loop act as time delay

        let num = 0
        for (let i = 0; i <= 1000000000; i++) {
            num += i
        }

        console.log(event.currentTarget)
    })
})

let Var = 0
for (let i = 0; i <= 1000000000; i++) {
    Var += i
}
console.log('code ends...')