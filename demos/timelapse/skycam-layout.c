
import (
    "github.com/peter-mount/go-anim/script/image"
    "github.com/peter-mount/go-anim/script/layout"
    "github.com/peter-mount/piweather.center/script/weather/keogram"
)

createLayout(cfg) {
    cfg.layout = layout.New(image.Width4K,image.Height4K).
        RowContainer().
            ColScaleContainer(1/3.0,1/3.0,1/3.0).
                Font("luxi 32 mono bold").
                Fill( cfg.white ).
                Text("",cfg.title).End().
                Text("","ME15Weather").Align("center").End().
                Text("timeDisplay","%s").Align("right").End().
            End().
            ColScaleContainer(0.4,0.4,0.2).
                Image("skyCamera").Inset(10).End().
                Image("auxView").Inset(10).End().
                RowContainer().
                    Font("luxi 20 mono bold").
                    Fill( cfg.white ).
                    Value("cloudCover", "Cloud Cover","%3.0f%% Sky %3.0f%% Obscured %3.0f%%",0,0,0).End().
                    Value("sunRaDec", "Sun Position","%s %s",0,0).End().
                    Value("sunAltAz", "Sun Position","%s %s",0,0).End().
                    Value("sunDist","Distance","%s").End().
                    Value("sunTime","Light Time","%s").End().
                    Value("moonAltAz", "Moon Position","%s %s",0,0).End().
                    Value("moonDist","Distance","%s").End().
                    Value("moonTime","Light Time","%s").End().
                End().
            End().
            ColScaleContainer(0.4,0.4,0.2).
                AddComponent( "keogram", keogram.Keogram() ).End().
            End().
        End().
    End().
    Build()
}