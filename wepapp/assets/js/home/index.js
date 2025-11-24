

window.onload = () => {
    const logoutBtn = document.querySelector('#logout-btn')
    if (!auth.hasToken()) {
        window.location = "/login"
    }

    api.get("/post")
  if (logoutBtn) {
    logoutBtn.addEventListener("click", async () => {
      await api.post("/logout")
      .then(res => {
        console.log(res)
      })
      window.location = "/login"

    })
  }
}
