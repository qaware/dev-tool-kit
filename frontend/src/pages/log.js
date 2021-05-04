export var log = {};

log.htmlButton = `
    <span class="button" id="pageLogButton" onclick="window.frontend.loadPage('pageLog'); window.frontend.log.load();">
        <i class="fas fa-clipboard-list"></i> <span>Log</span>
        <span class="tooltip">Save notes in your personal daily log file</span>
    </span>
`;

log.htmlPage = `
    <div id="pageLog" class="hidden">
        <input type="text" id="inputLog" class="clear" placeholder="Log message" onkeyup="window.frontend.log.append(event);">
        <textarea id="outputLog" placeholder="Log entries" onfocusout="window.frontend.log.save();"></textarea>
    </div>
`;

log.load = function () {
    window.frontend.sendEventToBackend("log", "load", [])
        .then(function (result) {
            document.getElementById("outputLog").value = result.value;
        });
};

log.append = function (event) {
    if (event.key !== "Enter") {
        return;
    }

    window.frontend.sendEventToBackend("log", "append", [document.getElementById("inputLog").value.trim()])
        .then(function (result) {
            document.getElementById("outputLog").value = result.value;
            document.getElementById("inputLog").value = "";
        });
};

log.save = function () {
    window.frontend.sendEventToBackend("log", "save", [document.getElementById("outputLog").value]);
};