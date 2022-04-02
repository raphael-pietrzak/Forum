function NewComment() {
  var counter = 0;
  var btn = document.getElementById('addComment');
  var form = document.getElementById('addNewComment');
  var addInput = function NewComment() {
    counter++;
    var input = document.createElement("input");
    input.id = 'input-' + counter;
    input.type = 'text';
    input.name = 'name';
    input.placeholder = 'Input number ' + counter;
    form.appendChild(input);
  };
  btn.addEventListener('click', function NewComment() {
    addInput();
  }.bind(this));
}

function addElement() {
  var newDiv = document.createElement("div");
  var newContent = document.createTextNode('Hi there and greetings!');
  newDiv.appendChild(newContent);
  var currentDiv = document.getElementById('container');
  document.body.insertBefore(newDiv, currentDiv);
}