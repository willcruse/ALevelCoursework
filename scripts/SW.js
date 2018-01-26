var timerInterval = -1; //Defines needed global variables
var i = -1;
var ms = 00;
var s = 0;
var m = 0;
var h = 0;

function startSW() { //Function to start the stopwatch with a timer interval of 10ms
    if (i == -1) {
        timerInterval = setInterval(increment, 10);
        i = timerInterval;
    }
}

function stopSW() {
    clearInterval(i); //Clears the interval
    i = -1; //Sets i to ready state
}

function resetSW() {
    clearInterval(timerInterval); //Clears timer interval
    ms = 0;
    s = 0;
    m = 0;
    h = 0; //Changes all time elements to 0
    var timerPJS = document.getElementById("SWP");
    timerPJS.textContent = (h + " : " + m + " : " + s + " . " + ms); //Updates html
    i = -1; //Sets i to ready state
}

function increment() {
    var timerPJS = document.getElementById("SWP");
    timerPJS.textContent = (h + " : " + m + " : " + s + " . " + ms);
    ms++;
    if (ms == 99) { //Increments ms and checks to see if the next variable needs increasing
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





