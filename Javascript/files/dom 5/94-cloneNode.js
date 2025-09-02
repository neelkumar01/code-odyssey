// html file : dom5.html


const container = document.querySelector('.container')

const p1 = document.createElement('p')
const content = document.createTextNode('Hi, I am neel :)')
p1.append(content)

const p2 = p1.cloneNode(true)  // true means to clone childs also

container.append(p1)
container.prepend(p2)

p1.style.backgroundColor = 'pink'  // appended content with pink

p2.style.backgroundColor = 'gold'


// querySelectorAll  -  give static list

// getElementBy...  -  give live list

const bag = document.querySelector('.bag')

const bag_items = document.querySelectorAll('.bag li')

const items = bag.getElementsByTagName('li')

const item5 = document.createElement('li')
const text = document.createTextNode('item 5')
item5.append(text)
bag.append(item5)

console.log(bag_items)
console.log(items)