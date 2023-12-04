function update_value(id, idx) {
    Object.keys(idx).forEach(i => {
        setText(id, i, idx[i].formatted)
    })
}
