export function customizeUi() {
    window.frontend.sendEventToBackend("customize", "", [])
        .then(function (result) {
            var items = JSON.parse(result.value);

            for (var i = 0; i < items.length; i++) {
                if (items[i].name.length > 0) {
                    var nav = document.getElementById("nav");
                    var index = nav.innerHTML.indexOf("<!--placeholder-->");
                    var button = "<span class=\"button\" id=\"pageCustom" + items[i].name + "Button\" onclick=\"window.frontend.loadPage('pageCustom" + items[i].name + "');\"><i class=\"fas " + items[i].icon + "\"></i> <span>" + items[i].title + "</span><span class=\"tooltip\">" + items[i].tooltip + "</span></span>";

                    nav.innerHTML = nav.innerHTML.slice(0, index) + button + nav.innerHTML.slice(index);

                    document.getElementById("page").innerHTML += "<div id=\"pageCustom" + items[i].name + "\" class=\"hidden\">"+ items[i].body + "</div>";
                }
            }
        });
}