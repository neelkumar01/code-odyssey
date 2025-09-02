// html file : dom2.html

const heading = document.querySelector('.heading')

heading.classList.add('topText')

heading.classList.remove('heading')

const ans = heading.classList.contains('heading')

console.log(ans)

/*
toggle
- add class if not there
- remove class if there
*/

heading.classList.toggle('topText')

