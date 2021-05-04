export var uuid = {};

uuid.htmlButton = `
    <span class="button" id="pageUuidButton" onclick="window.frontend.loadPage('pageUuid'); window.frontend.uuid.load();">
        <i class="fas fa-address-card"></i> <span>UUID</span>
        <span class="tooltip">Generate UUIDs</span>
    </span>
`;

uuid.htmlPage = `
    <div id="pageUuid" class="hidden">
        <textarea id="outputUuid" readonly></textarea>
    </div>
`;

uuid.load = function () {
    window.frontend.sendEventToBackend("uuid", "", [])
        .then(function (result) {
            document.getElementById("outputUuid").value = result.value;
        });
};