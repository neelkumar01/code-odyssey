const person1 = {
    Name: 'Neel Kumar',
    id: 12
}
// method1 

let printDetails = (obj) => {
    console.log(obj.Name)
    console.log(obj.id)
}

printDetails(person1)



// method 2 : object parameters destructuring

const person2 = {
    Name: 'Krishna',
    id: 7,
    location: 'global'
}

let getDetails = ( {Name, id, location}) => {
    console.log(Name)
    console.log(id)
    console.log(location)
}

getDetails(person2)
