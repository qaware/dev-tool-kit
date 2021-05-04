export var http_client = {};

http_client.htmlButton = `
    <span class="button" id="pageHttpClientButton" onclick="window.frontend.loadPage('pageHttpClient');">
        <i class="fas fa-globe-europe"></i> <span>HTTP client</span>
        <span class="tooltip">Send an HTTP request</span>
    </span>
`;

http_client.htmlPage = `
    <div id="pageHttpClient" class="hidden">
        <textarea id="inputHttpClient" class="clear" placeholder="HTTP request"></textarea>
        <div class="row">
            <span class="button" id="buttonHttpSend" onclick="window.frontend.http_client.send();">
                <i class="fas fa-hourglass-half hidden" id="iconSpinnerHttp"></i><i class="fas fa-play" id="iconSendHttp"></i> Send request
            </span>
            <span class="button" onclick="window.frontend.http_client.curl();">
                <i class="fas fa-terminal"></i> Generate cURL
            </span>
            <span class="button" onclick="window.frontend.http_client.example();">
                <i class="fas fa-question-circle"></i> Example
            </span>
        </div>
        <textarea id="outputHttpClient" class="clear" placeholder="Received HTTP response" readonly></textarea>
    </div>
`;

http_client.send = function () {
    var button = document.getElementById("buttonHttpSend");
    if (!button.classList.contains("active")) {
        button.classList.add("active");
        document.getElementById("iconSendHttp").classList.add("hidden");
        document.getElementById("iconSpinnerHttp").classList.remove("hidden");
        document.getElementById("inputHttpClient").classList.remove("clear");
        document.getElementById("outputHttpClient").value = "";
        document.getElementById("outputHttpClient").classList.remove("soft-wrap");
        document.getElementById("info").classList.add("hidden");
        document.getElementById("error").classList.add("hidden");
        document.getElementById("footerText").innerText = "";

        window.frontend.sendEventToBackend("httpClient", "send", [document.getElementById("inputHttpClient").value])
            .then(function (result) {
                document.getElementById("outputHttpClient").value = result.value;
            })
            .finally(function () {
                document.getElementById("buttonHttpSend").classList.remove("active");
                document.getElementById("iconSendHttp").classList.remove("hidden");
                document.getElementById("iconSpinnerHttp").classList.add("hidden");
                document.getElementById("inputHttpClient").classList.add("clear");
            });
        document.body.classList.remove("waiting");
    }
};

http_client.curl = function () {
    document.getElementById("outputHttpClient").classList.add("soft-wrap");
    if (document.getElementById("iconSpinnerHttp").classList.contains("hidden")) {
        window.frontend.sendEventToBackend("httpClient", "curl", [document.getElementById("inputHttpClient").value])
            .then(function (result) {
                document.getElementById("outputHttpClient").value = result.value;
            });
    }
};

http_client.example = function () {
    if (document.getElementById("iconSpinnerHttp").classList.contains("hidden")) {
        window.frontend.sendEventToBackend("httpClient", "example", [])
            .then(function (result) {
                document.getElementById("inputHttpClient").value = result.value;
            });
    }
};