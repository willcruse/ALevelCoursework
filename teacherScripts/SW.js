var timerInterval = -1;
var i = -1;
var ms = 00;
var s = 0;
var m = 0;
var h = 0;

function startSW() {
    if (i == -1) {
        timerInterval = setInterval(increment, 10);
        i = timerInterval;
    }
}

function stopSW() {
    clearInterval(i);
    i = -1;
}

function resetSW() {
    console.log("reset");
    clearInterval(timerInterval);
    ms = 0;
    s = 0;
    m = 0;
    h = 0;
    var timerPJS = document.getElementById("SWP");
    timerPJS.textContent = (h + " : " + m + " : " + s + " . " + ms);
    i = -1;
}

function increment() {
    var timerPJS = document.getElementById("SWP");
    timerPJS.textContent = (h + " : " + m + " : " + s + " . " + ms);
    ms++;
    if (ms == 99) {
        ms = 0;
        s++;
    }
    if (s == 60) {
        m++;
        s = 0;
    }
    if (m == 60) {
        h++;
        m = 0;
    }
}





