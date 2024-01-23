function update_inclinometer(id, idx) {
    Object.keys(idx).forEach(i => {
        let m = idx[i],
            e = document.getElementById(id + ".ptr" + i),
            v = 90 - ensureWithin(m.value, -90, 90)
        setText(id, i, m.formatted)
        e.setAttribute("transform", "rotate(" + v + ")")
    })
}
