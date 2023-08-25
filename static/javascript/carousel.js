const previous = document.getElementById("previous");
const next = document.getElementById("next");
const carousel = document.getElementById("carousel");
const slider = document.getElementById("slider");

console.log(carousel.children)

// for (let i = 0; i < carousel.children.length; i++) {
//     console.log(carousel.children[i].clientWidth)
// }

let step = 0;
let mounToMove = carousel.firstElementChild.clientWidth + 32 //  window.innerWidth - 200
let width = carousel.offsetWidth  - (slider.offsetWidth);

const carouselEvent = ()=>{
    mounToMove = carousel.firstElementChild.clientWidth + 32 //  window.innerWidth - 200
    width = carousel.offsetWidth  - (slider.offsetWidth);
}

(function (){
    carouselEvent()
})()
    window.addEventListener("resize",()=>{
        carouselEvent()
})


previous.onclick = () => {
    if (step < 0) {
        step += mounToMove;
    }
    carousel.style.transform = "translate(" + `${step}` + "px,0%)";
};
next.onclick = () => {
    // let mounToMove = window.innerWidth - 200
    let limit = width - mounToMove;
    if (step > -limit) {
        step -= mounToMove;
    }
    carousel.style.transform = "translate(" + `${step}` + "px,0%)";
};