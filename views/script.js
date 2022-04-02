

function NewComment(id) {
  //add input div
  var input = document.createElement("div");
  input.className = "input";
  var inputText = document.createElement("input");
  inputText.type = "text";
  inputText.placeholder = "Add a comment...";
  input.appendChild(inputText);
  var inputButton = document.createElement("button");
  inputButton.innerHTML = "Post";
  input.appendChild(inputButton);
  console.log(id);
  document.getElementById("comments"+id).appendChild(input);
}

function addElement() {
  var newDiv = document.createElement("div");
  var newContent = document.createTextNode('Hi there and greetings!');
  newDiv.appendChild(newContent);
  var currentDiv = document.getElementById('container');
  document.body.insertBefore(newDiv, currentDiv);
}