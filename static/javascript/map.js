
// initialize communication with the platform
// In your own code, replace variable window.apikey with your own apikey
var platform = new H.service.Platform({
    apikey: process.env.MAP_KEY,
});

var defaultLayers = platform.createDefaultLayers();
document
    .getElementById("gt-map-container-resize")
    .addEventListener("click", () => {
        document.getElementById("gt-map").classList.toggle("absolute");
        map.getViewPort().resize();
    });

function addMarkerToGroup(group, coordinate, html) {
    var marker = new H.map.Marker(coordinate);
    // add custom data to the marker
    marker.setData(html);
    group.addObject(marker);
}
/**
 * Add two markers showing the position of Liverpool and Manchester City football clubs.
 * Clicking on a marker opens an infobubble which holds HTML content related to the marker.
 * @param {H.Map} map A HERE Map instance within the application
 */
async function addInfoBubble(map) {
    let res;
    let path = window.location.href.match(/artists\/\d/);
    if (path == null) {
        res = await fetch("allartists");
    } else {
        res = await fetch("/singleartist/" + path[0].match(/\d/)[0]);
    }
    let data = await res.json();
    data = Array.isArray(data) ? data : [data];

    var group = new H.map.Group();
    map.addObject(group);
    // add 'tap' event listener, that opens info bubble, to the group
    group.addEventListener(
        "tap",
        function (evt) {
            // event target is the marker itself, group is a parent event target
            // for all objects that it contains
            var bubble = new H.ui.InfoBubble(evt.target.getGeometry(), {
                // read custom data
                content: evt.target.getData(),
            });
            // show info bubble
            ui.addBubble(bubble);
        },
        false
    );

    let response = await fetch("/map")
    let storedData = await response.json()
 
    storedData.forEach((value) => {
        var artist = [];

        for (var i = 0; i < data.length; i++) {
            if (
                data[i].Location.locations.includes(value.location)
            ) {
                artist.push(
                    `<div style="display: flex; justify-content: flex-start; align-items:center; gap:4px; width: 100%;">
                        <img style="width: 2.2rem; height: 2.2rem; border-radius:50%;" src="${data[i].Artist.image}" alt=${data[i].Artist.name} />
                        <h4 style="font-weight: bold; color: #0c840c; line-height: initial;">${data[i].Artist.name}</h4>
                    </div>`
                );
            }
        }

        if (artist.length !== 0) {
            addMarkerToGroup(
                group,
                {
                    lat: value.LocationInformation.position?.lat,
                    lng: value.LocationInformation.position?.lng,
                },
                `<div class="w-[180px]">
                    <div class="flex gap-[4px]">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 mt-[6px]">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z" />
                  </svg>
                        <h3 class="locations text-green-400" >${
                            formatText(value.location)
                        }</h3>
                    </div>
                    <div style="display: flex; flex-wrap: wrap; justify-content: center; align-items:center; gap: 10px;">
                        ${artist.join(" ")}
                    </div>
                </div>`
            );
        }
    });
    
}
/**
 * Boilerplate map initialization code starts below:
 */
// initialize a map - this map is centered over Europe
var map = new H.Map(
    document.getElementById("gt-map-container"),
    defaultLayers.vector.normal.map,
    {
        // center: { lat: 53.43, lng: -2.961 },
        zoom: 2,
        pixelRatio: window.devicePixelRatio || 1,
    }
);
// add a resize listener to make sure that the map occupies the whole container
window.addEventListener("resize", () => map.getViewPort().resize());
// MapEvents enables the event system
// Behavior implements default interactions for pan/zoom (also on mobile touch environments)
var behavior = new H.mapevents.Behavior(new H.mapevents.MapEvents(map));
// create default UI with layers provided by the platform
var ui = H.ui.UI.createDefault(map, defaultLayers);
// Now use the map as required...
addInfoBubble(map);
function formatText(str) {
    splittedDateAndLocation = str.split("-")
    return `<div class="mb-2 ">
            <p class="font-bold text-green-500 capitalize">${splittedDateAndLocation[1].replace("_"," ")}</p>
            <p class="capitalize text-xs text-gray-500">${splittedDateAndLocation[0].replace("_"," ")}</p>
    </div>
    `
}