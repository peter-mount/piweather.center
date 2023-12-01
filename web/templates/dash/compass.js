function update_compass(m, id, idx) {
    idx.forEach(i => {
        setRotate(id, i, m.value)
        setText(id, i, "" + Math.floor(m.value) + '°')
    })
}
