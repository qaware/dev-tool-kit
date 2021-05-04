export var time = {};

time.htmlButton = `
    <span class="button" id="pageTimeButton" onclick="window.frontend.loadPage('pageTime'); window.frontend.time.load();">
        <i class="fas fa-clock"></i> <span>Timestamp</span>
        <span class="tooltip">Convert a time given in ISO 8601 format to a timestamp and vice versa</span>
    </span>
`;

time.htmlPage = `
    <div id="pageTime" class="hidden">
        <input id="inputDateTime" type="text" placeholder="Time in ISO 8601 format" class="autodetect-input"
            onkeyup="window.frontend.time.convertDateTime();"
            onfocusout="window.frontend.time.convertDateTime();">
        <input id="inputTimestamp" type="text" placeholder="Timestamp in seconds"
            onkeyup="window.frontend.time.convertTimestamp();">
    </div>
`;

time.load = function () {
    window.frontend.sendEventToBackend("time", "load", [])
        .then(function (result) {
            document.getElementById("inputDateTime").value = result.value;
        })
        .then(window.frontend.time.convertDateTime);
};

time.convertDateTime = function () {
    window.frontend.sendEventToBackend("time", "convertDateTime", [document.getElementById("inputDateTime").value.trim()])
        .then(function (result) {
            document.getElementById("inputTimestamp").value = result.value;
        });
};

time.convertTimestamp = function () {
    window.frontend.sendEventToBackend("time", "convertTimestamp", [document.getElementById("inputTimestamp").value.trim()])
        .then(function (result) {
            document.getElementById("inputDateTime").value = result.value;
        });
};