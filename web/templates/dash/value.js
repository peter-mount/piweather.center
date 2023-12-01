/* value.js */
function update_value(r) {
    for (let id of r.actions['value']) {
        let e = document.getElementById(id)
        if (e !==null) {e.textContent=r.metric.formatted}
    }
}
