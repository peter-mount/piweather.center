function update_gauge(m, id, idx) {
    idx.forEach(i => {
        let d = document.getElementById(id + ".svg"),
            min = d.dataset.min,
            max = d.dataset.max
        idx.forEach(i => {
            setText(id, i, m.formatted)
            let v = ensureWithin(m.value, min, max)
            setRotate(id, i, ((v - d.dataset.min) * d.dataset.delta) - 90)
        })
    })
}
