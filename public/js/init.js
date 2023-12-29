fetch("/init", { headers: { 'accept': 'application/json' } })
    .then(response => response.json())
    .catch(err => {
        msg = "There was a problem getting your IP address"
        document.getElementById("ip").innerHTML = msg
        console.error(err)
    })