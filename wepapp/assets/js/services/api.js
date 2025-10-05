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
            const options = {
                method: method,
                headers: {
                    'Content-Type': 'application/json',
                    ...headers
                }
            }
            if (method != "GET") {
                options["body"] = JSON.stringify(data)
            }
            fetch(url, options)
            .then(async response => {
                if (response.status >= 400) {
                    const erro = await response.json()
                    reject(erro)
                    alert(erro.message)
                }else{
                    resolve(response.json())
                }
            })
            
        })
    }
}

const api = new Api()