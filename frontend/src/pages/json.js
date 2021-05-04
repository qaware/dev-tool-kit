export var json = {};

json.htmlButton = `
    <span class="button" id="pageJsonButton" onclick="window.frontend.loadPage('pageJson');">
        <span><span class="icon-font">{ }</span> JSON</span>
        <span class="tooltip">Pretty-print a JSON</span>
    </span>
`;

json.htmlPage = `
    <div id="pageJson" class="hidden">
        <textarea id="inputJson" class="clear autodetect-input" placeholder="Raw JSON"
            onfocusout="window.frontend.json.format();"
            onpaste="window.frontend.json.formatPasted();"></textarea>
        <textarea id="outputJson" class="clear" placeholder="Pretty-printed JSON" readonly></textarea>
    </div>
`;

json.format = function () {
    window.frontend.sendEventToBackend("json", "", [document.getElementById("inputJson").value])
        .then(function (result) {
            document.getElementById("outputJson").value = result.value;
        });
};

json.formatPasted = function () {
    setTimeout(function () {
        window.frontend.json.format();
    }, 10);
};