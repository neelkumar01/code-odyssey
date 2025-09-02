// callback recap

function myfunc(callback) {
  console.log("hello");
  callback();
}

myfunc((a, b) => {
  a = 5;
  b = 4;
  console.log("a + b =", a + b);
});

// callback hell, pyramid of doom

const p1 = document.querySelector("#p1");
const p2 = document.querySelector("#p2");
const p3 = document.querySelector("#p3");

const id = setTimeout(() => {
  p1.textContent = "web content 1";
  setTimeout(() => {
    p2.textContent = "web content 2";
    setTimeout(() => {
      p3.textContent = "web content 3";
    }, 1000);
  }, 1000);
}, 1000);
/*
- above code is callback hell
- callback hell : when there is callback inside callback inside callback...
*/

// doing same thing above with functions

clearTimeout(id); // stoping our code to execute further one!

const para1 = document.querySelector("#para1");
const para2 = document.querySelector("#para2");
const para3 = document.querySelector("#para3");
const para4 = document.querySelector("#para4");
const para5 = document.querySelector("#para5");
const para6 = document.querySelector("#para6");
const para7 = document.querySelector("#para7");

function changeText(element, text, time, successCallback, failureCallback) {
  setTimeout(()=>{
    if (element) {
      element.textContent = text
      if (successCallback) {
        successCallback()
      }
    } else {
      if (failureCallback) {
        failureCallback()
      }
    }

  }, time)
}

changeText(para1, 'my para 1', 1000, ()=>{
  changeText(para2, 'my para 2', 1000, ()=>{
    changeText(para3, 'my para 3', 1000, ()=>{
      changeText(para4, 'my para 4', 1000, ()=>{
        changeText(para5, 'my para 5', 1000, ()=>{
          changeText(para6, 'my para 6', 1000, ()=>{
            changeText(para7, 'my para 7', 1000, ()=>{

            }, ()=>{console.log('para7 dont exist')})
          }, ()=>{console.log('para6 dont exits')})
        }, ()=>{console.log('para5 dont exist')})
      }, ()=>{console.log('para4 dont exist')})
    }, ()=>{console.log('para3 dont exist')})
  }, ()=>{console.log('para2 dont exits')})
}, ()=>{console.log('para1 dont exist')})

// above structure is 'pyramid of doom'