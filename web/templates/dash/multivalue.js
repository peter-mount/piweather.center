function update_multivalue(id, idx) {
    Object.keys(idx).forEach(i => {
        setText(id, i, idx[i].formatted)
    })
}
