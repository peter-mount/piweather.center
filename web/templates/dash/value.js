function update_value(m, id, idx) {
    idx.forEach(i => {
        setText(id, i, m.formatted)
    })
}
