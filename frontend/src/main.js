"use strict";

import "core-js/stable";
import "./main.css";
import "./fontawesome-5.14.0/css/all.min.css";

import {clearPage, loadPage, setupPages} from "./pages.js";
import {navigation, setupNavigation} from "./navigation.js";
import {customizeUi} from "./customize.js";
import {receiveEventFromBackend, sendEventToBackend} from "./events.js";
import {licenses} from "./licenses.js";
import {upgrade} from "./upgrade.js";

import {auto} from "./pages/auto.js";
import {ascii} from "./pages/ascii_art.js";
import {base64} from "./pages/base64.js";
import {auth} from "./pages/auth.js";
import {calculator} from "./pages/calculator.js";
import {diff} from "./pages/diff.js";
import {gzip} from "./pages/gzip.js";
import {hex} from "./pages/hex.js";
import {http_client} from "./pages/http_client.js";
import {http_server} from "./pages/http_server.js";
import {json} from "./pages/json.js";
import {jwt} from "./pages/jwt.js";
import {log} from "./pages/log.js";
import {ports} from "./pages/ports.js";
import {tunnel} from "./pages/tunnel.js";
import {time} from "./pages/time.js";
import {url} from "./pages/url.js";
import {uuid} from "./pages/uuid.js";
import {xml} from "./pages/xml.js";

const runtime = require("@wailsapp/runtime");

window.frontend = {};
window.frontend.loadPage = loadPage;
window.frontend.clearPage = clearPage;
window.frontend.sendEventToBackend = sendEventToBackend;
window.frontend.receiveEventFromBackend = receiveEventFromBackend;
window.frontend.navigation = navigation;

window.frontend.auto = auto;
window.frontend.ascii = ascii;
window.frontend.base64 = base64;
window.frontend.auth = auth;
window.frontend.calculator = calculator;
window.frontend.diff = diff;
window.frontend.gzip = gzip;
window.frontend.hex = hex;
window.frontend.http_client = http_client;
window.frontend.http_server = http_server;
window.frontend.json = json;
window.frontend.jwt = jwt;
window.frontend.log = log;
window.frontend.ports = ports;
window.frontend.tunnel = tunnel;
window.frontend.time = time;
window.frontend.url = url;
window.frontend.uuid = uuid;
window.frontend.xml = xml;
window.frontend.licenses = licenses;
window.frontend.upgrade = upgrade;

function start() {
    setupPages();
    setupNavigation();
    customizeUi();

    upgrade.checkUpgrade();

    window.wails.Events.On("httpServer.receive", function (value) {
        window.frontend.http_server.receive(window.frontend.receiveEventFromBackend(value));
    });
}

runtime.Init(start);