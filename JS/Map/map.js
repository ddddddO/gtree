// https://www.tam-tam.co.jp/tipsnote/javascript/post7755.html
function init() {
    var map = new google.maps.Map(document.getElementById("gmap"), {
        center: { lat: 35.630135, lng: 139.5501 },
        zoom: 13
    });

    var marker1 = new google.maps.Marker({
        position: {
            lat: 35.629559,
            lng: 139.549676
        },
        map: map
    })
    var info1 = new google.maps.InfoWindow({
        content: 'relax space'
    })
    marker1.addListener('click', function () {
        info1.open(map, marker1);
    })
}
