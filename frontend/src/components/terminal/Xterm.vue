<template>
    <div ref="terminalRef" class="terminalArea"></div>
  </template>
  
  <script setup>
  import { Terminal } from "@xterm/xterm";
  import "@xterm/xterm/css/xterm.css";
  import { FitAddon } from "@xterm/addon-fit";
  import { ref, onMounted, onUnmounted } from "vue";
  
  const terminalRef = ref();
  const terminal = new Terminal({
    cursorBlink: true,
    fontSize: 14,
    theme: {
      background: "#1e1e1e",
      foreground: "#ffffff",
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
    ws = new WebSocket(`ws://${import.meta.env.VITE_APP_BASE_URL}/terminal?token=${token}`);
  
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
  });
  
  onUnmounted(() => {
    window.removeEventListener("resize", fitTerminal);
  
    if (ws) {
      ws.close();
      ws = null;
    }
  });
  </script>
  
  <style scoped lang="scss">
  .terminalArea {
    height: 100vh;
    background-color: #1e1e1e;
  }
  </style>
  