function genQuiz(id) {
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/game/getFirstTerm', true);
  xhr.onload = function() {
    var jsonRes = JSON.parse(xhr.responseText);
    var termFirst = jsonRes.term;
    var newHTML = "<tr><th>First<th><Second</th><th>Check</th><th>Answer</th><th>Correct?</th>";
    for (var i = 0; i < termFirst.length; i++) {
      var tempHTML = "<tr><th>";
      tempHTML += termFirst[i];
      tempHTML += "</th><input type='text' id='";
      tempHTML += i;
      tempHTML += "'></input></th><th></th></tr>";
      newHTML += tempHTML;
    }
    document.getElementById("quizTable").innerHTML = newHTML;
  }
  xhr.send(id);
}

function check(id){
  var answers = [];
  var rowLen = (document.getElementById("quizTable").getElementByTagName("tr").length) -1;
  for (var i = 1; i < rowLen; i++) {
    answers.push(document.getElementById("quizTable").rows[i].cells.item(1).innerHTML)
  }
  var obj = new Object();
  obj.id = id;
  obj.answers = answers;
  var xhr = new XMLHttpRequest();
  xhr.open('POST', '/game/checkQuizRes', true);
  xhr.onload = function (){
    var jsonR = JSON.parse(xhr.responseText);
    var correct = jsonR.corArr;
    var ans = jsonR.ansArr;
    var score = jsonR.score;
    document.getElementById(infoP).innerHTML = score;
    for (var i = 1; i < rowLen; i++){
      document.getElementById("quizTable").rows[i].cells.item(2).innerHTML = ansArr[i-1];
      if (correct[i-1]) {
        document.getElementById("quizTable").rows[i].cells.item(3).innerHTML = "Corrrect";
      }else{
        document.getElementById("quizTable").rows[i].cells.item(3).innerHTML = "Incorrect";
      }
    }
  }
}
