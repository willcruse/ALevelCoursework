window.onload = getSetsTable();

function getSetsTable() {
    var uID;
    var cookieName = "uID=";
    var cookies = decodeURIComponent(document.cookie);
    var cookiesSplit = cookies.split(';');
    for (var i = 0; i < cookiesSplit.length; i++) {
        var c = cookiesSplit[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(cookieName) == 0) {
            uID = c.substring(cookieName.length, c.length);
        }
    }
    console.log(uID);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/setsPage/getSets', true);
    xhr.setRequestHeader("Content-Type", "text; charset=UTF-8");
    xhr.onload = function() {

    }
    xhr.send(uID);
}