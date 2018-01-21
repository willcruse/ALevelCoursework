var timerInterval = -1;
var finishedInterval;
var i = -1; //timer interval check to disallow more than one instance of the timer decrease to allow
var s = 0; //seconds variable representing remaining seconds on the timer
var m = 0; //minutes variable representing remaining minutes on the timer
var h = 0; //hours variable representing remaining hours on the timer
var iH = 0; //input hours that gets the number on the hours in the webpage
var iM = 0; //input minutes that gets the number on the hours in the webpage
var iS = 0; //input seconds that gets the number on the hours in the webpage
var TMS = -1;


function startTimer() {
    if (i == -1) {
        setInputs();
        TMS = (iH * 3600000) + (iM * 60000) + (iS * 1000) - 1000;
        i = setInterval(decrement, 1000);
        clearNums();
        updateDisplay();
    } else if (i == -2) {
        i = setInterval(decrement, 1000);
    }
}

function stopTimer() {
    clearInterval(i);
    i = -2;
}

function resetTimer() {
    clearInterval(i);
    i = -1;
    addNums();
}

function decrement() {
    if (TMS != 0) {
        TMS = TMS - 1000;
        updateDisplay();
    } else {
        finished();
        clearInterval(i);
    }
}

function updateDisplay() {
    //finds how many hours, minutes and seconds are in the TMS
    var n = TMS; //n represents the TMS in a local variable
    h = Math.floor(n / 3600000); //divides n by the number of ms in a hour then floors it to a whole number
    n = n - (h * 3600000); //takes the total hours in ms from the TMS to allow it not to be doubled
    m = Math.floor(n / 60000); //does the same operation on minutes
    n = n - (m * 60000); //takes the minutes in ms from the TMS remaining
    s = Math.floor(n / 1000); //repeats the same on seconds but I do not take these from n as not needed
    var timerP = document.getElementById("timerP");
    timerP.textContent = (h + " : " + m + " : " + s);
}

function finished() {
    var audio = new Audio("/scripts/finishSound.mp3");
    audio.play();
}


function clearNums() {
    function setBlank(id) {
        var a = document.getElementById(id);
        a.textContent = "";
    }
    setBlank("h2Im");
    setBlank("h1Im");
    setBlank("m2Im");
    setBlank("m1Im");
    setBlank("s2Im");
    setBlank("s1Im");
    setBlank("emptyP");
    setBlank("emptyP1");
}

function addNums() {
    function setZero(id) {
        var a = document.getElementById(id);
        a.textContent = "0";
    }

    function setColon(id) {
        var a = document.getElementById(id);
        a.textContent = ":";
    }
    var n = document.getElementById("timerP");
    n.textContent = "";
    setZero("h2Im");
    setZero("h1Im");
    setZero("m2Im");
    setZero("m1Im");
    setZero("s2Im");
    setZero("s1Im");
    setColon("emptyP");
    setColon("emptyP1")
}

//Code below for inputting numbers
function increase(i) {
    var n = i.id;
    var k = document.getElementById(n);
    var j = k.textContent;
    j = parseInt(j, 10);
    if (n.charAt(0) == "h" || n.charAt(1) == "1") {
        if (j < 9) {
            j++;
        } else {
            j = 0;
        }
    } else {
        if (j < 5) {
            j++;
        } else {
            j = 0;
        }
    }
    k.textContent = j;
}

function decrease(i) {
    var n = i.id;
    var k = document.getElementById(n);
    var j = k.textContent;
    j = parseInt(j, 10);
    if (j > 0) {
        j--;
    } else {
        j = 0;
    }
    k.textContent = j;
}

function setInputs() {
    iH = parseInt((document.getElementById("h2Im").textContent) + (document.getElementById("h1Im").textContent), 10);
    iM = parseInt((document.getElementById("m2Im").textContent) + (document.getElementById("m1Im").textContent), 10);
    iS = parseInt((document.getElementById("s2Im").textContent) + (document.getElementById("s1Im").textContent), 10);
}
