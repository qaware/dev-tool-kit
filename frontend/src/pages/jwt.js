export var jwt = {};

jwt.htmlButton = `
    <span class="button" id="pageJwtButton" onclick="window.frontend.loadPage('pageJwt');">
        <i class="fas fa-key"></i> <span>JWT</span>
        <span class="tooltip">View the claims of a JWT and verify its signature</span>
    </span>
`;

jwt.htmlPage = `
    <div id="pageJwt" class="hidden">
        <textarea id="inputJwt" class="soft-wrap clear autodetect-input" placeholder="JSON web token"
            onfocusout="window.frontend.jwt.decode();"
            onpaste="window.frontend.jwt.decodePasted();"></textarea>
        <textarea id="inputJwtKey" class="clear" placeholder="Public key or JSON web key"
            onkeyup="window.frontend.jwt.decode();"></textarea>
        <textarea class="clear" id="outputJwt" placeholder="Payload" readonly></textarea>
    </div>
`;

jwt.decode = function () {
    window.frontend.sendEventToBackend("jwt", "decode", [document.getElementById("inputJwt").value.trim(), document.getElementById("inputJwtKey").value.trim()])
        .then(function (result) {
            document.getElementById("outputJwt").value = result.value;
        });
};

jwt.decodePasted = function () {
    setTimeout(function () {
        window.frontend.jwt.decode();
    }, 10);
};