// difference b/w dot and bracket notation

const id = {
    name: 'neel',
    age:18,
    'user hobbies':['programming', 'open source contrbution']
}

// console.log(id.user hobbies)  ---> this will give error
console.log(id['user hobbies'])

const key = 'email'
id.key = 'abc@gmail.com'
id[key] = 'xyz@gmail.com'

console.log(id)