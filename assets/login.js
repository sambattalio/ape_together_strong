// modified but based off of same medium article in login.html
const loginForm = document.getElementById("login-form");
const loginButton = document.getElementById("login-form-submit");
const loginErrorMsg = document.getElementById("login-error-msg");

loginButton.addEventListener("click", (e) => {
    e.preventDefault();
    const username = loginForm.username.value;
    const password = loginForm.password.value;
    basicLogin(username, password);
    if (username === "sigma" && password === "male") {
        alert("You have successfully logged in.");
        location.reload();
    } else {
        loginErrorMsg.style.opacity = 1;
    }
})

async function basicLogin (email, password) {
    // based on concepts from https://zellwk.com/blog/frontend-login-system/
    const response = await fetch('/attempt-login', {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        body: {
            username: email,
            password: password
        }
    });
    return response.json();
}
