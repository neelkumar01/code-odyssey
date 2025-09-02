// first local variable is checked when called
// if not found, it go in lexical environment

let connection = () => {
    let port = 'connected'

    let func = () => {
        let port = 'not connected'
        console.log(port)
    }

    console.log(port)
    func()
 }

 connection()

 