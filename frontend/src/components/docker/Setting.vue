<template>
  <div class="flex flex-col flex-1 overflow-hidden">
    <div class="p-4 sm:p-6">
      <div class="flex items-center gap-2 mb-6">
        <button
          class="flex items-center text-sm text-blue-600 hover:text-blue-700 md:hidden"
          @click="$emit('openSideBarFn')"
        >
          <ChevronRight class="mr-1" />
        </button>
        <h1 class="text-xl font-semibold text-gray-900">Gỡ cài đặt</h1>
      </div>

      <div class="my-8">
        <div class="flex items-center gap-2 mb-8">
          <span class="text-gray-900">Link:</span>
          <a
            href="https://docs.docker.com/engine/install/ubuntu/#uninstall-docker-engine"
            target="_blank"
            rel="noopener noreferrer"
          >
            https://docs.docker.com/engine/install/ubuntu/#uninstall-docker-engine
          </a>
        </div>

        <button
          @click="openMiniTerminal"
          className="flex items-center space-x-2 px-4 py-2 text-sm bg-[#24283b] text-white rounded-md hover:bg-[#414868] transition-colors mb-3"
        >
          <TerminalIcon class="w-4 h-4" />
          <span>Mở terminal</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { TerminalIcon } from "lucide-vue-next";

function openMiniTerminal() {
  window.open(
    "/#/mini-terminal",
    "_blank",
    "width=800,height=450,top=100,left=200"
  );
}
</script>
