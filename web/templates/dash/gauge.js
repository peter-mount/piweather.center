function update_gauge(m,id,idx) {
    idx.forEach(i => {
        let d = document.getElementById(id + ".svg")
        idx.forEach(i => {
            setText(id, i, m.formatted)
            let v = Math.max(d.dataset.min, Math.min(m.value, d.dataset.max))
            setRotate(id, i, ((v - d.dataset.min) * d.dataset.delta) - 90)
        })
    })
}
