const buttons = document.querySelectorAll('.button-collection button')

console.log(buttons)

for (let button of buttons) {
    button.addEventListener('click', function() {
        console.log('i got clicked!')
        console.log(this.textContent)
    })
}