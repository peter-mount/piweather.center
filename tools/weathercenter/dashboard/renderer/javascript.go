package renderer

import (
	"strings"
	"sync"
)

var (
	mutex sync.Mutex
	js    map[string]string
)

const (
	liveWsPrefix = "/ws/d/"
	mainJs1      = `
window.addEventListener("load",wsListener);function wsListener(evt){` +
		`let url=(location.protocol==="http:"?"ws":"wss")+"://"+location.host+"`
	mainJs2 = `",dashUid="`
	mainJs3 = `` +
		`",ws=new WebSocket(url);` +
		`ws.onclose=function(evt){ws=null;setTimeout(wsListener,2000)};` +
		// on message dispatch to correct handler
		`ws.onmessage=function(evt){` +
		` let msg=JSON.parse(evt.data),acts=msg.actions;` +
		` if(msg.uid!==dashUid){location.reload();return}` +
		` Object.keys(acts).forEach(k=>{` +
		`  let f=actions[k];if(f){` +
		`   let ids=acts[k];Object.keys(ids).forEach(id=>{f(id,ids[id])})` +
		`  }` +
		` })` +
		`};` +
		`return false;}` +
		// rotate by angle
		`;function setRotate(id,i,ang){` +
		`let e=document.getElementById(id+".ptr"+i);` +
		`if(e!==null) {ang=ang+(0>ang?360:ang>=360?-360:0);` +
		`let from=e.getAttribute("to"),` +
		`d=Math.abs(from-ang)>180,` +
		`td=d&&ang>from,` +
		`fd=d&&from>ang;` +
		`e.setAttribute("from",from-(fd?360:0));` +
		`e.setAttribute("to",ang-(td?360:0));` +
		`e.beginElement()}` +
		`}` +
		// setText(id,i,t)
		`;function setText(id,i,t){` +
		`let e=document.getElementById(id+".txt"+i);` +
		`if(e!==null){e.textContent=t}` +
		`}` +
		// ensure within bounds
		`;function ensureWithin(v,min,max){return Math.max(min,Math.min(v,max))}`
)

func LiveWsPath(s, d string) string {
	return liveWsPrefix + s + "/" + d
}

// registerJs registers javascript for a specific component type.
// Only the body of the update function is required. The definition will be created during registration.
//
// Note: the following signature is used (here for type "value"):
// function update_value(id,idx){ body }
func registerJs(t, body string) {
	if js == nil {
		js = make(map[string]string)
	}
	js[t] = ";function " + updateFunc(t) + "(id,idx){" + body + "}"
}

func GetJavaScript(t string) string {
	mutex.Lock()
	defer mutex.Unlock()
	return js[t]
}

func HasJavaScript(t string) bool {
	if t == "" {
		return false
	}

	mutex.Lock()
	defer mutex.Unlock()
	_, exists := js[t]
	return exists
}

func updateFunc(t string) string {
	return "update_" + strings.ReplaceAll(t, "-", "_")
}

func GenerateJavaScript(stId, dashId, dashUid string, actions map[string]interface{}) string {
	// Ensure we only include types we have javascript for
	lookup := make(map[string]string)
	for k, _ := range actions {
		if HasJavaScript(k) {
			lookup[k] = updateFunc(k)
		}
	}
	// Nothing so stop here
	if len(lookup) == 0 {
		return ""
	}

	// Generate the action lookup map
	var s []string
	s = append(s, `let actions={`, strings.Join(s, ","))
	for k, v := range lookup {
		s = append(s, k, ":", v, ",")
	}
	// Replace last element "," with "};"
	s[len(s)-1] = "};"

	s = append(s, mainJs1, LiveWsPath(stId, dashId), mainJs2, dashUid, mainJs3)

	for k, _ := range lookup {
		s = append(s, GetJavaScript(k))
	}

	return strings.Join(s, "")
}
