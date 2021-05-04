export var base64 = {};

base64.htmlButton = `
    <span class="button" id="pageBase64Button" onclick="window.frontend.loadPage('pageBase64');">
        <i class="fas fa-mask"></i> <span>Base64</span>
        <span class="tooltip">Encode plain text into base64 and vice versa</span>
    </span>
`;

base64.htmlPage = `
    <div id="pageBase64" class="hidden">
        <textarea id="inputPlainBase64" class="clear" placeholder="Plain text" onkeyup="window.frontend.base64.encode();"></textarea>
        <textarea class="soft-wrap clear autodetect-input" id="inputEncodedBase64" placeholder="Base64-encoded text"
            onfocusout="window.frontend.base64.decode();"
            onpaste="window.frontend.base64.decodePasted();"></textarea>
    </div>
`;

base64.encode = function () {
    window.frontend.sendEventToBackend("base64", "encode", [document.getElementById("inputPlainBase64").value.trim()])
        .then(function (result) {
            document.getElementById("inputEncodedBase64").value = result.value;
        });
};

base64.decode = function () {
    window.frontend.sendEventToBackend("base64", "decode", [document.getElementById("inputEncodedBase64").value.trim()])
        .then(function (result) {
            document.getElementById("inputPlainBase64").value = result.value;
        });
};

base64.decodePasted = function () {
    setTimeout(function () {
        window.frontend.base64.decode();
    }, 10);
};