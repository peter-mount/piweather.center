/* value.js */
function update_compass(r) {
    for (let id of r.actions['compass']) {
        let e = document.getElementById(id+".ptr")
        if (e !==null) {e.setAttribute("transform",'rotate('+r.metric.value+')')}
        e = document.getElementById(id+".txt")
        if (e !==null) {e.textContent=r.metric.formatted}
    }
}
