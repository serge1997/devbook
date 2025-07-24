class Api {

    post(url, data = null, headers = {}, extern = false) {
        return this.send(url, "POST", data, headers, extern)
    }
    get(url, headers = null, extern = false) {
        return this.send(url, "GET", null, headers, extern)
    }
    send(url, method, data = null, headers = {}, extern = false) {

        return new Promise((resolve, reject) => {
            console.log(`===== REQUEST ${method} ${url} =====`)
            fetch(url, {
                body: JSON.stringify(data),
                method: method,
                headers: {
                    'Content-Type': 'application/json',
                    ...headers
                }
            })
            .then(response => {
                resolve(response.json())
            })
            .catch(error => {
                reject(error)
                console.log(error)
            });
            
        })
    }
}

const api = new Api()