// const previous = document.getElementById("previous");
// const next = document.getElementById("next");
// const carousel = document.querySelectorAll(".carousel");

// let step = 0;
// let width = carousel.offsetWidth;
// let mounToMove = window.innerWidth - 200;

// const carouselEvent = () => {
//     let isMobile = window.innerWidth <= 768;
//     mounToMove = isMobile
//         ? carousel.firstElementChild.clientWidth + 64
//         : window.innerWidth - 200;
// };
// (function () {
//     carouselEvent();
// })();

// window.addEventListener("resize", () => {
//     carouselEvent();
// });

// previous.onclick = () => {
//     if (step < 0) {
//         step += mounToMove;
//     }
//     carousel.style.transform = "translate(" + `${step}` + "px,-50%)";
// };
// next.onclick = () => {
//     // let mounToMove = window.innerWidth - 200
//     let limit = width - mounToMove;
//     if (step > -limit) {
//         step -= mounToMove;
//     }
//     carousel.style.transform = "translate(" + `${step}` + "px,-50%)";
// };

// document.querySelectorAll(".carousel").forEach((current) => {
//     let one =
//         ((current.firstElementChild.clientWidth + 64) * 100) /
//         current.clientWidth;
//     console.log(one * 52);

//     previous.onclick = () => {
//         current.style.transform = `translateX(${
//             current.getBoundingClientRect().x -
//             (current.firstElementChild.clientWidth + 64)
//         }px)`;
//         // console.log(current);
//     };
// });
