window.addEventListener("load", wsListener)

function wsListener(evt) {
    let ws = new WebSocket("ws://127.0.0.1:8080/live/dash/{{$.dash}}"),
        dashUuid = "{{$.board.Uuid}}";

    ws.onopen = function (evt) {
        console.log("WS Open")
    }

    ws.onclose = function (evt) {
        ws = null;
        console.log("WS Closed")
        setTimeout(wsListener, 2000)
    }

    ws.onerror = function (evt) {
        //console.log("WS Error", evt)
    }

    ws.onmessage = function (evt) {
        let msg = JSON.parse(evt.data), m = msg.metric, acts = msg.actions

        // Reload the page if the uuid's differ
        if (msg.uuid !== dashUuid) {
            location.reload()
            return
        }

        Object.keys(acts).forEach(k => {
            let f = actions[k]
            if (f) {
                let ids = acts[k]
                Object.keys(ids).forEach(id => {
                    f(m, id, ids[id])
                })
            }
        })
    }

    return false;
}

function setRotate(id, i, ang) {
    let e = document.getElementById(id + ".ptr" + i)
    if (e !== null) {
        e.setAttribute("transform", 'rotate(' + ang + ')')
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