import { envWebSocket } from '../config/env'
import { wsAccion } from '@astric-store'

const setSocket = async (ws: any) => {
    const data = JSON.stringify({ databases: ["totems", "atric"], usuario: "dev" })
    ws.send(data)
}

const heartbeat = async (ws: any) => {
    if (!ws) return;
    if (ws.readyState !== 1) return;
    await setSocket(ws)
    let time = parseInt(envWebSocket.time == undefined ? "1" : envWebSocket.time, 10)
    setTimeout(heartbeat, time * 1000);
}

export let webSocket = () => {
    if (envWebSocket.enable) {
        const ws = new WebSocket(envWebSocket.URL || "ws://localhost:8000")
        ws.addEventListener("open", () => heartbeat(ws))
        ws.addEventListener("close", webSocket)
        ws.addEventListener("message", (dato: any) => {
            wsAccion.set(dato)
        })
        return ws
    }
    return
}






