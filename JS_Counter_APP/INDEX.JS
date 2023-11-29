console.log("hello")


let count = 0;
const countHeader = document.getElementById("count")
const countAddbtn = document.getElementById("ADD")
const countSubtractbtn = document.getElementById("Subtract")

countHeader.innerText = count

countAddbtn.addEventListener("click" , ()=>{
    count++;
    countHeader.innerText = count
})

countSubtractbtn.addEventListener("click" , ()=>{
    count--;
    countHeader.innerText = count
})