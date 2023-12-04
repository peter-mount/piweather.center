function update_compass(m, id, idx) {
    idx.forEach(i => {
        let d = document.getElementById(id + ".svg"),
            v = m.value,
            a = v - d.dataset["d" + idx]
        setRotate(id, i, a)
        setText(id, i, "" + Math.floor(v) + '°')
    })
}
