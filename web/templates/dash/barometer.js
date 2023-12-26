function update_barometer(id, idx) {
    Object.keys(idx).forEach(i => {
        let m = idx[i],
            d = document.getElementById(id + ".svg"),
            min = d.dataset.min,
            max = d.dataset.max,
            delta = d.dataset.delta,
            ofs = d.dataset["d" + i],
            v = ensureWithin(m.value, min, max)
        setText(id, i, m.formatted)
        setRotate(id, i, ((v - min) * delta) - 112.5 - ofs)
    })
}
