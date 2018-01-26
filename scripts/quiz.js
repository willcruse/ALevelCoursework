function genQuiz(id) {
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/games/getFirstTerm', true);
  var obj = new Object();
  obj.ID = id;
  var json = JSON.stringify(obj);
  xhr.onload = function() {
    console.log(xhr.responseText);
    var jsonRes = JSON.parse(xhr.responseText);
    var termFirst = jsonRes.terms;
    var newHTML = "<tr><th>First</th><th>Second</th><th>Answer</th><th>Correct?</th>";
    for (var i = 0; i < termFirst.length; i++) {
      var tempHTML = "<tr><td>";
      tempHTML += termFirst[i];
      tempHTML += "</td><td><input type='text' id='";
      tempHTML += i;
      tempHTML += "'></input></td><td><p></p></td><td><p></p></td></tr>";
      newHTML += tempHTML;
    }
    document.getElementById("quizTable").innerHTML = newHTML;
  };
  xhr.send(json);
}

function check(id){
  var answers = [];
  var rowLen = (document.getElementById("quizTable").getElementsByTagName("tr").length) -1;
  for (var i = 1; i <= rowLen; i++) {
    answers.push(document.getElementById("quizTable").rows[i].cells.item(1).firstChild.value);
  }
  var obj = new Object();
  obj.id = id;
  obj.ans = answers;
  console.log("Obj", obj);
  var json = JSON.stringify(obj);
  console.log("json", json);
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/games/checkQuizRes', true);
  xhr.onload = function (){
      console.log(xhr.responseText);
      var jsonR = JSON.parse(xhr.responseText);
      var correct = jsonR.corArr;
      var ans = jsonR.ansArr;
      var score = jsonR.score;
      document.getElementById("infoP").innerHTML = score;
      for (var i = 1; i <= rowLen; i++){
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

