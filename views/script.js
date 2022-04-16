

var a = true;


function NewComment(id) {
  if (a == true) {
    // .forEach(function(el) {  
    // el.addEventListener("click",
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


function UserPermissions(type){
  if (type == "guest") {
    document.getElementById("addpost").style.display = "none";
    document.getElementById("name").style.display = "none";
  } else {
    document.getElementById("addpost").style.display = "block";
    document.getElementById("user-profile").style.display = "none";
  }
}

function Like(){
  //favorite heart like/unlike
  document.querySelectorAll(".fa-heart")
  .forEach(function(el) {  
  el.addEventListener("click",
      function() {
      if (this.classList.contains("unlike"))
          this.classList.toggle("like");
      else this.classList.toggle("unlike");
      a = false;
      });
  });

}
