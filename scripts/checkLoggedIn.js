window.onload = function () {
    var uID = -1;
    var x = decodeURIComponent(document.cookie); //Finds the uID cookie if there
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
    if (uID != -1){ //If the cookie is found updates to tell user they are logged in
        document.getElementById("infoP").innerText = "you are logged in";
    }
};