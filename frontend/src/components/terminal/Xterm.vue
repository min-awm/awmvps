<template>
  <div ref="terminalRef" class="pl-2 *:terminalArea"></div>
</template>

<script setup>
import { Terminal } from "@xterm/xterm";
import "@xterm/xterm/css/xterm.css";
import { FitAddon } from "@xterm/addon-fit";
import { ref, onMounted, onUnmounted } from "vue";
import { useTerminalStore } from "@/store/terminal";

const terminalStore = useTerminalStore();
const terminalRef = ref();
// const terminal = new Terminal({
//   cursorBlink: true,
//   fontSize: 14,
//   theme: {
//     background: "#1e1e1e",
//     foreground: "#ffffff",
//   },
// });

const terminal = new Terminal({
  cursorBlink: true,
  fontSize: 14,
  theme: {
    background: "#ffffff",
    foreground: "#000000",
    cursor: "#000000",
    selection: "#d3d3d3", 
  },
});


const fitAddon = new FitAddon();
terminal.loadAddon(fitAddon);

function fitTerminal() {
  fitAddon.fit();
}

let ws = null;

onMounted(() => {
  terminal.open(terminalRef.value);
  fitTerminal();
  window.addEventListener("resize", fitTerminal);

  const token = localStorage.getItem("accessToken");

  const hostname = window.location.hostname;
  const port = window.location.port;
  const ipHost = `${hostname}${port ? `:${port}` : ""}`;
  
  ws = new WebSocket(
    `ws://${
      import.meta.env.VITE_APP_MODE == "dev"
        ? import.meta.env.VITE_APP_BASE_URL
        : ipHost
    }/terminal?token=${token}`
  );

  ws.onmessage = (event) => {
    terminal.write(event.data);
  };

  terminal.onData((data) => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(data);
    }
  });

  ws.onclose = () => {
    terminal.write("\r\n*** Connection closed ***\r\n");
  };

  ws.onerror = (error) => {
    console.error("WebSocket error:", error);
    terminal.write("\r\n*** WebSocket error ***\r\n");
  };

  function writeTerminal(value) {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(value);
    }
  }

  terminalStore.setTerminalWriteFn(writeTerminal);
});

onUnmounted(() => {
  window.removeEventListener("resize", fitTerminal);

  if (ws) {
    ws.close();
    ws = null;
  }
});
</script>
