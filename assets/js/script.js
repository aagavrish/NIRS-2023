let button = document.querySelector(".inputblock > button");
if (button) {
    button.onclick = function (e) {
        let District = document.querySelector('input[name="District"]');
        let Region = document.getElementById("Region")

        let data = {}
        data[Region.name] = Region.value;
        data[District.name] = District.value;
        
        let xhr = new XMLHttpRequest();

        xhr.open("POST", "/calculation");
        xhr.onload = function (e) {
            window.location.href="/";
        }

        xhr.send(JSON.stringify(data))
    }
}