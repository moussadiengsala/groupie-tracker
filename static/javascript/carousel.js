
const previous = document.getElementById("previous");
const next = document.getElementById("next");
const carousel = document.getElementById("carousel");
const slider = document.getElementById("slider");

let step = 0;
let cardWidth = carousel.querySelector('.active')?.getBoundingClientRect().width + 24; // 24px gap
let visibleCards = Math.floor(slider.getBoundingClientRect().width / cardWidth);
let maxStep = (carousel.children.length - visibleCards) * cardWidth;

// Initialize button states
updateButtonStates();

next.addEventListener('click', () => {
    let currentCard = carousel.querySelector(".active");
    let nextCard = currentCard.nextElementSibling;
    
    if (nextCard && step > -maxStep) {
        step -= cardWidth;
        carousel.style.transform = `translateX(${step}px)`;
        currentCard.classList.remove("active");
        nextCard.classList.add("active");
        updateButtonStates();
    }
});

previous.addEventListener('click', () => {
    let currentCard = carousel.querySelector(".active");
    let prevCard = currentCard.previousElementSibling;
    
    if (prevCard && step < 0) {
        step += cardWidth;
        carousel.style.transform = `translateX(${step}px)`;
        currentCard.classList.remove("active");
        prevCard.classList.add("active");
        updateButtonStates();
    }
});

function updateButtonStates() {
    previous.disabled = step >= 0;
    next.disabled = step <= -maxStep;
}

// Handle window resize
window.addEventListener('resize', () => {
    cardWidth = carousel.querySelector('.active')?.getBoundingClientRect().width + 24;
    visibleCards = Math.floor(slider.getBoundingClientRect().width / cardWidth);
    maxStep = (carousel.children.length - visibleCards) * cardWidth;
    updateButtonStates();
});





// const previous = document.getElementById("previous");
// const next = document.getElementById("next");
// const carousel = document.getElementById("carousel");
// const slider = document.getElementById("slider");

// let step = 0
// let width = carousel.getBoundingClientRect().width  - (slider.getBoundingClientRect().width);

// next.onclick = () => {
//     let currentelem = carousel.querySelector(".active");
//     let nextElement = currentelem.nextElementSibling;

//     console.log(currentelem.getBoundingClientRect().width)
//     if ((width > -step) && (nextElement != null )) {
//         step -= currentelem.getBoundingClientRect().width
//         carousel.style.transform = "translate(" + `${step}` + "px,0%)";
//         currentelem.classList.remove("active");
//         nextElement.classList.add("active");
//     } 
// }

// previous.onclick = () => {
//     let currentelem = carousel.querySelector(".active");
//     let previouselem = currentelem.previousElementSibling;

//     if ((step < 0) && (previouselem != null )) {
//         step += currentelem.getBoundingClientRect().width
//         carousel.style.transform = "translate(" + `${step}` + "px,0%)";
//         currentelem.classList.remove("active");
//         previouselem.classList.add("active");
//     }
// }
