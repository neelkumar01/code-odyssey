// set : iterable!
// nod index based access
// no duplicate items

// making a set
const mySet = new Set([1,2,3])

console.log(mySet)

const set1 = new Set('abc')
console.log(set1)

const set2 = new Set()
set2.add(1)
set2.add(12)
set2.add(24)
console.log(set2)

// checking element in set
const set3 = new Set([10,20,30,40])
console.log(set3.has(30))


// for of loop on set
const folder = new Set()
folder.add('doc1')
folder.add('doc2')
folder.add('doc3')
console.log(folder)

for (item of folder) {
    console.log(item)
}