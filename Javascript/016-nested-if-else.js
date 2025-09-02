// number guess game

let winningNum = 19

let userGuess = 7

if (userGuess === winningNum) {
    console.log('your guess is correct')
} else {
    if (userGuess > winningNum) {
        console.log('too high!')
    } else {
        console.log('too low!')
    }
}
