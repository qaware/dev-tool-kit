export var upgrade = {};

upgrade.htmlButton = `
    <span class="button hidden" id="upgradeButton" onclick="window.frontend.upgrade.now();">
        <i class="fas fa-sync-alt"></i> Upgrade
        <span class="tooltip">Upgrade to the latest version</span>
    </span>
`;

upgrade.checkUpgrade = function () {
    window.frontend.sendEventToBackend("upgrade", "getVersion", [])
        .then(function (result) {
            document.getElementById("version").innerText = result.value;

            window.frontend.sendEventToBackend("upgrade", "checkUpgrade", [])
                .then(function (hasNewVersion) {
                    if (hasNewVersion.value === "true") {
                        document.getElementById("upgradeButton").classList.remove("hidden");
                    } else {
                        document.getElementById("upgradeButton").classList.add("hidden");
                    }
                });
        });
};

upgrade.now = function () {
    document.getElementById("upgradeButton").classList.add("hidden");
    document.getElementById("upgrading").classList.remove("hidden");

    window.frontend.sendEventToBackend("upgrade", "upgradeNow", [])
        .then(function (result) {
            document.getElementById("version").innerText = result.value;
        })
        .finally(function () {
            document.getElementById("upgrading").classList.add("hidden");
        });
    document.body.classList.remove("waiting");
};