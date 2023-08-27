const previous = document.getElementById("previous");
const next = document.getElementById("next");
const carousel = document.getElementById("carousel");
const slider = document.getElementById("slider");

let step = 0
let width = carousel.getBoundingClientRect().width  - (slider.getBoundingClientRect().width);

next.onclick = () => {
    let currentelem = carousel.querySelector(".active");
    let nextElement = currentelem.nextElementSibling;

    console.log(currentelem.getBoundingClientRect().width)
    if ((width > -step) && (nextElement != null )) {
        step -= currentelem.getBoundingClientRect().width
        carousel.style.transform = "translate(" + `${step}` + "px,0%)";
        currentelem.classList.remove("active");
        nextElement.classList.add("active");
    } 
}

previous.onclick = () => {
    let currentelem = carousel.querySelector(".active");
    let previouselem = currentelem.previousElementSibling;

    if ((step < 0) && (previouselem != null )) {
        step += currentelem.getBoundingClientRect().width
        carousel.style.transform = "translate(" + `${step}` + "px,0%)";
        currentelem.classList.remove("active");
        previouselem.classList.add("active");
    }
}
