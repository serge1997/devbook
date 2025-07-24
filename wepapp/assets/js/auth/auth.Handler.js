
const createAccountBtn = document.querySelector("#register-btn")
const LoginBtn = document.querySelector("#login-btn")

if (createAccountBtn){
    
    createAccountBtn.addEventListener("click", function(e) {
        api.post("/register", {name: null})
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