<template>
  <button
    @click="showTerminal"
    className="flex items-center space-x-2 px-4 py-2 text-sm bg-[#24283b] text-white rounded-md hover:bg-[#414868] transition-colors mb-3"
  >
    <TerminalIcon class="w-4 h-4" />
    <span>Mở terminal</span>
  </button>
</template>

<script setup>
import { TerminalIcon } from "lucide-vue-next";
import { useTerminalStore } from "@/store/terminal";

const terminalStore = useTerminalStore();

function showTerminal() {
  terminalStore.setShowTerminal(true);
}
</script>
