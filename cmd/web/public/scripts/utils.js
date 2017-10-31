function checkToken(token) {
    if (token == null) {
        window.location = "/";
    }
}

function httpErrorHandler(err, resp) {
    if (resp.statusCode == 401) {
        localStorage.removeItem("token")
        window.location = "/"
        alert("your token was expired.");
        return;
    }
    alert(err);
}

function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}