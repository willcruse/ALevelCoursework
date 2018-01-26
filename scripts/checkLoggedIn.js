window.onload = function () {
    var uID = -1;
    var x = decodeURIComponent(document.cookie);
    var split = x.split(";");
    var target = "uID=";
    for (var i = 0; i < split.length; i++) {
        var c = split[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
    }
    if (c.indexOf(target) == 0) {
        uID = c.substring(target.length, c.length);
    }
    if (uID != -1){
        document.getElementById("infoP").innerText = "you are logged in";
    }
};