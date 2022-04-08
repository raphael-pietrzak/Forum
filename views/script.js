
function NewComment(id) {
  console.log(id)
  //add input div
  var input = document.createElement("div");
  input.className = "input";

  var formCom = document.createElement("form");
  formCom.action = "/addcomment";
  formCom.method = "POST";

  var inputText = document.createElement("input");
  inputText.type = "text";
  inputText.name = "comment_content";
  inputText.autofocus = true;
  inputText.placeholder = "Add a comment...";
  formCom.appendChild(inputText);

  var inputhidden = document.createElement("input");
  // inputhidden.style.display = "none";
  inputhidden.name = "post_id";
  inputhidden.type = "hidden";
  inputhidden.value = id;
  formCom.appendChild(inputhidden);
  
  var inputButton = document.createElement("button");
  inputButton.innerHTML = "Post";
  formCom.appendChild(inputButton);
  input.appendChild(formCom);
  console.log(id);

  //delete div "comments"
  console.log(document.getElementById("comment"));
  var comments = document.getElementsByClassName("comment");
  console.log(comments);
  for (var i = 0; i < comments.length; i++) {
    comments[i].innerHTML = "";
  }
  document.getElementById("comment" + id).appendChild(input);
}


function UserProfile(username) {
  if (username != "") {
  document.getElementById("user-profile").style.display = "none";
  } else {
    document.getElementById("name").style.display = "none";
  }
}






