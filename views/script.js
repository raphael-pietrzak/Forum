
var typeUser = "";
var a = true;
var id_div = 1;


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
    document.getElementById('div' + id).style.display = "block";
    document.getElementById('div' + id_div).style.display = "none";
    id_div = id;
  }
}

// function DisplayComments(id){
//   if (document.getElementById('commentaires'+id).style.display == 'flex') {
//     document.getElementById('commentaires'+id).style.display = 'none'
//     document.getElementsByClassName('showComments')[id-1].innerHTML = "<a onclick='DisplayComments(id)' id='{{.Pid}}' class='showComments'>Show Comments</a>"
//   } else {
//     document.getElementById('commentaires'+id).style.display = 'flex'
//     document.getElementsByClassName('showComments')[id-1].innerHTML = "<a onclick='DisplayComments(id)' id='{{.Pid}}' class='showComments'>Hide Coms</a>"
//   }
// }

function onSignIn(googleUser) {
  console.log("hahahahahaha")
  var profile = googleUser.getBasicProfile();
  console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
  console.log('Name: ' + profile.getName());
  console.log('Image URL: ' + profile.getImageUrl());
  console.log('Email: ' + profile.getEmail()); // This is null if the 'email' scope is not present.
}

function getBasicProfile(){
  console.log("heeeheheheehhe")
  gapi.auth2.authorize({
    client_id: '822476215105-a1qeg4jvqdnjlut874jucd3lh0srpfr8.apps.googleusercontent.com',
    scope: 'email profile openid',
    response_type: 'id_token permission'
  }, function(response) {
    if (response.error) {
      // An error happened!
      return;
    }
    // The user authorized the application for the scopes requested.
    var accessToken = response.access_token;
    var idToken = response.id_token;
    // You can also now use gapi.client to perform authenticated requests.
  });
}
