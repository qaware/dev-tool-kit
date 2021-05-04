export var gzip = {};

gzip.htmlButton = `
    <span class="button" id="pageGzipButton" onclick="window.frontend.loadPage('pageGzip');">
        <i class="fas fa-archive"></i> <span>Gzip</span>
        <span class="tooltip">Compress and encode plain text with gzip and vice versa</span>
    </span>
`;

gzip.htmlPage = `
    <div id="pageGzip" class="hidden">
        <textarea id="inputPlainGzip" class="clear" placeholder="Plain text" onkeyup="window.frontend.gzip.encode();"></textarea>
        <textarea class="soft-wrap clear autodetect-input" id="inputEncodedGzip" placeholder="Base64-encoded gzip-compressed text"
            onfocusout="window.frontend.gzip.decode();"
            onpaste="window.frontend.gzip.decodePasted();"></textarea>
    </div>
`;

gzip.encode = function () {
    window.frontend.sendEventToBackend("gzip", "encode", [document.getElementById("inputPlainGzip").value.trim()])
        .then(function (result) {
            document.getElementById("inputEncodedGzip").value = result.value;
        });
};

gzip.decode = function () {
    window.frontend.sendEventToBackend("gzip", "decode", [document.getElementById("inputEncodedGzip").value.trim()])
        .then(function (result) {
            document.getElementById("inputPlainGzip").value = result.value;
        });
};

gzip.decodePasted = function () {
    setTimeout(function () {
        window.frontend.gzip.decode();
    }, 10);
};