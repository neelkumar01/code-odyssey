const para1 = document.querySelector("#para1");
const para2 = document.querySelector("#para2");
const para3 = document.querySelector("#para3");
const para4 = document.querySelector("#para4");
const para5 = document.querySelector("#para5");
const para6 = document.querySelector("#para6");
const para7 = document.querySelector("#para7");

function changeText(element, text, time) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            if (element) {
                element.textContent = text
                resolve()
            } else {
                reject('some error')
            }
        }, time)
    })
}

changeText(para1, 'my para 1', 1000)
    .then(() => { changeText(para2, 'my para 2', 1000) })
    .then(() => { changeText(para3, 'my para 3', 2000) })
    .then(() => { changeText(para4, 'my para 4', 3000) })
    .then(() => { changeText(para5, 'my para 5', 4000) })
    .then(() => { changeText(para6, 'my para 6', 5000) })
    .then(() => { changeText(para7, 'my para 7', 6000) })
    .catch((error) => {console.log(error)})

    