isLoggedIn();

function isLoggedIn() {
    console.log("checking if logged in");
    const token = getCookie("token");
    // verify token
    if (!token) {
        window.location.replace('/assets/login.html');
    }
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/check_token", true);
    xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
    xhr.send(JSON.stringify({
            jwttoken: token
    }))
    xhr.onload = function() {
      // new jwt
      console.log("response", this.responseText);
      if (this.responseText === "") {
        window.location.replace('/assets/login.html')
      }
    }
}

const exerciseButton = document.getElementById("submit");
const exerciseForm   = document.getElementById("exercise");

exerciseButton.addEventListener("click", (e) => {
    e.preventDefault();
    const workout = exerciseForm.exercise.value;
    const weight = exerciseForm.weight.value;
    submitExercise(workout, weight); 
})

function submitExercise(workout, weightvar) {
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/workout", true);
    xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
    xhr.setRequestHeader("Authorization", "Bearer 20340230413204");
    xhr.send(JSON.stringify({
	exercise: workout,
	weight: weightvar
    }))
    xhr.onload = function() {
    }
}


