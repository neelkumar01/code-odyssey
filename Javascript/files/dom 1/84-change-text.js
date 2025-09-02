const myHeading = document.getElementById('my-heading')

// changing text
console.log(myHeading.textContent)

myHeading.textContent = 'Making magic happen...'


// textContent and innerText

// textContent : also showing element set to display none
// innerText : don't show element set to display none

const aboutMe = document.querySelector('.about-me')

console.log('text content :', aboutMe.textContent)

console.log('inner text :', aboutMe.innerText)

