// html file : dom3.html

const fruits = document.querySelector('.fruits')

const element = document.createElement('li')

const text = document.createTextNode('Blue berry')

element.append(text)
fruits.append(element)

// prepend : like append but add to start


// before()
const portfolio = document.createElement('a')

const link = document.createTextNode('Personal Portfolio')

portfolio.append(link)
fruits.before(portfolio)

fruits.style.paddingTop = '100px'

// after() as name say same like before() but just appends after!

