export var tunnel = {};

tunnel.htmlButton = `
    <span class="button" id="pageTunnelButton" onclick="window.frontend.loadPage('pageTunnel');">
        <i class="fas fa-map-signs"></i> <span>SSH tunnel</span>
        <span class="tooltip">Open an SSH tunnel to a remote host</span>
    </span>
`;

tunnel.htmlPage = `
    <div id="pageTunnel" class="hidden">
        <div class="row">
            <input id="inputSshHost" type="text" placeholder="SSH host" class="grow clear">
            <input id="inputSshPort" type="text" placeholder="SSH port" class="right-align clear">
        </div>
        <div class="row">
            <input id="inputSshUsername" type="text" placeholder="SSH username" class="grow clear">
            <input id="inputSshPassword" type="password" placeholder="SSH password" class="grow right-align clear">
        </div>
        <div class="row">
            <input id="inputRemoteHost" type="text" placeholder="Remote host" class="grow clear">
            <input id="inputRemotePort" type="text" placeholder="Remote port" class="clear">
            <input id="inputLocalPort" type="text" placeholder="Local port" class="right-align clear">
        </div>
        <div class="row">
            <span class="button toggle-button" id="toggleTunnel" onclick="window.frontend.tunnel.toggle();">
                <span class="unchecked"><i class="fas fa-play"></i> Open SSH tunnel</span>
                <span class="checked"><i class="fas fa-stop"></i> Close SSH tunnel</span>
            </span>
        </div>
    </div>
`;

tunnel.toggle = function () {
    var sshHost = document.getElementById("inputSshHost").value.trim();
    var sshPort = document.getElementById("inputSshPort").value.trim();
    var sshUser = document.getElementById("inputSshUsername").value.trim();
    var sshPass = document.getElementById("inputSshPassword").value.trim();
    var remoteHost = document.getElementById("inputRemoteHost").value.trim();
    var remotePort = document.getElementById("inputRemotePort").value.trim();
    var localPort = document.getElementById("inputLocalPort").value.trim();

    window.frontend.sendEventToBackend("tunnel", "", [sshHost, sshPort, sshUser, sshPass, remoteHost, remotePort, localPort])
        .then(function () {
            document.getElementById("toggleTunnel").classList.toggle("toggle-active");
            document.getElementById("inputSshHost").classList.toggle("clear");
            document.getElementById("inputSshPort").classList.toggle("clear");
            document.getElementById("inputSshUsername").classList.toggle("clear");
            document.getElementById("inputSshPassword").classList.toggle("clear");
            document.getElementById("inputRemoteHost").classList.toggle("clear");
            document.getElementById("inputRemotePort").classList.toggle("clear");
            document.getElementById("inputLocalPort").classList.toggle("clear");
        });
};