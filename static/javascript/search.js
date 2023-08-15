const filterInput = document.querySelector("#search");
const suggestionSection = document.querySelector(".suggestions");
const divs = document.querySelectorAll(".suggestions div");
filterInput.addEventListener("input", filterDivs);

function filterDivs() {
    const filterText = filterInput.value.toLowerCase();
    let hasMatch = false;

    if (filterText === "") {
        suggestionSection.style.display = "none";
    } else {
        divs.forEach((div) => {
            const pTagText = div.querySelector("p").textContent.toLowerCase();
            if (pTagText.includes(filterText)) {
                div.style.display = "block";
                hasMatch = true;
            } else {
                div.style.display = "none";
            }
        });

        suggestionSection.style.display = hasMatch ? "block" : "none";
    }
}
// filterInput.addEventListener("blur", function() {
//     // Hide the suggestion section when the input loses focus
//     suggestionSection.style.display = "none";
// });
filterDivs();
