// splice method
// start, delete(items?), insert(items?)

const folder = ['doc1', 'doc2', 'doc3', 'doc4', 'doc5', 'doc6']

// deleting doc2, doc3, doc4
folder.splice(1, 3)

console.log(folder)

// inserting doc after doc5
folder.splice(2, 0, 'my doc')

console.log(folder)