export var auth = {};

auth.htmlButton = `
    <span class="button" id="pageAuthButton" onclick="window.frontend.loadPage('pageAuth');">
        <i class="fas fa-lock"></i> <span>Basic auth</span>
        <span class="tooltip">Create base64-encoded credentials for HTTP Basic authentication</span>
    </span>
`;

auth.htmlPage = `
    <div id="pageAuth" class="hidden">
        <input type="text" id="user" class="clear" placeholder="Username" onkeyup="window.frontend.auth.encode();">
        <input type="password" id="pass" class="clear" placeholder="Password" onkeyup="window.frontend.auth.encode();">
        <input type="text" id="outputAuth" class="clear" placeholder="Base64-encoded credentials" readonly>
    </div>
`;

auth.encode = function () {
    var user = document.getElementById("user").value.trim();
    var pass = document.getElementById("pass").value.trim();

    window.frontend.sendEventToBackend("auth", "", [user, pass])
        .then(function (result) {
            document.getElementById("outputAuth").value = result.value;
        });
};