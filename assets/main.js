isLoggedIn();

function isLoggedIn() {
    const token = getCookie("token");
    // verify token
    if (!token) {
        window.location.replace('/assets/login.html');
    }
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/check_token", true);
    xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
    xhr.setRequestHeader("Authorization", "Bearer " + token);
    xhr.send();
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
    const date   = new Date(exerciseForm.date.value).getTime() / 1000;
    submitExercise(workout, weight, date); 
})

function submitExercise(workout, weightvar, datetime) {
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/workout", true);
    xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
    xhr.setRequestHeader("Authorization", "Bearer " + getCookie("token"));
    xhr.send(JSON.stringify({
	exercise: workout,
	weight: weightvar,
	date: datetime
    }))
    xhr.onload = function() {
	console.log("done");
	location.reload();
    }
}


