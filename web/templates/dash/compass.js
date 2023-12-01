/* value.js */
function update_compass(r) {
    for (let id of r.actions['compass']) {
        let e = document.getElementById(id)
        if (e !==null) {e.setAttribute("transform",'rotate('+r.metric.value+')')}
    }
}
