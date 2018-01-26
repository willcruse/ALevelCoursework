window.onload = getSets();

function getSets() {
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
    var xhr = new XMLHttpRequest();
    var obj = new Object();
    obj.uID = uID;
    var json = JSON.stringify(obj);
    xhr.open('POST', '/setsPage', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function() {
        var newHTML = "<tr><th>SetName</th><th>View</th><th>Delete</th><th>quiz!</th></tr>";
        var table = document.getElementById("mySetsTable");
        var jsonResp = JSON.parse(xhr.responseText);
        var jsonSets = jsonResp.sets;
        if (jsonSets == null) {
            newHTML += "<tr><td><input id='setName' placeholder='setName'></input></td><td><button onclick='newSets();'>Add</button></td></tr>";
            table.innerHTML = newHTML;
            return;
        }
        for (var i = 0; i < jsonSets.length; i++) {
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
        newHTML += "<tr><td><input id='setName' placeholder='setName'></input></td><td><button onclick='newSets();'>Add</button></td></tr>";
        table.innerHTML = newHTML;
    }
    xhr.send(json);
}

function view(id) {
    var obj = new Object();
    obj.SetID = id;
    var json = JSON.stringify(obj);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/setsPage/getTerms', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function() {
        var jsonRes = JSON.parse(xhr.responseText);
        var JSONterms = jsonRes.terms;
        if (JSONterms == null) {
            var newHTML = "<tr><th>First Term</th><th>Second Term</th><th>Delete Term</th></tr>";
            newHTML += "</br><tr><td><input type='text' placeholder='TermA' id='newTermA'></input></td><td><input type='text' placeholder='TermB' id='newTermB'></input></td><td><button onclick='newTerms(";
            newHTML += id;
            newHTML += ");'>Add</button></td></tr>";
            document.getElementById("termsTable").innerHTML = newHTML;
            return;
        }
        var newHTML = "<tr><th>First Term</th><th>Second Term</th><th>Delete Term</th></tr>";
        for (var i = 0; i < JSONterms.length; i++) {
            var tempArr = JSONterms[i];
            var tempHTML = "</br><tr><td>"; //Not most efficient but clear to see what is happening therefore easier to program
            tempHTML += tempArr[0];
            tempHTML += "</td><td>";
            tempHTML += tempArr[1];
            tempHTML += "</td><td><button onclick='deleteTerms(`"
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
            newHTML += ");'>Add</button></td></tr>";
        document.getElementById("termsTable").innerHTML = newHTML;
    }
    xhr.send(json)
}

function deleteTerms(fir, sec, id) {
    var obj = new Object();
    var tempArr = [fir, sec];
    console.log(tempArr);
    obj.setID = id;
    obj.term = tempArr;
    console.log("Object ", obj);
    var json = JSON.stringify(obj);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/setsPage/deleteTerms', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function() {
        view(id);
    }
    xhr.send(json);
}

function newTerms(id){
    var termA = document.getElementById("newTermA").value;
    var termB = document.getElementById("newTermB").value;
    var obj = new Object();
    obj.termA = termA;
    obj.termB = termB;
    obj.setID = id;
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
    obj.uID = uID;
    var json = JSON.stringify(obj);
    console.log("JSON", json);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/setsPage/addTerms', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function() {
        view(id);
    };
    xhr.send(json);
}

function newSets() {
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
    if (uID == -1) {
        return;
    }
    var setName = document.getElementById("setName").value;
    var obj = new Object();
    obj.SetName = setName;
    obj.UID = uID;
    var jsonObj = JSON.stringify(obj);
    console.log("JSON ", jsonObj);
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/setsPage/newSets', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function () {
        var jsonResp = JSON.parse(xhr.responseText);
        console.log("JSON", jsonResp);
        getSets();
    }
    xhr.send(jsonObj)
}

function deleteSets(id) {
    var uID = -1;
    var x = decodeURIComponent(document.cookie);
    var split = x.split(";")
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
    var obj = new Object();
    obj.SetID = id;
    obj.UID = uID;
    var jsonObj = JSON.stringify(obj);
    console.log("JSON ", jsonObj);
    var xhr = new XMLHttpRequest()
    xhr.open('POST', '/setsPage/deleteSets', true);
    xhr.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    xhr.onload = function () {
        getSets();
    }
    xhr.send(jsonObj)
}

function quiz(id) {
  var obj = new Object();
  obj.id = id;
  var json = JSON.stringify(obj);
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/games/quizMove", true);
  xhr.onload = function (){
    console.log(xhr.responseText);
    window.location = "/html/cache/" + id + ".html";
  };
  xhr.send(json);
}
