export var url = {};

url.htmlButton = `
    <span class="button" id="pageUrlButton" onclick="window.frontend.loadPage('pageUrl');">
        <i class="fas fa-network-wired"></i> <span>URL</span>
        <span class="tooltip">Encode special characters in a URL</span>
    </span>
`;

url.htmlPage = `
    <div id="pageUrl" class="hidden">
        <textarea class="soft-wrap clear" id="inputPlainUrl" placeholder="Plain URL"
            onkeyup="window.frontend.url.encode();"></textarea>
        <textarea class="soft-wrap clear" id="inputEncodedUrl" placeholder="Encoded URL"
            onfocusout="window.frontend.url.decode();"
            onpaste="window.frontend.url.decodePasted();"></textarea>
    </div>
`;

url.encode = function () {
    window.frontend.sendEventToBackend("url", "encode", [document.getElementById("inputPlainUrl").value.trim()])
        .then(function (result) {
            document.getElementById("inputEncodedUrl").value = result.value;
        });
};

url.decode = function () {
    window.frontend.sendEventToBackend("url", "decode", [document.getElementById("inputEncodedUrl").value.trim()])
        .then(function (result) {
            document.getElementById("inputPlainUrl").value = result.value;
        });
};

url.decodePasted = function () {
    setTimeout(function () {
        window.frontend.url.decode();
    }, 10);
};