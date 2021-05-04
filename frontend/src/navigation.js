export var navigation = {};

export function setupNavigation() {
    var clearButton = `
        <span class="button" id="clearButton" onclick="window.frontend.clearPage();">
            <i class="fas fa-trash-alt"></i>
        <span class="tooltip">Clear all visible input fields</span>
        </span>
    `;

    var filterInput = `
        <input type="text" id="filter" class="hidden" placeholder="Filter" onkeyup="window.frontend.navigation.filter(event);" onfocusout="window.frontend.navigation.stopFilter();">
    `;

    var nav = document.getElementById("nav");
    nav.innerHTML = window.frontend.auto.htmlButton
        + window.frontend.ascii.htmlButton
        + window.frontend.base64.htmlButton
        + window.frontend.auth.htmlButton
        + window.frontend.calculator.htmlButton
        + window.frontend.diff.htmlButton
        + window.frontend.gzip.htmlButton
        + window.frontend.hex.htmlButton
        + window.frontend.http_client.htmlButton
        + window.frontend.http_server.htmlButton
        + window.frontend.json.htmlButton
        + window.frontend.jwt.htmlButton
        + window.frontend.log.htmlButton
        + window.frontend.ports.htmlButton
        + window.frontend.tunnel.htmlButton
        + window.frontend.time.htmlButton
        + window.frontend.url.htmlButton
        + window.frontend.uuid.htmlButton
        + window.frontend.xml.htmlButton
        + "<!--placeholder-->"
        + filterInput
        + clearButton
        + window.frontend.upgrade.htmlButton;

    document.addEventListener("keydown", function (event) {
        if (event.key === "Control") {
            window.frontend.navigation.ctrl = true;
        } else if (window.frontend.navigation.ctrl && event.key === "f") {
            window.frontend.navigation.startFilter();
            window.frontend.navigation.ctrl = false;
            event.stopPropagation();
            event.preventDefault();
        }
    });

    document.addEventListener("keyup", function (event) {
        window.frontend.navigation.ctrl = false;
    });
}

navigation.ctrl = false;

navigation.startFilter = function () {
    var input = document.getElementById("filter");
    input.classList.remove("hidden");
    input.focus();
    input.select();
};

navigation.stopFilter = function () {
    var input = document.getElementById("filter");
    input.value = "";
    input.classList.add("hidden");
    document.body.focus();

    var buttons = document.getElementById("nav").getElementsByClassName("button");
    for (var i = 0; i < buttons.length; i++) {
        buttons[i].classList.remove("fade");
    }
};

navigation.filter = function (event) {
    if (event.key === "Enter" && window.frontend.navigation.clickFilteredButton()) {
        window.frontend.navigation.stopFilter();
        return;
    }

    if (event.key === "Escape" || event.key === "Esc") {
        window.frontend.navigation.stopFilter();
        return;
    }

    var query = document.getElementById("filter").value.trim().toLowerCase();
    var buttons = document.getElementById("nav").getElementsByClassName("button");
    for (var i = 0; i < buttons.length; i++) {
        var buttonText = buttons[i].getElementsByTagName("span")[0];
        if (query.length === 0 || (buttonText !== undefined && buttonText.innerText.toLowerCase().includes(query))) {
            buttons[i].classList.remove("fade");
        } else {
            buttons[i].classList.add("fade");
        }
    }
};

navigation.clickFilteredButton = function () {
    var filteredButton = null;
    var buttons = document.getElementById("nav").getElementsByClassName("button");
    for (var i = 0; i < buttons.length; i++) {
        if (!buttons[i].classList.contains("fade")) {
            if (filteredButton != null) {
                return false;
            }
            filteredButton = buttons[i];
        }
    }

    if (filteredButton != null) {
        filteredButton.click();
        return true;
    }
    return false;
};