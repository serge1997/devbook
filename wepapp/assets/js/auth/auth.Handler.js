
const createAccountBtn = document.querySelector("#register-btn")
const LoginBtn = document.querySelector("#login-btn")
if (createAccountBtn){
    
    createAccountBtn.addEventListener("click", function(e) {
        const data = {
            name: document.getElementById("name").value,
            email: document.getElementById("email").value,
            password: document.getElementById("password").value,
            nick: document.getElementById("username").value
        }
        api.post("/register", data)
        .then(response => {
            alert(response.message)
            window.location.href = "/login"
        })
    })
}

if (LoginBtn){
    LoginBtn.addEventListener("click", function(e) {
        console.log("login user")
    })
}