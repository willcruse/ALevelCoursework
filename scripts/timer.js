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
    if (i == -1) { //Checks i is in ready state
        setInputs();
        TMS = (iH * 3600000) + (iM * 60000) + (iS * 1000) - 1000; //Calculates number of ms
        i = setInterval(decrement, 1000); //Sets interval to go every 1s
        clearNums(); //Clears the display
        updateDisplay(); //Updates the display
    } else if (i == -2) {//Allows for resuming of timer once paused
        i = setInterval(decrement, 1000);
    }
}

function stopTimer() {
    clearInterval(i); //Clears interval to stop decrementation
    i = -2; //Sets i to pause state
}

function resetTimer() {
    clearInterval(i); //Clears interval
    i = -1; //Set i to the ready state
    addNums(); //Populates the display
}

function decrement() {
    if (TMS != 0) { //Decrements the TMS then updates display
        TMS = TMS - 1000;
        updateDisplay();
    } else {
        finished(); //Calls finished func
        clearInterval(i); //Clears the interval
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
    var audio = new Audio("/scripts/finishSound.mp3"); //Locates the finished sound then plays it
    audio.play();
}


function clearNums() {
    function setBlank(id) {
        var a = document.getElementById(id); //Gets the element and sets to empty string
        a.textContent = "";
    }
    setBlank("h2Im"); //Sets everything to blank
    setBlank("h1Im");
    setBlank("m2Im");
    setBlank("m1Im");
    setBlank("s2Im");
    setBlank("s1Im");
    setBlank("emptyP");
    setBlank("emptyP1");
}

function addNums() {
    function setZero(id) { //Gets the element and sets to 0
        var a = document.getElementById(id);
        a.textContent = "0";
    }

    function setColon(id) { //Gets the element and sets to :
        var a = document.getElementById(id);
        a.textContent = ":";
    }
    var n = document.getElementById("timerP");
    n.textContent = "";
    setZero("h2Im"); //For each element does the right thing
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
    if (n.charAt(0) == "h" || n.charAt(1) == "1") { //Makes sure time cant be made to not real values
        if (j < 9) { //Cant exceed 9
            j++;
        } else {
            j = 0;
        }
    } else { //Cant exceed 5
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
    j = parseInt(j, 10); //Makes sure time is above 0
    if (j > 0) {
        j--;
    } else {
        j = 0;
    }
    k.textContent = j;
}

function setInputs() {
    iH = parseInt((document.getElementById("h2Im").textContent) + (document.getElementById("h1Im").textContent), 10); //Gets the input from the display
    iM = parseInt((document.getElementById("m2Im").textContent) + (document.getElementById("m1Im").textContent), 10);
    iS = parseInt((document.getElementById("s2Im").textContent) + (document.getElementById("s1Im").textContent), 10);
}
