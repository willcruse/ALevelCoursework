function genQuiz(id) { //Generates the quiz html
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/games/getFirstTerm', true); //XHR to get the first termss
  var obj = new Object();
  obj.ID = id;
  var json = JSON.stringify(obj); //Creates json out of the id
  xhr.onload = function() {
    var jsonRes = JSON.parse(xhr.responseText); //Parses response text to json
    var termFirst = jsonRes.terms; //Extracts terms array from json response
    var newHTML = "<tr><th>First</th><th>Second</th><th>Answer</th><th>Correct?</th>"; //Sets table headers
    for (var i = 0; i < termFirst.length; i++) { //Iterates for each item in termFirst array appending it to the quiz table
      var tempHTML = "<tr><td>";
      tempHTML += termFirst[i];
      tempHTML += "</td><td><input type='text' id='";
      tempHTML += i;
      tempHTML += "'></input></td><td><p></p></td><td><p></p></td></tr>";
      newHTML += tempHTML;
    }
    document.getElementById("quizTable").innerHTML = newHTML; //Adds the newHTML to the quizTable
  };
  xhr.send(json);
}

function check(id){
  var answers = [];
  var rowLen = (document.getElementById("quizTable").getElementsByTagName("tr").length) -1; //Finds how many tows there are
  for (var i = 1; i <= rowLen; i++) {
    answers.push(document.getElementById("quizTable").rows[i].cells.item(1).firstChild.value); //For each row except the header row iterates over appending the input to the rows table
  }
  var obj = new Object(); //Creates a json object then appends with setID and user answers
  obj.id = id;
  obj.ans = answers;
  var json = JSON.stringify(obj);
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/games/checkQuizRes', true); //XHR to get the answers, score and number correct
  xhr.onload = function (){
      var jsonR = JSON.parse(xhr.responseText); //Get json and the data sent
      var correct = jsonR.corArr;
      var ans = jsonR.ansArr;
      var score = jsonR.score;
      document.getElementById("infoP").innerHTML = "Score: " + score + "/" + correct.length; //Adds score to the infoP
      for (var i = 1; i <= rowLen; i++){ //Iterates adding the answer and correct/incorrect to the table
        document.getElementById("quizTable").rows[i].cells.item(2).innerHTML = "<p>" + ans[i-1] + "</p>";
        if (correct[i-1]) {
            document.getElementById("quizTable").rows[i].cells.item(3).innerHTML = "<p>Correct</p>";
        }else{
            document.getElementById("quizTable").rows[i].cells.item(3).innerHTML = "<p>Incorrect</p>";
        }
      }
  };
  xhr.send(json);
}

