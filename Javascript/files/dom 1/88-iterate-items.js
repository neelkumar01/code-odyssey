const navItems = document.getElementsByClassName('item')

for (navItem of navItems) {
    console.log(navItem)
    navItem.style.color = '#FF549B'
}

const lines = document.getElementsByTagName('p')

for (line of lines) {
    line.style.color = '#007DFF'
}

