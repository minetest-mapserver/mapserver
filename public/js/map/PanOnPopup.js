export default function (map) {
    map.on("popupopen", (e) => {
        map.panTo(e.popup.getLatLng(), {
            duration: 0.5,
        });
    })
}
