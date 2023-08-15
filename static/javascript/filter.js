var slideOne;
let ranger = document.querySelectorAll(".range-tracker");
ranger.forEach((current) => {
    let sliderOne = current.querySelectorAll(".slider-0")[0];
    let sliderTwo = current.querySelectorAll(".slider-1")[0];
    let displayValOne = current.querySelectorAll(".slider-1-number")[0];
    let displayValTwo = current.querySelectorAll(".slider-2-number")[0];

    let minGap = 1;
    console.log(minGap);
    let sliderTrack = current.querySelectorAll(".slider-track")[0];
    let sliderMaxValue = current.querySelectorAll(".slider-0")[0].max;
    let sliderMinValue = current.querySelectorAll(".slider-1")[0].min;

    slideOne();
    slideTwo();

    function slideOne() {
        if (parseInt(sliderTwo.value) - parseInt(sliderOne.value) <= minGap) {
            sliderOne.value = parseInt(sliderTwo.value) - minGap;
        }
        displayValOne.textContent = sliderOne.value;
        fillColor();
    }
    function slideTwo() {
        if (parseInt(sliderTwo.value) - parseInt(sliderOne.value) <= minGap) {
            sliderTwo.value = parseInt(sliderOne.value) + minGap;
        }
        displayValTwo.textContent = sliderTwo.value;
        fillColor();
    }

    function fillColor() {
        percent1 =
            ((sliderOne.value - sliderMinValue) /
                (sliderMaxValue - sliderMinValue)) *
            100;
        percent2 =
            ((sliderTwo.value - sliderMinValue) /
                (sliderMaxValue - sliderMinValue)) *
            100;
            sliderTrack.style.background = `#dadae5`;
            sliderTrack.style.background = `linear-gradient(to right, #dadae5 ${percent1}% , #22c55e ${percent1}% , #22c55e ${percent2}%, #dadae5 ${percent2}%)`;

    }

    sliderOne.addEventListener("input", slideOne);
    sliderTwo.addEventListener("input", slideTwo);
});

let isOpen = false;
document.querySelector(".show-filter").addEventListener("click", function (e) {
    e.target.textContent = isOpen ? "Hide Fields" : "Show Fields";
    isOpen = !isOpen;
    document.querySelector(".fields").style.display = isOpen
        ? "none"
        : "inline";
});
