const authStore =  {
    _hasToken: () => {
        const storage = authStore._storage()
        return  storage != null && storage != ""
    },
    _auth: () => {
        const user = JSON.parse(authStore._storage())
        return user;
    },
    _token: () => {
        return authStore._auth()?.token;
    },
    _storage: () => {
        return localStorage.getItem("user")
    }
}

const auth = {
    hasToken: authStore._hasToken,
    token: authStore._token
}