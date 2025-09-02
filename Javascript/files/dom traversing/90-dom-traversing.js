// root node of webpage

const root_node = document.getRootNode()

console.log(root_node)

const html_node = root_node.childNodes[0]

console.log(html_node)

const html_childs = html_node.childNodes   // [head, text, body]

console.log(html_childs)  

const head = html_childs[0]
const text = html_childs[1]
const body = html_childs[2]

console.log(head)
console.log(text)
console.log(body)

// [head, text, body]  -  siblings
console.log(head.nextSibling)
console.log(head.nextElementSibling)

console.log(head.childNodes)


// reaching div through h1
const h1 = document.querySelector('h1')
const div = h1.parentNode

div.style.color = 'pink'

const web_body = div.parentNode
web_body.style.backgroundColor = 'black'


