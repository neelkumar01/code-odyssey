// event bubbling / propagation
// event capturing
// event delegation

const grandparent = document.querySelector(".grandparent");
const parent = document.querySelector(".parent");
const child = document.querySelector(".child");
const body = document.body;

child.addEventListener("click", () => {
  console.log("bubble - child");
});

parent.addEventListener("click", () => {
  console.log("bubble - parent");
});

grandparent.addEventListener("click", () => {
  console.log("bubble - grandparent");
});

body.addEventListener("click", () => {
  console.log("bubble - body");
});

/*
when clicked on child
- event called for child and also for parent, grandparent, body
- this is event bubbling / propagation
*/




// event capturing

child.addEventListener("click", () => {
  console.log("capture - child");
}, true);

parent.addEventListener("click", () => {
  console.log("capture - parent");
}, true);

grandparent.addEventListener("click", () => {
  console.log("capture - grandparent");
}, true);

body.addEventListener("click", () => {
  console.log("capture - body");
}, true);



