include "tableselect/daily-3d.c"

main() {
    img := daily3d(map(
        "title": "Hourly Rainfall",
        "yAxisLabel": "Rainfall mm",
        "metric": "home.ecowitt.hrain_piezo",
        "startDate": "2024-07-01",
        "endDate": "2024-08-01",
        "every": "30m",
        // Size of plot area
        "w": 600,
        "h": 250,
        "background": "white",
        "axesColour": "black",
        // Colour of the plot
        "fillColour": "white",
        "strokeColour": "red",
        // DB to use
        "dbUrl": "http://127.0.0.1:8080"
    ))

    try( r := render.New("test%d.png",1) ) {
        r.WriteImage(img)
    } catch(e) {
        fmt.Println("Failed to write image:", e)
    }

}