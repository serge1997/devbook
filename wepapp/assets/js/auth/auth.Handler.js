
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
        console.log(data)
        api.post("/register", data)
        .then(response => {
            console.log(response)
        })
    })
}

if (LoginBtn){
    LoginBtn.addEventListener("click", function(e) {
        console.log("login user")
    })
}