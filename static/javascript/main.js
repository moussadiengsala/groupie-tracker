document.querySelectorAll(".dates").forEach((date) => {
    const splittedDate = date.textContent.split("-");
    const dat = new Date(splittedDate[2], splittedDate[1] - 1, splittedDate[0]);
    const formattedDate = dat.toLocaleDateString("en-US", {
        day: "numeric",
        month: "short",
        year: "numeric",
    });
    date.innerHTML = formattedDate;
});
let darkMode = localStorage.getItem("dark-mode");
document.querySelectorAll(".dates-and-locations").forEach((dateAndLocation) => {
    splittedDateAndLocation = dateAndLocation.textContent.split("-");
    dateAndLocation.innerHTML = `<p class="">${splittedDateAndLocation[0].replace(
        "_",
        " "
    )}</p>`;
    dateAndLocation.insertAdjacentHTML(
        "afterend",
        `<p class="capitalize">${splittedDateAndLocation[1].replace(
            "_",
            " "
        )}</p>`
    );
});

