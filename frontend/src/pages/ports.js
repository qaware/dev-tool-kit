export var ports = {};

ports.htmlButton = `
    <span class="button" id="pagePortsButton" onclick="window.frontend.loadPage('pagePorts');">
        <i class="fas fa-dungeon"></i> <span>Port scanner</span>
        <span class="tooltip">Scan for open TCP ports on a remote host</span>
    </span>
`;

ports.htmlPage = `
    <div id="pagePorts" class="hidden">
        <div class="row">
            <input id="inputPortsHost" class="grow clear" type="text" placeholder="Remote host" value="">
            <input id="inputPortsFrom" class="clear" type="text" placeholder="Port range start" value="">
            <input id="inputPortsTo" class="right-align clear" type="text" placeholder="Port range end" value="">
        </div>
        <div class="row">
            <span class="button" id="buttonPortsScan" onclick="window.frontend.ports.scan();">
                <i class="fas fa-hourglass-half hidden" id="iconSpinnerPorts"></i><i class="fas fa-play" id="iconScanPorts"></i> Scan ports
            </span>
            <span class="button" onclick="window.frontend.ports.example();">
                <i class="fas fa-question-circle"></i> Example
            </span>
        </div>
        <textarea id="outputPorts" class="clear" placeholder="Open TCP ports" readonly></textarea>
    </div>
`;

ports.scan = function () {
    var button = document.getElementById("buttonPortsScan");
    if (!button.classList.contains("active")) {
        button.classList.add("active");
        document.getElementById("iconScanPorts").classList.add("hidden");
        document.getElementById("iconSpinnerPorts").classList.remove("hidden");
        document.getElementById("inputPortsHost").classList.remove("clear");
        document.getElementById("inputPortsFrom").classList.remove("clear");
        document.getElementById("inputPortsTo").classList.remove("clear");
        document.getElementById("outputPorts").value = "";
        document.getElementById("info").classList.add("hidden");
        document.getElementById("error").classList.add("hidden");
        document.getElementById("footerText").innerText = "";

        var hostname = document.getElementById("inputPortsHost").value.trim();
        var portFrom = document.getElementById("inputPortsFrom").value.trim();
        var portTo = document.getElementById("inputPortsTo").value.trim();

        window.frontend.sendEventToBackend("ports", "scan", [hostname, portFrom, portTo])
            .then(function (result) {
                document.getElementById("outputPorts").value = result.value;
            })
            .finally(function () {
                document.getElementById("buttonPortsScan").classList.remove("active");
                document.getElementById("iconScanPorts").classList.remove("hidden");
                document.getElementById("iconSpinnerPorts").classList.add("hidden");
                document.getElementById("inputPortsHost").classList.add("clear");
                document.getElementById("inputPortsFrom").classList.add("clear");
                document.getElementById("inputPortsTo").classList.add("clear");
            });
        document.body.classList.remove("waiting");
    }
};

ports.example = function () {
    if (document.getElementById("iconSpinnerPorts").classList.contains("hidden")) {
        document.getElementById("inputPortsHost").value = "httpbin.org";
        document.getElementById("inputPortsFrom").value = "80";
        document.getElementById("inputPortsTo").value = "80";
    }
};