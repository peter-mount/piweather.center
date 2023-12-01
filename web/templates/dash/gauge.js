/* value.js */
function update_gauge(r) {
    for (let id of r.actions['gauge']) {
        let e = document.getElementById(id+".txt")
        if (e !==null) {e.textContent=r.metric.formatted}

        let d = document.getElementById(id+".svg")
        e = document.getElementById(id+".ptr")
        if (d!==null && e !==null) {
            let v=Math.max(d.dataset.min,Math.min(r.metric.value,d.dataset.max))
            let a=((v-d.dataset.min)*d.dataset.delta)-90
            e.setAttribute("transform",'rotate('+a+')')
        }

    }
}
