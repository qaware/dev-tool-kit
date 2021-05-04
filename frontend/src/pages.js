export function setupPages() {
    var app = document.getElementById("app");
    app.innerHTML = `
        <div id="nav"></div>
        
        <div id="page">
            <div id="pageLicense" class="hidden">
                <div id="licenseContent"></div>
            </div>
        </div>
        
        <div id="footer">
            <i class="fas fa-info-circle hidden" id="info"></i>
            <i class="fas fa-exclamation-triangle hidden" id="error"></i>
            <span id="footerText"></span>
            <span id="version"></span>
            <i class="fas fa-hourglass-half hidden" id="upgrading"></i>
            <i class="fas fa-balance-scale" id="pageLicenseButton" onclick="window.frontend.loadPage('pageLicense');"></i>
        </div>
    `;

    var page = document.getElementById("page");
    page.innerHTML = page.innerHTML
        + window.frontend.auto.htmlPage
        + window.frontend.ascii.htmlPage
        + window.frontend.base64.htmlPage
        + window.frontend.auth.htmlPage
        + window.frontend.calculator.htmlPage
        + window.frontend.diff.htmlPage
        + window.frontend.gzip.htmlPage
        + window.frontend.hex.htmlPage
        + window.frontend.http_client.htmlPage
        + window.frontend.http_server.htmlPage
        + window.frontend.json.htmlPage
        + window.frontend.jwt.htmlPage
        + window.frontend.log.htmlPage
        + window.frontend.ports.htmlPage
        + window.frontend.tunnel.htmlPage
        + window.frontend.time.htmlPage
        + window.frontend.url.htmlPage
        + window.frontend.uuid.htmlPage
        + window.frontend.xml.htmlPage;

    document.getElementById("licenseContent").innerHTML = window.frontend.licenses;
}

export function loadPage(page) {
    document.getElementById("info").classList.add("hidden");
    document.getElementById("error").classList.add("hidden");
    document.getElementById("footerText").innerText = "";

    var pages = document.getElementById("page").children;
    for (var i = 0; i < pages.length; i++) {
        pages[i].classList.add("hidden");
    }
    var buttons = document.getElementById("nav").getElementsByClassName("button");
    for (var j = 0; j < buttons.length; j++) {
        buttons[j].classList.remove("active");
    }
    document.getElementById(page).classList.remove("hidden");
    document.getElementById(page + "Button").classList.add("active");
}

export function clearPage() {
    var pages = document.getElementById("page").children;
    var currentPage = pages[0];
    for (var i = 0; i < pages.length; i++) {
        if (!pages[i].classList.contains("hidden")) {
            currentPage = pages[i];
            break;
        }
    }

    var fields = currentPage.getElementsByClassName("clear");
    for (var j = 0; j < fields.length; j++) {
        if (fields[j].tagName.toLowerCase() === "div") {
            fields[j].innerHTML = "";
        } else {
            fields[j].value = "";
        }
    }

    document.getElementById("info").classList.add("hidden");
    document.getElementById("error").classList.add("hidden");
    document.getElementById("footerText").innerText = "";
}