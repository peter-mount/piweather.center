window.addEventListener("load", function (evt) {
    let ws = new WebSocket("ws://127.0.0.1:8080/live/dash/{{$.dash}}");

    ws.onclose = function (evt) {
        ws = null;
    }

    ws.onerror = function (evt) {
    }

    ws.onmessage = function (evt) {
        let msg = JSON.parse(evt.data), m=msg.metric,acts=msg.actions
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
});

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
