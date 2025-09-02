const grandparent = document.querySelector('.grandparent')

grandparent.addEventListener('click', (e)=>{
    console.log('you click grand pa!')
    console.log(e.target)
    console.log(e.target.textContent)
})

/*
- we click on child, still we get 'you click grand pa'
- it goes through cycle
- dont see capturing to do nothin
- then as we click on child, dont found any call back
- going through cycle, it found and return call back func of grandparent
*/