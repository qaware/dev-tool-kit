export var hex = {};

hex.htmlButton = `
    <span class="button" id="pageHexButton" onclick="window.frontend.loadPage('pageHex');">
        <i class="fas fa-edit"></i> <span>Hex</span>
        <span class="tooltip">View and edit hex codes of input characters</span>
    </span>
`;

hex.htmlPage = `
    <div id="pageHex" class="hidden">
        <textarea id="inputPlainHex" class="clear" placeholder="String" onkeyup="window.frontend.hex.encode();"></textarea>
        <textarea class="soft-wrap-word clear autodetect-input" id="inputEncodedHex" placeholder="Hex"
            onfocusout="window.frontend.hex.decode();"></textarea>
    </div>
`;

hex.encode = function () {
    window.frontend.sendEventToBackend("hex", "encode", [document.getElementById("inputPlainHex").value])
        .then(function (result) {
            document.getElementById("inputEncodedHex").value = result.value;
        });
};

hex.decode = function () {
    window.frontend.sendEventToBackend("hex", "decode", [document.getElementById("inputEncodedHex").value])
        .then(function (result) {
            document.getElementById("inputPlainHex").value = result.value;
        });
};