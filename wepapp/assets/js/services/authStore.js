const authStore =  {
    _hasToken: () => {
        const storage = authStore._storage()
        return  storage != null && storage != ""
    },
    _auth: () => {
      const user = authStore._storage() ? JSON.parse(authStore._storage()) : null;
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
