window.onload = getSets();

function getSets() {
    var uID = getCookieUID();
    var xhr = new XMLHttpRequest();
    var obj = new Object(); //Creates new object with the userID
    obj.uID = uID;
    var json = JSON.stringify(obj);
    xhr.open('POST', '/setsPage', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function() {
        var newHTML = "<tr><th>SetName</th><th>View</th><th>Delete</th><th>quiz!</th></tr>"; //Creates table header
        var table = document.getElementById("mySetsTable");
        var jsonResp = JSON.parse(xhr.responseText);
        var jsonSets = jsonResp.sets;
        if (jsonSets == null) { //If no sets returned adds add term row to the table
            newHTML += "<tr><td><input id='setName' placeholder='setName'></input></td><td><button onclick='newSets();'>Add</button></td></tr>";
            table.innerHTML = newHTML; //Adds to the page
            return;
        }
        for (var i = 0; i < jsonSets.length; i++) { //For each value returned adds row to setsTable with option to view and delete
            var tempArr = jsonSets[i];
            var tempHTML = "</br><tr><td>";
            tempHTML += tempArr[1];
            tempHTML += "</td><td>";
            tempHTML += "<button onclick='view(";
            tempHTML += tempArr[0];
            tempHTML += ");'>view</button></td>";
            tempHTML += "<td><button onclick='deleteSets(";
            tempHTML += tempArr[0];
            tempHTML += ");'>Delete</button>";
            tempHTML += "<td><button onclick='quiz(";
            tempHTML += tempArr[0];
            tempHTML += ");'>quiz!</button>";
            newHTML += tempHTML;
        }
        newHTML += "<tr><td><input id='setName' placeholder='setName'></input></td><td><button onclick='newSets();'>Add</button></td></tr>"; //Adds add set row
        table.innerHTML = newHTML; //Updates to table
    };
    xhr.send(json);
}

function view(id) {
    var obj = new Object(); //Creates a new object and adds setID to it then turns to JSON
    obj.SetID = id;
    var json = JSON.stringify(obj);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/setsPage/getTerms', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function() {
        var jsonRes = JSON.parse(xhr.responseText); //Gets json from resp text
        var JSONterms = jsonRes.terms;
        if (JSONterms == null) { //If no terms returned adds the add term row
            var newHTML = "<tr><th>First Term</th><th>Second Term</th><th>Delete Term</th></tr>";
            newHTML += "</br><tr><td><input type='text' placeholder='TermA' id='newTermA'></input></td><td><input type='text' placeholder='TermB' id='newTermB'></input></td><td><button onclick='newTerms(";
            newHTML += id;
            newHTML += ");'>Add</button></td></tr>";
            document.getElementById("termsTable").innerHTML = newHTML; //Updates table with the new html
            return;
        }
        var newHTML = "<tr><th>First Term</th><th>Second Term</th><th>Delete Term</th></tr>"; //Adds header
        for (var i = 0; i < JSONterms.length; i++) { //For each item in terms array appends to the table
            var tempArr = JSONterms[i];
            var tempHTML = "</br><tr><td>"; //Not most efficient but clear to see what is happening therefore easier to program
            tempHTML += tempArr[0];
            tempHTML += "</td><td>";
            tempHTML += tempArr[1];
            tempHTML += "</td><td><button onclick='deleteTerms(`";
            tempHTML += tempArr[0];
            tempHTML += "`, `";
            tempHTML += tempArr[1];
            tempHTML += "`, ";
            tempHTML += id;
            tempHTML += ");'>Delete</button></td></tr>";
            newHTML += tempHTML;
        }
        newHTML += "</br><tr><td><input type='text' placeholder='TermA' id='newTermA'></input></td><td><input type='text' placeholder='TermB' id='newTermB'></input></td><td><button onclick='newTerms(";
            newHTML += id;
            newHTML += ");'>Add</button></td></tr>"; //Adds new terms row
        document.getElementById("termsTable").innerHTML = newHTML; //Updates table
    };
    xhr.send(json)
}

function deleteTerms(fir, sec, id) { //Function to delete terms
    var obj = new Object(); //Creates object which then has setID and term added
    var tempArr = [fir, sec];
    obj.setID = id;
    obj.term = tempArr;
    var json = JSON.stringify(obj); //Turns to json
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/setsPage/deleteTerms', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function() {
        view(id); //Onload calls view to update users screen
    };
    xhr.send(json);
}

function newTerms(id){
    var termA = document.getElementById("newTermA").value;
    var termB = document.getElementById("newTermB").value;
    var obj = new Object(); //New object which has the terms and the setID appended to it
    obj.termA = termA;
    obj.termB = termB;
    obj.setID = id;
    var uID = getCookieUID()
    obj.uID = uID; //Appends to object
    var json = JSON.stringify(obj);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/setsPage/addTerms', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function() {
        view(id); //Onload updates user perspective
    };
    xhr.send(json);
}

function newSets() {
    var uID = getCookieUID(); //Gets user cookie
    var setName = document.getElementById("setName").value;
    var obj = new Object(); //Creates object with setName and uID appended to it
    obj.SetName = setName;
    obj.UID = uID;
    var jsonObj = JSON.stringify(obj); //Turns object to json
    console.log("JSON ", jsonObj);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/setsPage/newSets', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function () {
        getSets(); //Updates sets
    };
    xhr.send(jsonObj)
}

function deleteSets(id) {
    var uID = getCookieUID();
    var obj = new Object(); //New object which has setID and uID append to it
    obj.SetID = id;
    obj.UID = uID;
    var jsonObj = JSON.stringify(obj); //Converts to JSON
    console.log("JSON ", jsonObj);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/setsPage/deleteSets', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function () {
        getSets(); //Makes sets update onload
    };
    xhr.send(jsonObj)
}

function quiz(id) {
  var obj = new Object(); //Makes new object with setID appended
  obj.id = id;
  var json = JSON.stringify(obj); //Converts to json
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/games/quizMove", true);
  xhr.onload = function (){
    window.location = "/html/cache/" + id + ".html"; //Changes window location to that which the template is
  };
  xhr.send(json);
}

function getCookieUID () {
    var uID = -1;
    var x = decodeURIComponent(document.cookie); //Fetches cookies stored
    var split = x.split(";"); //Splits the string on ;
    var target = "uID="; //Sets target string to uID=
    for (var i = 0; i < split.length; i++) { //For every value in split array length looks for character
        var c = split[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
    }
    if (c.indexOf(target) == 0) { //If the char is 0 gets the substring starting at length of target and finishing at length of c
        uID = c.substring(target.length, c.length);
    }
    return uID;
}