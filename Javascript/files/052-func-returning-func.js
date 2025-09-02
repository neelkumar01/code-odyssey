let func1 = () => {
    let func2 = () => {
        console.log('Hello from func2')
    }
    return func2
}

let result = func1()
result()

// result func calling func1 which returns func2

