document.addEventListener("DOMContentLoaded", () => {
    const filterInput = document.querySelector("#search");
    const suggestionSection = document.querySelector("#suggestion");

    filterInput.addEventListener("input", () => {
        const filterText = filterInput.value.toLowerCase();
        const divs = suggestionSection.querySelectorAll("a");
        let hasMatch = false;

        if (filterText === "") {
            suggestionSection.classList.add("hidden");
        } else {
            divs.forEach((div) => {
                const pTag = div.querySelector("p");
                if (pTag && pTag.textContent.toLowerCase().includes(filterText)) {
                    div.classList.remove("hidden");
                    hasMatch = true;
                } else {
                    div.classList.add("hidden");
                }
            });

            if (hasMatch) {
                suggestionSection.classList.remove("hidden");
            } else {
                suggestionSection.classList.add("hidden");
            }
        }
    });
});