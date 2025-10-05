window.onload = () => {
    console.log(auth.token())
    if (!auth.hasToken()) {
        window.location = "/login"
    }

    api.get("/post")
}