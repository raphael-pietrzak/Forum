
var typeUser = "";
var a = true;
var id_div = 5;


function NewComment(id) {
  if (a == true) {
    var input = document.createElement("div");
    input.className = "input";

    var formCom = document.createElement("form");
    formCom.action = "/addcomment";
    formCom.method = "POST";

    var inputText = document.createElement("input");
    inputText.type = "text";
    inputText.name = "comment_content";
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

    var comments = document.getElementsByClassName("comment");
    console.log(comments);
    for (var i = 0; i < comments.length; i++) {
      comments[i].innerHTML = "";
    }
    document.getElementById("comment" + id).appendChild(input);
  } else {
    a = true;
  }
}


function SelectFilters() {
  if (document.getElementById("filters").style.display == "none") {
    document.getElementById("filters").style.display = "block";
  } else {
    document.getElementById("filters").style.display = "none";
  }
}

function modify(id){
  var elements = document.getElementsByClassName('modify-post');

  for (var i = 0; i < elements.length; i++){
      elements[i].style.display = "none";
  }
  if  (document.getElementById(id).style.display == "none"){
    document.getElementById(id).style.display = "block";
  } else {
    document.getElementById(id).style.display = "none";
  }
}

function UserPermissions(type){
  typeUser = type
  if (type == "guest") {
    document.getElementById("addpost").style.display = "none";
    document.getElementById("name").style.display = "none";
  } else {
    document.getElementById("addpost").style.display = "block";
    document.getElementById("user-profile").style.display = "none";
  }
}

function Like(){
  document.querySelectorAll(".fa-heart")
  .forEach(function(el) {  
  el.addEventListener("click",
      function() {
      if (typeUser != "guest"){
        if (this.classList.contains("unlike")){
          this.classList.toggle("like");
        } else {
          this.classList.toggle("unlike");
        }
        a = false;
        document.getElementById("inputlike").value = this.classList
        document.getElementById("submitlike").submit();
      }
      });
  });
}

function replace(id){
  if (id != id_div){
    document.getElementById('div' + id).style.display = "flex";
    document.getElementById('div' + id_div).style.display = "none";
    id_div = id;
  }
}


function modifypost(id){
  //replace div by an input
  var elements = document.getElementById('post-content'+id);
  if (elements.textContent.replace(/[\n\r]+|[\s]{2,}/g, ' ').trim() != ""){
    elements.innerHTML = '<input type="text" name="post_content" class="modify-text" autofocus value="' + elements.textContent.replace(/[\n\r]+|[\s]{2,}/g, ' ').trim() + '">';
  }
}
