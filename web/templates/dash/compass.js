function update_compass(id, idx) {
    Object.keys(idx).forEach(i => {
        let m=idx[i],
            d = document.getElementById(id + ".svg"),
            v = m.value,
            a = v - d.dataset["d" + i]
        setRotate(id, i, a)
        setText(id, i, "" + Math.floor(v) + '°')
    })
}
