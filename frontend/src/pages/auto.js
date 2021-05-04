export var auto = {};

auto.htmlButton = `
    <span class="button active" id="pageAutoButton" onclick="window.frontend.loadPage('pageAuto');">
        <i class="fas fa-chess"></i> <span>Smart mode</span>
        <span class="tooltip">Let the smart mode detect the required tool</span>
    </span>
`;

auto.htmlPage = `
    <div id="pageAuto">
        <textarea id="inputAuto" class="clear" placeholder="Paste anything here you want to decode or format"
            onpaste="window.frontend.auto.detectPasted();"></textarea>
    </div>
`;

auto.detect = function () {
    window.frontend.sendEventToBackend("auto", "detect", [document.getElementById("inputAuto").value.trim()])
        .then(function (result) {
            if (result.value.length > 0) {
                window.frontend.loadPage(result.value);

                var inputs = document.getElementById(result.value).getElementsByClassName("autodetect-input");
                for (var i = 0; i < inputs.length; i++) {
                    inputs[i].focus();
                    inputs[i].value = document.getElementById("inputAuto").value;
                    inputs[i].blur();
                }

                document.getElementById("inputAuto").value = "";
            }
        });
};

auto.detectPasted = function () {
    setTimeout(function () {
        window.frontend.auto.detect();
    }, 10);
};