function update_rain_gauge(id, idx) {
    Object.keys(idx).forEach(i => {
        let m = idx[i],
            d = document.getElementById(id + ".svg"),
            min = d.dataset.min,
            max = d.dataset.max,
            scale = d.dataset.scale,
            height = d.dataset.height,
            v = m.value,
            y = scale * (v - min)
        if (v > max) {
            // Update means we exceed the axis so reload to get a new axis
            location.reload()
            return
        }
        let e = document.getElementById(id + ".rect")
        e.setAttribute("y", height - y)
        e.setAttribute("height", y)
        setText(id, i, m.formatted)
    })
}
