<template>
  <div class="p-3 m-2 bg-white border rounded-md">
    <button
      @click="openMiniTerminal"
      className="flex items-center space-x-2 px-4 py-2 text-sm bg-[#24283b] text-white rounded-md hover:bg-[#414868] transition-colors mb-3"
    >
      <TerminalIcon class="w-4 h-4" />
      <span>Mở trong cửa sổ riêng</span>
    </button>

    <Xterm />
  </div>
</template>

<script setup>
import { TerminalIcon } from "lucide-vue-next";
import Xterm from "@/components/terminal/Xterm.vue";

function openMiniTerminal() {
  window.open(
    "/#/mini-terminal",
    "_blank",
    "width=800,height=450,top=100,left=200"
  );
}
</script>

<style>
.terminalArea {
  height: 500px !important;
}
</style>
