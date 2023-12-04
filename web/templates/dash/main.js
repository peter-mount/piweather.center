window.addEventListener("load", wsListener)

function wsListener(evt) {
    let pt = location.port,
        url=(location.protocol === "http:" ? "ws" : "wss") + "://" + location.host + "/live/dash/{{$.dash}}",
        dashUid = "{{$.board.Uid}}",
        ws = new WebSocket(url);

    ws.onopen = function (evt) {}

    ws.onclose = function (evt) {
        ws = null;
        setTimeout(wsListener, 2000)
    }

    ws.onerror = function (evt) {}

    ws.onmessage = function (evt) {
        let msg = JSON.parse(evt.data), acts = msg.actions

        if (msg.uid !== dashUid) {
            location.reload()
            return
        }

        Object.keys(acts).forEach(k => {
            let f = actions[k]
            if (f) {
                let ids = acts[k]
                Object.keys(ids).forEach(id => {
                    f(id, ids[id])
                })
            }
        })
    }

    return false;
}

function setRotate(id, i, ang) {
    let e = document.getElementById(id + ".ptr" + i)
    if (e !== null) {
        let from=e.getAttribute("to"),
            d=Math.abs(from-ang)>180,
            td=d && from<ang,
            fd=d && from>ang
        e.setAttribute("from",from-(fd?360:0))
        e.setAttribute("to",ang-(td?360:0))
        e.beginElement()
    }
}

function setText(id, i, t) {
    let e = document.getElementById(id + ".txt" + i)
    if (e !== null) {
        e.textContent = t
    }
}

function ensureWithin(v, min, max) {
    return Math.max(min, Math.min(v, max))
}