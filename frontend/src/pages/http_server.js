export var http_server = {};

http_server.htmlButton = `
    <span class="button" id="pageHttpServerButton" onclick="window.frontend.loadPage('pageHttpServer');">
        <i class="fas fa-globe-americas"></i> <span>HTTP server</span>
        <span class="tooltip">Run an HTTP server on localhost with a defined response</span>
    </span>
`;

http_server.htmlPage = `
    <div id="pageHttpServer" class="hidden">
        <textarea id="inputHttpServer" class="clear" placeholder="HTTP response"></textarea>
        <div class="row">
            <input id="inputHttpServerStatus" class="clear" type="text" placeholder="HTTP status code" value="">
            <input id="inputHttpServerPort" class="clear" type="text" placeholder="Local port" value="">
            <span class="button toggle-button" id="toggleHttpServer"
                onclick="window.frontend.http_server.toggle();">
                <span class="unchecked"><i class="fas fa-play"></i> Start HTTP server</span>
                <span class="checked"><i class="fas fa-stop"></i> Stop HTTP server</span>
            </span>
            <span class="button" onclick="window.frontend.http_server.example();">
                <i class="fas fa-question-circle"></i> Example
            </span>
        </div>
        <textarea id="outputHttpServer" class="clear" placeholder="Received HTTP request" readonly></textarea>
    </div>
`;

http_server.toggle = function () {
    var body = document.getElementById("inputHttpServer").value;
    var status = document.getElementById("inputHttpServerStatus").value.trim();
    var port = document.getElementById("inputHttpServerPort").value.trim();

    document.getElementById("outputHttpServer").value = "";

    window.frontend.sendEventToBackend("httpServer", "", [body, status, port])
        .then(function () {
            document.getElementById("toggleHttpServer").classList.toggle("toggle-active");
            document.getElementById("inputHttpServer").classList.toggle("clear");
            document.getElementById("inputHttpServerStatus").classList.toggle("clear");
            document.getElementById("inputHttpServerPort").classList.toggle("clear");
            document.getElementById("outputHttpServer").classList.toggle("clear");
        });
};

http_server.receive = function (value) {
    document.getElementById("outputHttpServer").value = value;
};

http_server.example = function () {
    if (!document.getElementById("toggleHttpServer").classList.contains("toggle-active")) {
        document.getElementById("inputHttpServer").value = "Hello world!";
        document.getElementById("inputHttpServerStatus").value = "200";
        document.getElementById("inputHttpServerPort").value = "8080";
    }
};