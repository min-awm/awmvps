<template>
  <div class="p-3 m-2 bg-white border rounded-md">
    <OpenTerminalButton />

    <Xterm />
  </div>
</template>

<script setup>
import OpenTerminalButton from "@/components/common/OpenTerminalButton.vue";
import Xterm from "@/components/terminal/Xterm.vue";
</script>
