// modified but based off of same medium article in login.html
const loginForm = document.getElementById("login-form");
const loginButton = document.getElementById("login-form-submit");
const loginErrorMsg = document.getElementById("login-error-msg");

isLoggedIn();

loginButton.addEventListener("click", (e) => {
    e.preventDefault();
    const username = loginForm.username.value;
    const password = loginForm.password.value;
    basicLogin(username, password); 
})

function isLoggedIn() {
    const token = getCookie("token");
    // verify token
    if (!token) return;
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/check_token", true);
    xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
    xhr.send(JSON.stringify({
            jwttoken: token
    }))
    xhr.onload = function() {
      // new jwt
      console.log(this.responseText);
      if (this.responseText != "") {
        window.location.replace('/assets/')
      }
    }
}

// w3schools lol
function getCookie(cname) {
  var name = cname + "=";
  var decodedCookie = decodeURIComponent(document.cookie);
  var ca = decodedCookie.split(';');
  for(var i = 0; i <ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}

function basicLogin (email, password) {
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/login_attempt", true);
    xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
    xhr.send(JSON.stringify({
            username: email,
            password: password
    }))
    xhr.onload = function() {
      // new jwt
      console.log(this.responseText);
      new_jwt = this.responseText;
      if (new_jwt) {
          console.log(new_jwt);
	  location.reload();
      } else {
	  loginErrorMsg.style.opacity = 1;
      }
    }
   return true;
}
